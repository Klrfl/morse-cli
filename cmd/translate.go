/*
Copyright © 2024 klrfl <efrayanglain@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/klrfl/morse-translator/pkg/translate"
	"github.com/spf13/cobra"
)

var (
	target string
	input  string
)

// translateCmd represents the translate command
var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "The translate command translates either to morse code or to plain text.",
	Long: `The translate command translates either to morse code or to plain text.
  Use the flags -t (or --target) to specify target (morse or plain). For example:

  --target m takes text and translates it into morse code
  --target p takes morse and translates it into plain text
  `,

	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
			return fmt.Errorf("too much arguments")
		}

		options := []string{"morse", "m", "plain", "p"}

		for _, value := range options {
			if target == value {
				return nil
			}
		}

		return fmt.Errorf("invalid translation target: %s", target)
	},

	Run: func(cmd *cobra.Command, args []string) {
		switch target {
		case "plain":
		case "p":
			fmt.Println(translate.TranslateToPlainText(input))
		default:
			fmt.Println(translate.TranslateToMorseCode(input))
		}
	},
}

func init() {
	rootCmd.AddCommand(translateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// translateCmd.PersistentFlags().String("foo", "", "A help for foo")
	translateCmd.Flags().StringVarP(&target, "target", "t", "morse", "translation target")
	translateCmd.Flags().StringVarP(&input, "input", "i", "", "string to be translated")
	translateCmd.MarkFlagRequired("input")
}
