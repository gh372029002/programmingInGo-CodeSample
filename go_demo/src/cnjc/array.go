//hello.go
package main

import (
	"fmt"
	"os"
	"strings"
)

/* 定义结构体 */
type Circle struct{
	radius float64
}

var balance [5] float32
var arr [5] int
var arrs [3][4] int

func main() {
	// 一维数组
	var i int
	for i=0;i<5;i++ {
		balance[i] = (float32)(i) + 10.5
		arr[i] = i
	}
	fmt.Println(balance)
	// 二维数组
	arrs[2][0] = 'a'  // 转换ASCII
	arrs[2][1] = 'b'
	arrs[2][2] = '1' 
	arrs[2][3] = 2
	fmt.Println(arrs)

	who := "World!"
	if len(os.Args) > 1 { /* os.Args[0]是"hello"或者"hello.exe" */
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello", who)

	a, b, _ := function_name("a", "b", 28)

	fmt.Println(&a, b)

	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	// 调用结构体类型method
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("Area of Circle(c1) = ", c1.getArea())
	fmt.Println(arr)
	// 传递数组到函数
	sum := getAverage(arr,4)
	fmt.Println(sum)
	

}

func function_name(name string, name1 string, age int)(int,int,int) {
	fmt.Println(name, name1, "age:", age)
	return 1, 2, 3
}

func getSequence() func() int{
	i:=0
	return func() int {
		i+=1
		return i
	}
}

// 该method 属于Circle 类型对象中的方法
func (c Circle) getArea() float64{
	// c.radius 即为Circle 类型对象中的属性
	return 3.14* c.radius * c.radius
}

//method 传递数组
func getAverage(arr [5]int, size int) float32{
	var i,sum int
	var avg float32

	for i=0;i<size;i++ {
		sum += arr[i]
	}

	avg = (float32)(sum) / (float32)(size)

	return avg
}
