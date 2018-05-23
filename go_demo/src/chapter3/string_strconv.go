package main

import (
	"fmt"
	"strconv"
)

func main() {
	// ParseBool
	for _, truth := range []string{"1", "t", "TRUE", "false", "F", "0", "5"} {
		if b, err := strconv.ParseBool(truth); err != nil {
			fmt.Printf("\n{%v}", err)
		} else {
			fmt.Print(b, " ")
		}
	}
	fmt.Println()

	//Other
	z, err := strconv.Atoi("71309")
	fmt.Printf("%8T ,%6v ,%v\n", z, z, err)
	s := strconv.FormatBool(z > 100)
	fmt.Println(s)
}
