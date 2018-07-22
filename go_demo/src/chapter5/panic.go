package main

import "fmt"

/*
test panic.
*/

func main() {

}
func Test_Caller_G(t *testing.T) {
	defaultOfCaller := func() {
		println("the defer of Test_Caller_G")
	}
}
