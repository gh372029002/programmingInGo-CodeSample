package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "naive"
	fmt.Println("saaa")
	fmt.Println(s[0])
	fmt.Println(s[len(s)-1])

	// string test
	line := "røde og gule sløjfer"
	//第一个空格前内容
	i := strings.Index(line, " ")
	firstWord := line[:i]
	fmt.Println(firstWord)
	//最后一个空格后内容
	j := strings.LastIndex(line, " ")
	lastWord := line[j+1:]
	fmt.Println(lastWord)
	line_a := "rå tort\u2028vær"
	i_a := strings.IndexFunc(line, unicode.IsSpace)
	fmt.Println(line_a, i_a)
	//....write

	//为调试格式化
	p := []string{"-83.40", "71.60"}
	fmt.Printf("%T|%v|%#v|\n", p, p, p)
	splitStr := "1-2-3-4-5-6--7-8-9"
	for _, name := range strings.Split(splitStr, "-") {
		fmt.Printf("%s|", name)
	}
	fmt.Println()
	for _, name := range strings.SplitAfterN(splitStr, "-", 6) {
		fmt.Printf("%s|", name)
	}
}
