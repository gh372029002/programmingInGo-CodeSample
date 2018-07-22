// 递归阶乘
package main

import "fmt"

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func
func main() {
	var i uint64 = 3
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(i))
}
