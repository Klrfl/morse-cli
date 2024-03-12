/*
Copyright © 2024 klrfl <efrayanglain@gmail.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "morse-translator",
	Short: "A small CLI to translate to and from morse code",
	Long: `
morse-translator translates your morse code to plain text and back. for
example:
$ morse-translator translate -i '-... .- -... -.-- / -.. --- -. .----. - / .... ..- .-. - / -- .'

will ouput BABY DON'T HURT ME`,
	Version: version,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
