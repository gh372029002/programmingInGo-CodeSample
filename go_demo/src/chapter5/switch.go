package main

import (
	"fmt"
	"path"
)

func main() {
	/*
		switch optionalStatement; optionalExpression{
			case expressionList1:  block1
			..
			case expressionListN:blockN
			default: blockD
		}
	*/
	// case 1
	res := BoundedInt(3, 0, 5)
	fmt.Println(res)
	// case 2
	file := "d:/abc.zip"
	res = getSuffix(file)
	fmt.Println(res)

}

func BoundedInt(minimum, value, maximum int) int {

	switch {
	case value < minimum:
		return minimum
	case value > maximum:
		return maximum
	default:
		return value
	}
	panic("unreachable")
}

//获取文件后缀
func getSuffix(file string) string {
	fv := Suffix(file)
	switch fv { //经典用法
	case ".gz":
		fmt.Println("@1")
	case ".tar", ".tar.gz", ".tgz":
		fmt.Println("@2")
	case ".zip":
		fmt.Println("@3")
	default:
		fmt.Println("@4")
	}
	return value
}

// 获取文件后缀
func Suffix(file string) string {
	return path.Ext(file) //获取文件后缀
}
