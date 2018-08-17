package main

import "fmt"

func main() {
	counterA := createCounter(2)
	counterB := createCounter(102)
	for i := 0; i < 5; i++ {
		a := <-counterA
		fmt.Printf("(A->%d,B->%d)", a, <-counterB)
	}
	fmt.Println()
}

func createCounter(start int) chan int {
	next := make(chan int)
	go func(i int) {
		for {
			next <- i
			fmt.Printf("%s,%s\n", "@", "#")
			i++
		}
	}(start)
	return next
}
