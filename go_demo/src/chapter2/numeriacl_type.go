package main

import (
	"fmt"
	"math"
	"os"
)

// chapter 2.3
func main() {
	const factor = 3
	i := 2000
	i *= factor
	j := int16(20)
	i += int(j)
	k := uint8(0)
	// 缩小i小尺寸
	if _k, err := Uint8FromInt(i); err != nil {
		fmt.Println(err)
		os.Exit(0)
	} else {
		k = _k
	}
	fmt.Println(i, j, k)
}

//创建合适的转换函数
func Uint8FromInt(x int) (uint8, error) {
	if 0 <= x && x <= math.MaxUint8 {
		return uint8(x), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", x)
}
