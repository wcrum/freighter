package flags

import "github.com/spf13/cobra"

type CliRootOpts struct {
	LogLevel     string
	FreighterDir string
	IgnoreErrors bool
}

func AddRootFlags(cmd *cobra.Command, ro *CliRootOpts) {
	pf := cmd.PersistentFlags()

	pf.StringVarP(&ro.LogLevel, "log-level", "l", "info", "Set the logging level (i.e. info, debug, warn)")
	pf.StringVarP(&ro.FreighterDir, "freighterdir", "d", "", "Set the location of the freighter directory (default $HOME/.freighter)")
	pf.BoolVar(&ro.IgnoreErrors, "ignore-errors", false, "Ignore/Bypass errors (i.e. warn on error) (defaults false)")
}
