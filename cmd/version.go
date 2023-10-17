package cmd

import "github.com/spf13/cobra"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version.",
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
