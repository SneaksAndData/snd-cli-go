package ml

import (
	"fmt"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmd/util"
	"snd-cli/pkg/shared/esd-client/crystal"
)

var cause, message, sasUri string

func NewCmdSubmit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit",
		Short: "Submit a result for a ML Algorithm run",
		RunE: func(cmd *cobra.Command, args []string) error {
			return submitRun()
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Specify the Crystal Job ID")
	cmd.Flags().StringVarP(&cause, "cause", "c", "", "Cause for submitting the result")
	cmd.Flags().StringVarP(&message, "message", "m", "", "Result message")
	cmd.Flags().StringVarP(&sasUri, "sas_uri", "u", "", "Sas Uri")

	return cmd
}

func submitRun() error {
	// TODO: url is wrong (needs to be receiver url)
	url := fmt.Sprintf("https://crystal.%s.sneaksanddata.com", env)
	var crystalConn crystal.Connector
	crystalConn = crystal.NewConnector(url, "", "v1.2")
	token, err := util.ReadToken()
	if err != nil {
		return err
	}
	resp, err := crystalConn.SubmitResult(id, algorithm, cause, message, sasUri, token)
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}
