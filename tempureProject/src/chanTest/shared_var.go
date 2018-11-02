package main

import "fmt"

type shared_var struct {
	reader chan int
	writer chan int
}

func shared_var_watchdog(v shared_var){
	go func() {
		var value int = 0
		for {
			select {
			case value = <- v.writer:
			case v.reader <- value:
			}
		}
	}()
}

func main(){
	v := shared_var{make(chan int), make(chan int)}
	shared_var_watchdog(v)

	fmt.Println(<-v.reader)
	v.writer <- 1
	fmt.Println(<-v.reader)

	defer fmt.Printf("end main\n")
}
