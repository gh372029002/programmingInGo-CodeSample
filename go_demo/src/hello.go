//hello.go
package main
import(
	"fmt"
	"os"
	"strings"
)
func main() {
	who := "World!"
	if len(os.Args) > 1 {/* os.Args[0]是"hello"或者"hello.exe" */
	   who = strings.Join(os.Args[1:], " ") 
	}
    fmt.Println("Hello", who)
}

  