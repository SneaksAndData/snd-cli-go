package main

import (
	"fmt"
	"os"
	"snd-cli"
	"snd-cli/pkg/cmd/root"
)

func main() {
	rootCmd, _ := root.NewCmdRoot()
	rootCmd.Version = snd_cli_go.Version
	rootCmd.SetVersionTemplate(fmt.Sprintf("snd-cli version %s", snd_cli_go.Version))
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
