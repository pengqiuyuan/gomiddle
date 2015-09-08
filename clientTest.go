package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"./codec"
)

var quitSemaphore chan bool

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	fmt.Println("connected!")

	sendMessage(conn)
	go onMessageRecived(conn)
	//go sendMessage(conn)
	<-quitSemaphore
}

func sendMessage(conn *net.TCPConn) {
	/*
		for {
			time.Sleep(3 * time.Second)
			// var msg string
			// fmt.Scanln(&msg)
			s := strings.Split(conn.LocalAddr().String(), ":")
	        jsonStr := conn.LocalAddr().String()+`|serverStatus|request|{"serverZoneId":1,"platForm":["1","2"],"serverId":"fb_server_1","storeId":1,"ip":"`+s[0]+`","port":"`+s[1]+`","status":"1"}`
	        b, _ := codec.Encode(jsonStr)
	        conn.Write(b)
		}
	*/
	s := strings.Split(conn.LocalAddr().String(), ":")
	jsonStr := conn.LocalAddr().String() + `|serverStatus|{"serverZoneId":1,"platForm":["1","2"],"serverId":"fb_server_3","storeId":1,"ip":"` + s[0] + `","port":"` + s[1] + `","status":"1"}`
	b, _ := codec.Encode(jsonStr)
	conn.Write(b)
}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := codec.Decode(reader)
		fmt.Println(msg)
		if err != nil {
			quitSemaphore <- true
			break
		}
		s, _ := codec.Encode(conn.LocalAddr().String() + `|addPlacards|{"choose":1,"success":0,"objFail":["我是返回来的保存失败的消息"],"fail":1}|post`)
		conn.Write(s)
	}
}
