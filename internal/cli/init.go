package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jaronline/blocksmith/internal/app"
	"github.com/spf13/cobra"
)

var withDefaults bool
var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"create"},
	Short:   "Set up a new modpack.",
	Run: func(cmd *cobra.Command, args []string) {
		name := reader("Fill in the name of the modpack", "greg")
		version := reader("Fill in the version of the modpack", "1.0.0")

		app.CreatePackage(&app.Package{
			Name:    name,
			Version: version,
		})
	},
}

func init() {
	initCmd.Flags().BoolVarP(&withDefaults, "yes", "y", false, fmt.Sprintf(
		"Automatically answer \"yes\" to any prompts that \"%s\" might print on the command line.", rootCmd.Name()))
	rootCmd.AddCommand(initCmd)
}

func reader(prompt string, defaultValue string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(fmt.Sprintf("%s (default %s): ", prompt, defaultValue))

	response, _, err := reader.ReadLine()
	if err != nil {
		os.Exit(1)
	}

	if len(response) == 0 {
		return defaultValue
	}

	return string(response)
}
