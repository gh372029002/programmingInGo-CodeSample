package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//Args[0]存放程序的名字
	fmt.Println(filepath.Base(os.Args[0]))
	//‘内容’解ASCII
	fmt.Println('3' - '0')
}
