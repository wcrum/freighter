package flags

import (
	"freighter.dev/go/freighter/pkg/consts"
	"github.com/spf13/cobra"
)

type SaveOpts struct {
	*StoreRootOpts
	FileName string
	Platform string
}

func (o *SaveOpts) AddFlags(cmd *cobra.Command) {
	f := cmd.Flags()

	f.StringVarP(&o.FileName, "filename", "f", consts.DefaultFreighterArchiveName, "(Optional) Specify the name of outputted haul")
	f.StringVarP(&o.Platform, "platform", "p", "", "(Optional) Specify the platform for runtime imports... i.e. linux/amd64 (unspecified implies all)")
}
