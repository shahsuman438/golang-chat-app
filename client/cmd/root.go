package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "cobra",
		Short: "Chat Application CLI App",
		Long:  `Chat Application CLI app with postgres and cobra and socket.io`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
