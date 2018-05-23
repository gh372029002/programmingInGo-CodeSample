package main

import (
	"fmt"
	"strings"
)

func main() {
	asciiOnly := func(char rune) rune {
		if char > 127 {
			return '?'
		} else {
			return char
		}
	}
	fmt.Println(strings.Map(asciiOnly, "JerØmesterreich"))

}
