package main

import (
	"os"
	"snd-cli/pkg/cmd/root"
)

func main() {
	rootCmd, _ := root.NewCmdRoot()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
