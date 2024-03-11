package main

import (
	"fmt"
	"strings"
)

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

func translator(word string) string {
	if len(word) <= 0 {
		return ""
	}

	var result strings.Builder

	for _, char := range strings.ToUpper(word) {
		varc := string(char)
		result.WriteString(translationTable[varc])
		result.WriteString(" ")
	}

	return result.String()
}

func main() {
	fmt.Println(translator("kimi"))
	fmt.Println(translator(""))
}
