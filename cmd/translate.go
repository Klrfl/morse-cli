/*
Copyright © 2024 klrfl <efrayanglain@gmail.com>

*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var target string

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

		return fmt.Errorf("invalid translation target: ", args[0])
	},

	Run: func(cmd *cobra.Command, args []string) {
		if target == "morse" {
			fmt.Println(translateToMorseCode(args[0]))
		} else {
			fmt.Println(translateToPlainText(args[0]))
		}
	},
}

func translateToMorseCode(sentence string) string {
	if len(sentence) <= 0 {
		return ""
	}

	var result strings.Builder

	for _, char := range strings.ToUpper(sentence) {
		varc := string(char)
		result.WriteString(translationTable[varc])
		result.WriteString(" ")
	}

	return result.String()
}

var translationTable map[string]string = map[string]string{
	"A": ".-",
	"B": "-...",
	"C": "-.-.",
	"D": "-..",
	"E": ".",
	"F": "..-.",
	"G": "--.",
	"H": "....",
	"I": "..",
	"J": ".----",
	"K": "-.-",
	"L": ".-..",
	"M": "--",
	"N": "-.",
	"O": "---",
	"P": ".--.",
	"Q": "--.-",
	"R": ".-.",
	"S": "...",
	"T": "-",
	"U": "..-",
	"V": "...-",
	"W": ".--",
	"X": "-..-",
	"Y": "-.--",
	"Z": "--..",
	" ": "/",

	"0": "-----",
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
}

func init() {
	rootCmd.AddCommand(translateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// translateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	translateCmd.Flags().StringVarP(&target, "target", "t", "morse", "translation target")
}
