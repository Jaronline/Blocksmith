package cmd

import (
	"os"

	"github.com/jaronline/blocksmith/internal/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "blocksmith",
	Short:   "Command-line tool for managing Minecraft modpacks",
	Version: "0.1.0",
	Run: func(cmd *cobra.Command, args []string) {
		ui.StartTUI()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
