package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"strings"
)

var user = make(map[string]map[string]*net.TCPConn)

func addUser(name string, conn *net.TCPConn, conntype string) {

	if _, ok := user[name]; !ok {
		user[name] = make(map[string]*net.TCPConn)
		str := name + "上线了！！"
		addMessage(name, str)
	}
	user[name][conntype] = conn
}

func deleteUser(name string) {
	if _, ok := user[name]; !ok {
		str := name + "同学下线了！！"
		addMessage(name, str)
	}
}

func addMessage(name, msg string) {
	msg = name + ": " + msg + "\n"
	fmt.Println(user)
	for i, _ := range user {
		fmt.Println("message:", msg, i, name, len(i), len(name))
		if i != name {

			n, err := user[i]["read"].Write([]byte(msg))
			if err != nil {
				log.Println(err)
				deleteUser(name)
				return
			}
			fmt.Printf("send %d bytes to %s %s\n", n, i, name)
		}
	}
}

func main() {
	address := net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 9999,
	}
	listener, err := net.ListenTCP("tcp4", &address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("远程地址：", conn.RemoteAddr())

		go func() {
			for {
				var res = make([]byte, 1024)
				_, err := conn.Read(res)
				fmt.Println(string(res))
				if err != nil {
					log.Fatal(err)
				}
				index := bytes.IndexByte(res, 0)
				res = res[0:index]
				arr := strings.Split(string(res), "|")
				if string(res[0]) == "w" {
					name := strings.Trim(arr[1], " ")

					addUser(name, conn, "write")
				} else if string(res[0]) == "r" {

					name := strings.Trim(arr[1], " ")
					addUser(name, conn, "read")
				} else {
					name := strings.Trim(arr[0], " ")
					msg := arr[1]
					go addMessage(name, msg)
				}
			}
		}()

	}
}
