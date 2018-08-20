package main

import (
	"bytes"
	"fmt"
	"strings"
)

type memoizeFunction func(int, ...int) interface{}

var Fibonacci memoizeFunction

var RomanForDecimal memoizeFunction

func init() {
	// test1
	Fibonacci = Memoize(func(x int, xs ...int) interface{} {
		if x < 2 {
			return x
		}
		return Fibonacci(x-1).(int) + Fibonacci(x-2).(int)
	})

	//test2
	decimals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	RomanForDecimal = Memoize(func(x int, xs ...int) interface{} {
		if x < 0 || x > 3999 {
			panic("RomanForDecimal() only handles integers [0,3999]")
		}
		var buffer bytes.Buffer
		for i, decimal := range decimals {
			remainder := x / decimal
			x %= decimal
			if remainder > 0 {
				buffer.WriteString(strings.Repeat(romans[i], remainder))
			}
		}
		return buffer.String()
	})

}

func Memoize(function memoizeFunction) memoizeFunction {
	fmt.Println("222")
	cache := make(map[string]interface{})
	return func(x int, xs ...int) interface{} {
		key := fmt.Sprint(x)
		for _, i := range xs {
			key += fmt.Sprintf(",%d", i)
		}
		if value, found := cache[key]; found {
			return value
		}
		value := function(x, xs...) //疑问?
		cache[key] = value
		return value
	}
}

func main() {
	fmt.Println("Fibonacci(45) =", Fibonacci(46, 42, 40).(int), Fibonacci(46).(int))
}
