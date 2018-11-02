package main

import (
	"log"
	"net"
)

var content = `HTTP/1.1 200 OK
Date: Sat, 29 Jul 2017 06:18:23 GMT
Content-Type: text/html
Connection: Keep-Alive
Server: BWS/1.1
X-UA-Compatible: IE=Edge,chrome=1
BDPAGETYPE: 3
Set-Cookie: BDSVRTM=0; path=/

<html>
<head  name="尹正杰" age="25">  <!--标签的开头,其和面跟的内容（name="尹正杰"）是标签的属性,其属性可以定义多个。-->
    <meta charset="UTF-8"/>     <!--指定页面编码，-->
    <meta http-equiv="refresh" content="30; Url=http://www.cnblogs.com/yinzhengjie/"> <!--这是做了一个界面的跳转，表示30s不运行的话就跳转到指定的URL-->
    <title>尹正杰的个人主页</title> <!--定义头部的标题-->
</head> <!--标签的结尾-->

<body>
<h1 style="color:red">尹正杰</h1>
<h1 style="color:green">hello golang</h1>

</body>
</html>
`

func handleConnection(conn net.Conn){
	conn.Write([]byte(content))
	conn.Close()
}

func main(){
	l,err := net.Listen("tcp", ":8888")
	if err!= nil{
		log.Println("error listen:", err)
		return
	}

	defer l.Close()
	log.Println("listen ok")

	for {
		conn, err := l.Accept()
		if err != nil{
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}
