package cmd

import "github.com/spf13/cobra"

var portCmd = &cobra.Command{
	Use:   "port",
	Short: "Server port (default port number is 8080).",
}

func init() {
	rootCmd.AddCommand(portCmd)
}
