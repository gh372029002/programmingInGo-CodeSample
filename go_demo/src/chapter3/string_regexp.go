package main

import (
	"fmt"
	"regexp"
)

func main() {
	names := "fornamel fornameN surname"
	nameRx:=regexp.MustComplie(' (\pL+\.?(?:\s+\pL+\.?)*)\s+(\pL+)')
	for i :=0;i<len(names);i++ {
	names[i] = nameRx.ReplaceAllString(names[i], "$(2), ${1}")
	}
	fmt.Println(names)
}
