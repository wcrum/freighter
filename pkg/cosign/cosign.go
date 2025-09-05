package cosign

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"freighter.dev/go/freighter/internal/flags"
	"freighter.dev/go/freighter/pkg/artifacts/image"
	"freighter.dev/go/freighter/pkg/consts"
	"freighter.dev/go/freighter/pkg/log"
	"freighter.dev/go/freighter/pkg/store"
	"github.com/sigstore/cosign/v2/cmd/cosign/cli"
	"github.com/sigstore/cosign/v2/cmd/cosign/cli/options"
	"github.com/sigstore/cosign/v2/cmd/cosign/cli/verify"
	"oras.land/oras-go/pkg/content"
)

// VerifySignature verifies the digital signature of a file using Sigstore/Cosign.
func VerifySignature(ctx context.Context, s *store.Layout, keyPath string, useTlog bool, ref string, rso *flags.StoreRootOpts, ro *flags.CliRootOpts) error {
	l := log.FromContext(ctx)
	operation := func() error {
		v := &verify.VerifyCommand{
			KeyRef:     keyPath,
			IgnoreTlog: true, // Ignore transparency log by default.
		}

		// if the user wants to use the transparency log, set the flag to false
		if useTlog {
			v.IgnoreTlog = false
		}

		err := log.CaptureOutput(l, true, func() error {
			return v.Exec(ctx, []string{ref})
		})
		if err != nil {
			return err
		}

		return nil
	}

	return RetryOperation(ctx, rso, ro, operation)
}

// VerifyKeylessSignature verifies the digital signature of a file using Sigstore/Cosign.
func VerifyKeylessSignature(ctx context.Context, s *store.Layout, identity string, identityRegexp string, oidcIssuer string, oidcIssuerRegexp string, ghWorkflowRepository string, useTlog bool, ref string, rso *flags.StoreRootOpts, ro *flags.CliRootOpts) error {
	l := log.FromContext(ctx)
	operation := func() error {

		certVerifyOptions := options.CertVerifyOptions{
			CertOidcIssuer:               oidcIssuer,
			CertOidcIssuerRegexp:         oidcIssuer,
			CertIdentity:                 identity,
			CertIdentityRegexp:           identityRegexp,
			CertGithubWorkflowRepository: ghWorkflowRepository,
		}

		v := &verify.VerifyCommand{
			CertVerifyOptions:            certVerifyOptions,
			IgnoreTlog:                   false, // Ignore transparency log is set to false by default for keyless signature verification
			CertGithubWorkflowRepository: ghWorkflowRepository,
		}

		// if the user wants to use the transparency log, set the flag to false
		if useTlog {
			v.IgnoreTlog = false
		}

		err := log.CaptureOutput(l, true, func() error {
			return v.Exec(ctx, []string{ref})
		})
		if err != nil {
			return err
		}

		return nil
	}

	return RetryOperation(ctx, rso, ro, operation)
}

// SaveImage saves image and any signatures/attestations to the store.
func SaveImage(ctx context.Context, s *store.Layout, ref string, platform string, rso *flags.StoreRootOpts, ro *flags.CliRootOpts) error {
	l := log.FromContext(ctx)

	if !ro.IgnoreErrors {
		envVar := os.Getenv(consts.FreighterIgnoreErrors)
		if envVar == "true" {
			ro.IgnoreErrors = true
		}
	}

	operation := func() error {
		o := &options.SaveOptions{
			Directory: s.Root,
		}

		// check to see if the image is multi-arch
		isMultiArch, err := image.IsMultiArchImage(ref)
		if err != nil {
			return err
		}
		l.Debugf("multi-arch image [%v]", isMultiArch)

		// Note: Platform support removed in newer cosign versions
		// Platform filtering is now handled differently in cosign v2.5.3+
		if platform != "" && isMultiArch {
			l.Debugf("platform for image [%s] - note: platform filtering not supported in current cosign version", platform)
		}

		err = cli.SaveCmd(ctx, *o, ref)
		if err != nil {
			return err
		}

		return nil

	}

	return RetryOperation(ctx, rso, ro, operation)
}

// LoadImage loads store to a remote registry.
func LoadImages(ctx context.Context, s *store.Layout, registry string, only string, ropts content.RegistryOptions, ro *flags.CliRootOpts) error {
	l := log.FromContext(ctx)

	o := &options.LoadOptions{
		Directory: s.Root,
		Registry:  options.RegistryOptions{
			// Note: Name field removed in newer cosign versions
			// Registry is now specified as part of the command arguments
		},
	}

	// Note: LoadOnly field removed in newer cosign versions
	// Filtering is now handled differently in cosign v2.5.3+
	if len(only) > 0 {
		l.Debugf("load only filter [%s] - note: filtering not supported in current cosign version", only)
	}

	if ropts.Insecure {
		o.Registry.AllowInsecure = true
	}

	if ropts.PlainHTTP {
		o.Registry.AllowHTTPRegistry = true
	}

	if ropts.Username != "" {
		o.Registry.AuthConfig.Username = ropts.Username
		o.Registry.AuthConfig.Password = ropts.Password
	}

	// execute the cosign load and capture the output in our logger
	err := log.CaptureOutput(l, false, func() error {
		return cli.LoadCmd(ctx, *o, registry)
	})
	if err != nil {
		return err
	}

	return nil
}

func RetryOperation(ctx context.Context, rso *flags.StoreRootOpts, ro *flags.CliRootOpts, operation func() error) error {
	l := log.FromContext(ctx)

	if !ro.IgnoreErrors {
		envVar := os.Getenv(consts.FreighterIgnoreErrors)
		if envVar == "true" {
			ro.IgnoreErrors = true
		}
	}

	// Validate retries and fall back to a default
	retries := rso.Retries
	if retries <= 0 {
		retries = consts.DefaultRetries
	}

	for attempt := 1; attempt <= rso.Retries; attempt++ {
		err := operation()
		if err == nil {
			// If the operation succeeds, return nil (no error)
			return nil
		}

		if ro.IgnoreErrors {
			if strings.HasPrefix(err.Error(), "function execution failed: no matching signatures: rekor client not provided for online verification") {
				l.Warnf("warning (attempt %d/%d)... failed tlog verification", attempt, rso.Retries)
			} else {
				l.Warnf("warning (attempt %d/%d)... %v", attempt, rso.Retries, err)
			}
		} else {
			if strings.HasPrefix(err.Error(), "function execution failed: no matching signatures: rekor client not provided for online verification") {
				l.Errorf("error (attempt %d/%d)... failed tlog verification", attempt, rso.Retries)
			} else {
				l.Errorf("error (attempt %d/%d)... %v", attempt, rso.Retries, err)
			}
		}

		// If this is not the last attempt, wait before retrying
		if attempt < rso.Retries {
			time.Sleep(time.Second * consts.RetriesInterval)
		}
	}

	// If all attempts fail, return an error
	return fmt.Errorf("operation unsuccessful after %d attempts", rso.Retries)
}
