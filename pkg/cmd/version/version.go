package version

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	snd "snd-cli/cmd"
)

func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of the SnD CLI.",
		Long:  `All software has versions. This is SnD CLI's version.`,
		Run: func(cmd *cobra.Command, args []string) {
			pterm.DefaultBasicText.Println(
				pterm.Sprintf("snd-cli version %s", snd.Version),
			)
		},
	}

	return cmd
}
