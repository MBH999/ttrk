package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ttrk",
	Short: "ttrk is a CLI tool",
	Long:  "ttrk is a command line interface tool for various operations.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to ttrk CLI!")
		fmt.Println("Use 'ttrk --help' for available commands.")
	},
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add global flags here if needed
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ttrk.yaml)")
}