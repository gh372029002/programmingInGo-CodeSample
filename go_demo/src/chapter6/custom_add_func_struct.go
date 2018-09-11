package main

import (
	"fmt"
	"strings"
)

type Part struct {
	Id   int
	Name string
}

func (part *Part) LowerCase() {
	part.Name = strings.ToLower(part.Name)
}
func (part *Part) UpperCase() {
	part.Name = strings.ToUpper(part.Name)
}

// 调用输出整个Part结构体时被调用.
func (part Part) String() string {
	return fmt.Sprintf("<<%d %q>>", part.Id, part.Name)
}

func (part Part) HasPrefix(prefix string) bool {
	return strings.HasPrefix(part.Name, prefix)
}
func main() {
	part := Part{5, "wrenth"}
	part.UpperCase()
	part.Id += 11
	fmt.Println("下面Println打印part的(String)...")
	fmt.Println(part, part.HasPrefix("W"))
}
