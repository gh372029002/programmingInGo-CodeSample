package main

import (
	"../../goeg/src/stacker/stack"
	"fmt"
)

func main() {
	// 自定义数据集合类型
	var haystack stack.Stack
	haystack.Push("hay")
	haystack.Push(-15)
	haystack.Push([]string{"pin", "clip", "needle"})
	haystack.Push(81.52)
	for {
		item, err := haystack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}
