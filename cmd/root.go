/*
Copyright © 2024 Peter Leung
*/
package cmd

import (
	"os"
	"webapp/internal/route"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Web Application",
	Long:  `Start a simple web application`,
	Run: func(cmd *cobra.Command, args []string) {
		route.StartServer()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
