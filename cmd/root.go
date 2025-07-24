/*
Copyright © 2025 github.com/dayiamin/imgtool
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "imgtool",
	Short: "A CLI Tool for Image Compression",
	Long:  `imgtool is a CLI tool to compress and manage images efficiently.`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.imgtool.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.SetUsageTemplate(`Usage:
  imgtool [command]

Available Commands:
  compress    Compress images from a file or folder
  help        Help about any command

Flags:
  -h, --help     help for imgtool

Use "imgtool [command] --help" for more information about a command.

Sample:
	"imgtool compress --help"
	"imgtool compress -i ./inputfolder -o ./output -f original  -s "_suffix" -q 80
	
	`)

	
}
