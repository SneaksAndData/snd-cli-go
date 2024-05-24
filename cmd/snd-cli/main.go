package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"log"
	"os"
	"snd-cli/pkg/cmd/root"
)

var genDocs bool

func main() {
	rootCmd, _ := root.NewCmdRoot()

	// Add a persistent flag to the root command for generating documentation
	rootCmd.PersistentFlags().BoolVar(&genDocs, "gen-docs", false, "Generate Markdown documentation for all commands")
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if genDocs {
			err := doc.GenMarkdownTree(rootCmd, "./docs")
			if err != nil {
				log.Fatalf("Failed to generate docs: %v", err)
			} else {
				fmt.Println("Documentation generated in ./docs")
			}
			os.Exit(0) // Exit after generating docs
		}
	}
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
