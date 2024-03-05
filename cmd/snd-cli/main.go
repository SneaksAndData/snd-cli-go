package main

import (
	"fmt"
	"os"
	"snd-cli/cmd"
	"snd-cli/pkg/cmd/root"
)

func main() {
	rootCmd, _ := root.NewCmdRoot()
	rootCmd.Version = cmd.Version
	rootCmd.SetVersionTemplate(fmt.Sprintf("snd-cli version %s", cmd.Version))
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
