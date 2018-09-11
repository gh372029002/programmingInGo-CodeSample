package main

import (
	"fmt"
	"strings"
	"unicode"
)

type RuneForRuneFunc func(rune) rune

var removePunctuation RuneForRuneFunc

func processPhrases(phrases []string, function RuneForRuneFunc) {
	for _, phrase := range phrases {
		fmt.Println(strings.Map(function, phrase))
	}
}

func main() {
	phrases := []string{"Day; dusk, and night.", "All day long"}
	removePunctuation = func(char rune) rune {
		if unicode.Is(unicode.Terminal_Punctuation, char) {
			return -1
		}
		return char
	}
	processPhrases(phrases, removePunctuation)
}
