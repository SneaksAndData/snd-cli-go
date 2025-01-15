package version

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	snd "snd-cli/cmd"
	"snd-cli/pkg/cmd/util/version"
)

func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of the SnD CLI.",
		Long:  `All software has versions. This is SnD CLI's version.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			err := version.CheckIfNewVersionIsAvailable()
			if err != nil {
				pterm.DefaultBasicText.Println("Unable to check if a new version is available: You can view the releases at: https://github.com/SneaksAndData/snd-cli-go/releases")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			pterm.DefaultBasicText.Println(
				pterm.Sprintf("snd-cli version %s", snd.Version),
			)
		},
	}

	return cmd
}
