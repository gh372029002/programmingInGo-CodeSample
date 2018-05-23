package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// TrimSpace start
	str := " abcdef 123456 "
	fmt.Printf("|%s|\n", SimpleSimplifyWhitespace(str))
	// TrimSpace END
	//
	// buffer write 实现trimSpace
	var buffer bytes.Buffer
	skip := true
	for _, char := range str {
		if unicode.IsSpace(char) {
			if !skip {
				buffer.WriteRune(' ')
				skip = true
			} else {
				buffer.WriteRune(char)
				skip = false
			}
		}

	}
	str = buffer.String()
	if skip && len(str) > 0 {
		str = str[:len(str)-1]
	}
	fmt.Printf(str)
	fmt.Println()
	// END
}
func SimpleSimplifyWhitespace(s string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(s)), " ")
}
