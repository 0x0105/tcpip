package commands

import (
	"github.com/spf13/cobra"
)

// root cmd
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goim",
	Short: "An instant message app implemented by Go.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(
		newConnectCmd().getCommand(),
	)
	cobra.CheckErr(rootCmd.Execute())
}
