package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// go run client.go zxf
// go run server.go

func main() {
	if len(os.Args) != 2 {
		fmt.Println("请输入用户名：", len(os.Args))
		return
	}
	name := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	connR, err := net.DialTCP("tcp", nil, tcpAddr) // 读
	connW, err := net.DialTCP("tcp", nil, tcpAddr) // 写
	defer connR.Close()
	defer connW.Close()

	connW.Write([]byte("w|" + name))
	connR.Write([]byte("r|" + name))
	go func() {
		for {
			var str = make([]byte, 1024)
			_, err := connR.Read(str)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("==> %s", str)
		}

	}()
	for {
		var msg string
		inputReader := bufio.NewReader(os.Stdin) //创建一个读取器，并将其与标准输入绑定。
		msg, err = inputReader.ReadString('\n')
		msg = strings.Trim(msg, " ")
		msg = strings.Trim(msg, "\n")
		if msg != "" {
			msg = name + "|" + msg
			connW.Write([]byte(msg))
		}
	}

}
