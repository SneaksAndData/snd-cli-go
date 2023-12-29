package ml

import (
	"fmt"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmd/util"
	"snd-cli/pkg/shared/esd-client/crystal"
)

var payload, tag string

func NewCmdRun() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run a ML Algorithm",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runRun()
		},
	}

	cmd.Flags().StringVarP(&payload, "payload", "p", "", "Path to the payload JSON file")
	cmd.Flags().StringVarP(&tag, "tag", "t", "", " Client-side submission identifier")
	return cmd
}

func runRun() error {
	url := fmt.Sprintf("https://crystal.%s.sneaksanddata.com", env)
	var crystalConn crystal.Connector
	crystalConn = crystal.NewConnector(url, "", "v1.2")
	token, err := util.ReadToken()
	if err != nil {
		return err
	}
	body, err := util.ReadJSONFile(payload)
	//TODO: conf := body["CustomConfiguration"]
	resp, err := crystalConn.CreateRun(algorithm, body["AlgorithmParameters"], crystal.AlgorithmConfiguration{}, tag, token)
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}
