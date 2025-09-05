package store

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"

	"freighter.dev/go/freighter/internal/flags"
	"freighter.dev/go/freighter/internal/mapper"
	"freighter.dev/go/freighter/pkg/log"
	"freighter.dev/go/freighter/pkg/reference"
	"freighter.dev/go/freighter/pkg/store"
)

func ExtractCmd(ctx context.Context, o *flags.ExtractOpts, s *store.Layout, ref string) error {
	l := log.FromContext(ctx)

	r, err := reference.Parse(ref)
	if err != nil {
		return err
	}

	// use the repository from the context and the identifier from the reference
	repo := r.Context().RepositoryStr() + ":" + r.Identifier()

	found := false
	if err := s.Walk(func(reference string, desc ocispec.Descriptor) error {
		if !strings.Contains(reference, repo) {
			return nil
		}
		found = true

		rc, err := s.Fetch(ctx, desc)
		if err != nil {
			return err
		}
		defer rc.Close()

		var m ocispec.Manifest
		if err := json.NewDecoder(rc).Decode(&m); err != nil {
			return err
		}

		mapperStore, err := mapper.FromManifest(m, o.DestinationDir)
		if err != nil {
			return err
		}

		pushedDesc, err := s.Copy(ctx, reference, mapperStore, "")
		if err != nil {
			return err
		}

		l.Infof("extracted [%s] from store with digest [%s]", pushedDesc.MediaType, pushedDesc.Digest.String())

		return nil
	}); err != nil {
		return err
	}

	if !found {
		return fmt.Errorf("reference [%s] not found in store (hint: use `freighter store info` to list store contents)", ref)
	}

	return nil
}
