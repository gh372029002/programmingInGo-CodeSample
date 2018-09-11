package main

import "fmt"

func main() {
	i := f(1)
	fmt.Println(i(2))
}
func f(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}
