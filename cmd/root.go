package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "mally-cli",
	Short: "A CLI tool for creating short URLs and pastes directly from your terminal.",
	Long:  `Mally is a website that provides a collection of web services such as URL shorteners and pastebins.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
