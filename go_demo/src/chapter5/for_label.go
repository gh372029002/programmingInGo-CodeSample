package main

import "fmt"

// 标签是个后面带一个冒号的标识符
func main() {
	table := [2][3]string{{"a", "b", "c"}, {"xxx", "xx", "x"}}
	// 循环语法2 = 循环语法1
	func1(table)
	func2(table)
}

func func1(table [2][3]string) {
	found := false
	for row := range table {
		for column := range table[row] {
			if table[row][column] == "x" {
				fmt.Println("跳出循环func1")
				found = true
				break
			}
			if found {
				break
			}

		}
	}
}

func func2(table [2][3]string) {
ABC:
	for row := range table {
		for column := range table[row] {
			if table[row][column] == "x" {
				fmt.Println("跳出循环func2") // found = true
				break ABC
			}
		}
	}
}
