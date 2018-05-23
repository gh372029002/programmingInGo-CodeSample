//Go 语言指针
package main

import "fmt"

func main() {
    var a int = 20
    var ip,_nil *int
    ip = &a
    fmt.Println(ip)
    fmt.Println(_nil)
    if _nil==nil {
        fmt.Println("_nil is nil")
    }
}