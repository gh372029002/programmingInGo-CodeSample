package main

import "fmt"

func main() {
	i := Minimum(4, 3, 8, 2, 9).(int)
	fmt.Printf("%T %v\n", i, i)
	f := Minimum(9.4, -5.4, 3.8, 17.0, 12.1, -3.1, 0.0).(float64)
	fmt.Printf("%T %v\n", f, f)
	s := Minimum("K", "X", "B", "C", "CC", "CA", "D", "M").(string)
	fmt.Printf("%T %q\n", s, s)
}

func Minimum(first interface{}, rest ...interface{}) interface{} {
	fmt.Println(rest)
	minimum := first
	for _, x := range rest {
		switch x := x.(type) {
		case int:
			if x < minimum.(int) {
				minimum = x
			}
		case float64:
			if x < minimum.(float64) {
				minimum = x
			}
		case string:
			if x < minimum.(string) {
				minimum = x
			}
		}
	}
	return minimum
}
