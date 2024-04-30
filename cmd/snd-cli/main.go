package main

import (
	"fmt"
	"github.com/spf13/cobra/doc"
	"log"
	"os"
	"snd-cli/pkg/cmd/root"
)

func main() {
	rootCmd, _ := root.NewCmdRoot()
	err := doc.GenMarkdownTree(rootCmd, "./docs")
	if err != nil {
		log.Fatal(err)
	}
	err = rootCmd.Execute()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
