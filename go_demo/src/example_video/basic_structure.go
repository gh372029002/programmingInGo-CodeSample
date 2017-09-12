// 当前程序的包名
package main

// 导入其他包
import (
	std "fmt"
)

// 导入自定义包
// import p1 " importPackage"

// 常量的定义
const (
	PI   = 3.14
	Year = 1
)

// 全局变量的声明与赋值π`
var (
	name  = "gopher"
	name1 = "name 1"
)

// 一般类型声明
type newType int

// 结构的声明
type gopher struct{}

// 接口的声明
type golang interface{}

// 由 main 函数作为程序入口点启动
func main() {
	std.Println("Hello world!你好 , 世界!")
	std.Println(name1)
	var1 := iota(PI)
	std.Pirntln(var1)
}
