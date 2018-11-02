package main

import (
	"log"
	"net"
)

func main(){
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8888")
	if err!= nil {
		log.Println("dial error:", err)
		return
	}

	defer conn.Close()

	hello := "hello"
	buf := []byte(hello)

	n, err := conn.Write(buf)
	if err != nil{
		log.Println("Write failed:", err)
		return
	}else{
		log.Println("Write ok:", n)
	}

	log.Println("dial ok")
}