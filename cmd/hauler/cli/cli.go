package cli

import (
	"context"

	"freighter.dev/go/freighter/internal/flags"
	"freighter.dev/go/freighter/pkg/consts"
	"freighter.dev/go/freighter/pkg/log"
	cranecmd "github.com/google/go-containerregistry/cmd/crane/cmd"
	"github.com/spf13/cobra"
)

func New(ctx context.Context, ro *flags.CliRootOpts) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "freighter",
		Short:   "Airgap Swiss Army Knife",
		Example: "  View the Docs: https://docs.freighter.dev\n  Environment Variables: " + consts.FreighterDir + " | " + consts.FreighterTempDir + " | " + consts.FreighterStoreDir + " | " + consts.FreighterIgnoreErrors,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			l := log.FromContext(ctx)
			l.SetLevel(ro.LogLevel)
			l.Debugf("running cli command [%s]", cmd.CommandPath())

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	flags.AddRootFlags(cmd, ro)

	cmd.AddCommand(cranecmd.NewCmdAuthLogin("freighter"))
	addStore(cmd, ro)
	addVersion(cmd, ro)
	addCompletion(cmd, ro)

	return cmd
}
