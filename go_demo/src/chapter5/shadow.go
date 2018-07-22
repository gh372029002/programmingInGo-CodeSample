package main

import (
	"fmt"
	"strings"
)

func main() {
	shadow()
}

func shadow() (err error) {
	x, err := check1()
	if err != nil {
		return
	}
	if y, err := check2(x); err != nil {
		return
	} else {
		fmt.Println(y)
	}

	return

}

func check1() {
	return true
}

func check2(x) {
	if x != 1 {
		return nil
	}
	return true
}
