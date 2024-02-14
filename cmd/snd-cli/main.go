package main

import (
	"fmt"
	"os"
	"snd-cli/pkg/cmd/root"
)

func main() {
	rootCmd, _ := root.NewCmdRoot()
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
