package main

import (
	"fmt"
)

type Data struct {
	ID   int
	Name string
	Age  int
}

func init() {
	fmt.Printf("first%s", "is ok")
}

func main() {
	const allDone = 2
	doneCount := 0
	answera := make(chan int)
	answerb := make(chan int)
	defer func() {
		close(answera)
		close(answerb)
	}()
	done := make(chan bool)
	defer func() { close(done) }()
	data1 := Data{1, "ok", 11}
	data2 := Data{2, "ok", 12}
	go expensiveComputation(data1, answera, done)
	go expensiveComputation(data2, answerb, done)
	for doneCount != allDone {
		var which, result int
		select {
		case result = <-answera:
			which = 'a'
		case result = <-answerb:
			which = 'b'
		case <-done:
			doneCount++
		}
		if which != 0 {
			fmt.Printf("%c-%d", which, result)
		}
	}
	fmt.Println()
}
func expensiveComputation(data Data, answer chan int, done chan bool) {
	// 设置
	finished := false
	for !finished {
		// 计算
		answer <- data.Age //返回结果集 int
		data.Age += 1
		if data.Age >= 14 {
			finished = true
		}
	}
	done <- true
}
