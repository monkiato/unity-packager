/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "unity-packager",
	Short: "Create a .unitypackage file without Unity",
	Long: `Create a .unitypackage file without Unity`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var cachedir string

func init() {
	home, _ := os.UserHomeDir()
	rootCmd.PersistentFlags().StringVar(&cachedir, "cachedir", home+"/.unity-packager", "directory used for application cache (default is $HOME/.unity-packager)")

}
