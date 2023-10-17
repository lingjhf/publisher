package cmd

import "github.com/spf13/cobra"

var hostCmd = &cobra.Command{
	Use:   "host",
	
	Short: "Server bind address (default bind address is localhost).",
}

func init() {
	rootCmd.AddCommand(hostCmd)
}
