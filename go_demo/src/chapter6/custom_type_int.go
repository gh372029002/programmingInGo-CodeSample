package main

import "fmt"

// 基于整型的自定义类型,支持三个方法
type Count int

var count Count

func (count *Count) Increment()  { *count++ }
func (count *Count) Decrement()  { *count-- }
func (count Count) IsZero() bool { return count == 0 }

func main() {
	count = 3
	i := int(count)
	count.Increment()
	j := int(count)
	count.Decrement()
	k := int(count)
	fmt.Println(count, i, j, k, count.IsZero())
}
