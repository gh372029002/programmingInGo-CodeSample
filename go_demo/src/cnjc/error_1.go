//Go 错误处理
package main

import (
	"errors"
	"fmt"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// 实现
	return 1, nil
}

func main() {

	result, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
}
