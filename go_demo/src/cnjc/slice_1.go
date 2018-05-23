// 定义切片
package main

import "fmt"

//未指定大小的数组的切片
// var identifier []type

// 使用make()函数创建切片
// var slice1 []type = make([]type, len)

// 简写
// slice2 :=make([]type, len)

// 指定容量
// make([]T, length, capacity)

func main() {

	s := []int{1, 2, 3}
	fmt.Println(s)
	arr := []int{1, 2, 3, 4}
	fmt.Println(arr)
	s1 := arr[:]
	fmt.Println(s1)
	s2 := arr[1:2]
	fmt.Println(s2)
	s3 := arr[2:]
	fmt.Println(s3)
	s4 := arr[:3]
	fmt.Println(s4)
	printSlice(s4)
	s5 := s1[1:2]
	s1[1] = 6
	fmt.Println(s1)
	fmt.Println(s5)

	s6 := make([]int, 4, 5)
	s6[1] = 1
	s6[3] = 3
	s6[2] = 2
	s6[0] = 0

	fmt.Println(s6)
	printSlice(s6)
	// 空切片
	var s7 []int
	printSlice(s7)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
