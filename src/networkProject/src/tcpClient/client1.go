package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main(){
	addr := "www.baidu.com:80"
	conn, err := net.Dial("tcp", addr)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("访问公网IP地址是：", conn.RemoteAddr().String())
	fmt.Printf("客户端连接的地址及端口:%v\n", conn.LocalAddr())

	n, err := conn.Write([]byte("Get / HTTP/1.1\r\n\r\n"))
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("向服务端发送的数据长度是：",n)

	buf := make([]byte, 1024)

	for {
		n, err = conn.Read(buf)

		if err == io.EOF{
			conn.Close()
		}
		fmt.Println(string(buf[:n]))
	}

	fmt.Println(string(buf[:n]))
}