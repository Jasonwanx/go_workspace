package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

var(
	cmd string
	file_name string
)

func handleConnection(conn net.Conn){
	conn.Write([]byte("Welcome to ftp server!"))
	reader := bufio.NewReader(conn)

	for{
		line, err := reader.ReadString('\n')
		if err == io.EOF{
			conn.Close()
		}

		fmt.Println(line)
		line = strings.TrimSpace(line)
		if len(line) == 0{
			continue
		}

		cmd = strings.Fields(line)[0]
		if len(strings.Fields(line)) > 1{
			file_name = strings.Fields(line)[1]
		}

		pwd, err := os.Getwd()
		if err != nil{
			panic("获取路径出错")
		}

		file_name = filepath.Join(pwd, file_name)
		switch cmd {
		case "GET", "get":
			f, err := os.Open(file_name)
			if err != nil{
				fmt.Println("打开文件失败：",file_name)
				continue
			}

			defer f.Close()
			buf,err := ioutil.ReadAll(f)
			if err != nil{
				log.Println(err)
				return
			}

			conn.Write(buf)
		case "PUSH", "push":
			fmt.Println("上传文件的语句")
			conn.Write([]byte("上传文件的命令\n"))
		case "EXIT", "exit":
			conn.Close()
		default:
			fmt.Println("您输入的命令无效")
			conn.Write([]byte("您输入的命令无效!\n"))
		}
	}
}

func main(){
	listener, err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Fatal(err)
	}

	defer listener.Close()
	log.Println("listen ok!")

	for{
		conn, err := listener.Accept()
		if err != nil{
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}


