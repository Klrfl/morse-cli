package translate

import "strings"

func TranslateToMorseCode(sentence string) string {
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

func TranslateToPlainText(sentence string) string {
	if len(sentence) <= 0 {
		return ""
	}

	var result strings.Builder

	inputSentence := strings.Split(sentence, "/")
	for _, morseWord := range inputSentence {
		for _, morseChar := range strings.Split(strings.Trim(morseWord, " "), " ") {
			for letter, code := range translationTable {
				if code == morseChar {
					result.WriteString(letter)
				}
			}
		}
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

	"&": ".-...",
	"'": ".----.",
	"@": ".--.-.",
	")": "-.--.-",
	"(": "-.--.",
	":": "---...",
	",": "--..--",
	"=": "-...-",
	"!": "-.-.--",
	".": ".-.-.-",
	"-": "-....-",
	"×": "-..-",
	//TODO: add % sign
	"+":  ".-.-.",
	"\"": ".-..-.",
	"?":  "..--..",
	"/":  "-..-.",
}
