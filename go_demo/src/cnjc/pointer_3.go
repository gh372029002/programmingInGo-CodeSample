//Go 语言指向指针的指针
package main

import (
    "fmt"
)

func main(){
    var a int
    var ptr *int
    var pptr **int

    a = 3000
    /*
    指针 ptr 地址
     */
    ptr = &a
    /*
    指向指针 ptr 的地址
     */
    pptr = &ptr

    /*
    获取pptr的值
     */
    fmt.Printf("变量a = %d\n", a)
    fmt.Printf("指针变量 *ptr = %d\n", *ptr)
    fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}