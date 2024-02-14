package spark

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdLogs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logs",
		Short: "Get logs from a Spark Job",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("logs called")
		},
	}

	cmd.Flags().StringP("id", "i", "", "Beast Job ID")
	cmd.Flags().StringP("trim-logs", "t", "", "Trims log to anything after STDOUT")

	return cmd
}
