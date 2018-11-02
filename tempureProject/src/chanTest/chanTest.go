package main

import (
	"fmt"
)

var complete1 chan int = make(chan int)
var complete2 chan int = make(chan int)

func loop1() {
	for i:=0; i < 100; i++{
		fmt.Printf("loop1:%d ", i)
	}

	complete1 <- 0
}

func loop2() {
	for i:=0; i < 100; i++{
		fmt.Printf("loop2:%d ", i)
	}

	complete2 <- 0
}

var ch chan int = make(chan int)

func foo(id int) { //id: 这个routine的标号
	ch <- id
}

func main() {
	// 开启5个routine
	for i := 0; i < 5; i++ {
		go foo(i)
	}

	// 取出信道中的数据
	for i := 0; i < 5; i++ {
		fmt.Print(<- ch)
	}
}

//func main(){
//	c, quit := make(chan int), make(chan int)
//
//	go func() {
//		c <- 1
//		quit <- 0
//	}()
//
//	<- c
//	<- quit // quit 等待数据的写
//}
