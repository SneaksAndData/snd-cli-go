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
	versionTemplate := fmt.Sprintf("snd-cli version %s \n", cmd.Version)
	rootCmd.SetVersionTemplate(versionTemplate)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
