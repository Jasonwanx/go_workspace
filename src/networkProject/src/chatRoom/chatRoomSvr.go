package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

var globalRoom *Room = NewRoom()

type Room struct {
	users map[string]net.Conn
}

func NewRoom() *Room{
	return &Room{
		users:make(map[string]net.Conn),
	}
}

func (r *Room) Join(user string, conn net.Conn){
	if _,ok := r.users[user]; ok{
		r.Leave(user)
	}

	r.users[user] = conn
	fmt.Println(user, " login in!")
	conn.Write([]byte(user + ":join in chatroom!\n"))
}

func (r *Room) Leave(user string){
	conn, ok := r.users[user]
	if !ok{
		fmt.Println(user," not exists!")
		return
	}

	conn.Close()
	delete(r.users, user)
	fmt.Println(user, " leave")
}

func (r *Room) Broadcast(who string, msg string){
	timeInfo := time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006")
	toSend := fmt.Sprintf("%v %s:%s\n", timeInfo, who, msg)

	for user,conn := range r.users{
		if user == who{
			continue
		}

		conn.Write([]byte(toSend))
	}
}

func handleConn(conn net.Conn){
	defer conn.Close()

	r := bufio.NewReader(conn)
	line,err := r.ReadString('\n')
	if err != nil{
		fmt.Println(err)
		return
	}

	line = strings.TrimSpace(line)
	fields := strings.Fields(line)
	if len(fields) != 2{
		conn.Write([]byte("unvalid username and password, enforce to quit!\n"))
		return
	}

	user := fields[0]
	password := fields[1]
	if password != "123"{
		conn.Write([]byte("unvalid password, enforce to quit!"))
		return
	}

	globalRoom.Join(user, conn)
	globalRoom.Broadcast("System", fmt.Sprintf("%s join chatroom", user))
	for ; ;  {
		conn.Write([]byte("press enter to send message:>>>>"))
		line,err := r.ReadString('\n')
		if err != nil{
			break
		}

		line = strings.TrimSpace(line)
		fmt.Println(user, line)
		globalRoom.Broadcast(user, line)
	}

	globalRoom.Broadcast("System", fmt.Sprintf("%s leave room", user))
	globalRoom.Leave(user)
}

func main(){
	listener,err := net.Listen("tcp", ":8888")
	if err != nil{
		log.Fatal(err)
	}

	defer listener.Close()

	for{
		conn,err := listener.Accept()
		if err != nil{
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}