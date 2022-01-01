package api

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

// Client 客户端
func Client() {
	//建立与服务器的联系
	coon, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}

	defer coon.Close()
	//获取服务端传来的消息
	go ioCopy(os.Stdout, coon)
	//传送消息给服务端
	ioCopy(coon, os.Stdin)
}

func ioCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("获取消息服务端失败", err)
		return
	}
}

// Server 服务端
type c chan<- string            //定义发送消息的通道（客户）
var message = make(chan string) //消息通道

func Server() {
	//绑定端口
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("绑定端口失败:", err)
		return
	}

	//广播：将每一个客户端的消息全部发送到已连接的客户端
	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[c]bool)
	for {
		msg := <-message
		for cli := range clients {
			cli <- msg
		}
	}
}

func handleConn(conn net.Conn) {
	who := conn.RemoteAddr().String() //谁发送的信息
	content := bufio.NewScanner(conn) //发送的内容
	for content.Scan() {
		message <- who + ":" + content.Text() //将“名字:...”传给message
	}
	conn.Close()
}

