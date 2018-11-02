package main

import (
	"fmt"
	"sync"
)

func main() {
	people := []string{"Anna", "Bob", "Cody", "Dave", "Eva", "jason"}
	match := make(chan string, 1) // 给未匹配的元素预留空间
	wg := new(sync.WaitGroup)
	for _, name := range people {
		wg.Add(1)
		go Seek(name, match, wg)
	}
	wg.Wait()
	select {
	case name := <-match:
		fmt.Printf("No one received %s’s message.\n", name)
	default:
		// 没有待处理的发送操作.
	}
}

// 寻求发送或接收匹配上名称名称的通道,并在完成后通知等待组.
func Seek(name string, match chan string, wg *sync.WaitGroup) {
	select {
	case peer := <-match:
		fmt.Printf("%s received a message from %s.\n", name, peer)
	case match <- name:
		// 等待其他人接受消息.
	}
	wg.Done()
}
