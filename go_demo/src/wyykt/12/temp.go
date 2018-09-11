package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connect()
}
type PhoneConnecter struct {
	name string
	age  string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}
func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

func main() {
	var a USB
	a = PhoneConnecter{"PhoneConnecter", "21"}
	a.Connect()
	na := a.Name()
	fmt.Println(na)
}
