package ml

import (
	"fmt"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmd/util"
	"snd-cli/pkg/shared/esd-client/crystal"
)

var id string

func NewCmdGet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get the result for a ML Algorithm run",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getRun()
		},
	}
	cmd.Flags().StringVarP(&id, "id", "i", "", "Specify the Crystal Job ID")
	return cmd
}

func getRun() error {
	url := fmt.Sprintf("https://crystal.%s.sneaksanddata.com", env)
	var crystalConn crystal.Connector
	crystalConn = crystal.NewConnector(url, "", "v1.2")
	token, err := util.ReadToken()
	if err != nil {
		return err
	}
	resp, err := crystalConn.RetrieveRun(id, algorithm, token)
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}
