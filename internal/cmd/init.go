package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jaronline/blocksmith/internal/lib"
	"github.com/jaronline/blocksmith/internal/utils"
	"github.com/spf13/cobra"
)

var withDefaults bool
var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"create"},
	Short:   "Set up a new modpack.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var pkg *lib.Package

		_ = utils.ReturnsError(&pkg, &err, lib.GetDefaultPackage) &&
			!withDefaults &&
			utils.ReturnsError(&pkg.Name, &err, func() (string, error) {
				return reader("Fill in the name of the modpack", pkg.Name)
			}) &&
			utils.ReturnsError(&pkg.Version, &err, func() (string, error) {
				return reader("Fill in the version of the modpack", pkg.Version)
			})

		if err != nil {
			return fmt.Errorf("unable to initialize modpack: %w", err)
		}

		if err = pkg.Write(); err != nil {
			return fmt.Errorf("unable to write modpack: %w", err)
		}

		return nil
	},
}

func init() {
	initCmd.Flags().BoolVarP(&withDefaults, "yes", "y", false, fmt.Sprintf(
		"Automatically answer \"yes\" to any prompts that \"%s\" might print on the command line.", rootCmd.Name()))
	rootCmd.AddCommand(initCmd)
}

func reader(prompt string, defaultValue string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(fmt.Sprintf("%s (default %s): ", prompt, defaultValue))

	response, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}

	if len(response) == 0 {
		return defaultValue, nil
	}

	return string(response), nil
}
