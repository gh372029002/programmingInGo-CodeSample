//火箭发射
package main

import (
 "fmt"
 "time"
 "os"
)
var ch chan int

func listenstdin(){
	// 监控标准输入
	n,err  := os.Stdin.Read([512]byte)	//阻塞读
	if err!=nil{
		fmt.Println('print err')
	}
	ch<-1
	
}

func lanuch() {
 fmt.Println("发射")
}

func main() {
	ch = make(ch,3)
 tk := time.NewTicker(time.Second) //定时器,每隔1s发送内容到管道
 for i := 5; i > 0; i-- {
	select {
		case <-tk.C: //读管道
			fmt.Println(i)
		case <- ch:
			fmt.Println('结束发射')
			return			
	}
	lanuch()
  fmt.Println(i)
 }
