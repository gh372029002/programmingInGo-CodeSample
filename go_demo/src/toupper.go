package main

import (
	"fmt"
	"regexp"
	"strings"
)

var line string

func main() {
	line = "huwenfeng"
	wordRx := regexp.MustCompile("[A-Za-z]+")
	line = wordRx.ReplaceAllStringFunc(line,
		func(word string) string { return strings.ToUpper(word) })
	fmt.Println(line)

}
