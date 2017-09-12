package main

import (
	"fmt"
	"strconv"
)

const s0 = "123"

const (
	s1 = 'A'
	s2 = "1"
	s3 = len(s2)
	s4 = len(s0)
)
const (
	m1 = 1
	m2
	m3 = iota
	m4
)

// 类型别名
type (
	byte int8
	rune int32
	文本   string
)

func main() {
	//语法-运算符
	fmt.Println(-1 ^ 2)
	fmt.Println(!false)
	fmt.Println(1 << 10 << 10 >> 10)
	c1 := 0
	if c1 > 0 && (10/c1) > 1 {
		fmt.Println("ok")
	}

	return
	// 语法-常量的使用
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println("-------------------------")
	// 语法-iota
	fmt.Println(m3)
	fmt.Println(m4)
	fmt.Println("-------------------------")
	var a 文本
	a = "中文别名"
	fmt.Println(a)

	var b rune
	b = '2'
	fmt.Println(b)
	c := 1
	fmt.Println(c)
	A, _, C := 1, 2, 3 //_,表示忽略值为2的参数
	fmt.Println(A)
	// fmt.Println(B)
	fmt.Println(C)
	//类型转换
	var aa float32 = 1.1
	fmt.Println(aa)
	bb := int(aa)
	fmt.Println(bb)
	// var aa [12]byte

	var a1 int = 65
	fmt.Println(a1)
	b1 := strconv.Itoa(a1)
	fmt.Println(b1)
}
