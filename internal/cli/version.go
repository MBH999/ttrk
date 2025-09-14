package cli

import (
	"fmt"

	"github.com/MBH999/ttrk/pkg/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ttrk",
	Long:  "All software has versions. This is ttrk's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ttrk version:", version.GetVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}