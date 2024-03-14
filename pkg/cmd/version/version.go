package version

import (
	"fmt"
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
				fmt.Printf("Unable to check if a new version is available: %v\nYou can also view the releases at: https://github.com/SneaksAndData/snd-cli-go/releases\n", err)
			}
			return nil

		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("snd-cli version %s \n", snd.Version)
		},
	}

	return cmd
}
