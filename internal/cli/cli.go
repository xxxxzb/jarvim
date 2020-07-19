// Package cli provides ...
package cli

import (
	"fmt"
	"os"

	"github.com/glepnir/jarvis/internal/logic"
	"github.com/spf13/cobra"
)

var (
	cli_version = "0.0.1"
	version     bool
	genConfig   bool
)

var rootCmd = &cobra.Command{
	Use:   "jarvis",
	Short: "Jarvis is a cli tool to generate a module vim configruation which like a pro",
	Long:  "Jarvis is a cli tool to generate a module vim configruation which like a pro",
	RunE: func(cmd *cobra.Command, args []string) error {
		if version {
			return printVersion()
		}
		if genConfig {
			return logic.RunLogic()
		}
		return cmd.Help()
	},
}

func init() {
	cobra.OnInitialize()
	rootCmd.Flags().BoolVarP(&version, "Version", "v", false, "show current version of CLI")
	rootCmd.Flags().BoolVarP(&genConfig, "Generate vim config", "g", false, "generate new configuration")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printVersion() error {
	fmt.Println("Version: ", cli_version)
	return nil
}