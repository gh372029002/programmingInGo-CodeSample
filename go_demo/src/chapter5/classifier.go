package main

import (
	"fmt"
)

func main() {
	classifier(5, "ZIP", -17.9, nil, true, complex(1, 1))
}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%d is a float64\n", i)
		case int, int8, int16, int32, int64:
			fmt.Printf("param #%d is an int\n", i)
		case uint, uint8, uint16, uint32, uint64:
			fmt.Printf("param #%d is a unsigned int\n", i)
		case nil:
			fmt.Printf("param #%d is nil\n", i)
		case string:
			fmt.Printf("param #%d is string\n", i)
		default:
			fmt.Printf("param #%d's type is unknow\n", i)
		}
	}
}
