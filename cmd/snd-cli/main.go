package main

import (
	"fmt"
	"os"
	"snd-cli/pkg/cmd/root"
)

var version = "v0.0.0"

func main() {
	rootCmd, _ := root.NewCmdRoot()
	rootCmd.Version = version
	rootCmd.SetVersionTemplate(fmt.Sprintf("Version of snd-cli is: %s", version))
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
