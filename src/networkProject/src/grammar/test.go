package main

import (
	"fmt"
)

type Test struct {
	key 	string
	value	int32
}

func main()  {
	pointer := &Test{"123", 99}
	fmt.Printf("%v", pointer)

}