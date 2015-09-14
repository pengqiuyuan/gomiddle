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
	s := strings.Split(conn.LocalAddr().String(), ":")
	jsonStr := conn.LocalAddr().String() + `|serverStatus|{"serverZoneId":1,"platForm":["1","2"],"serverId":"fb_server_1","gameId":1,"ip":"` + s[0] + `","port":"` + s[1] + `","status":"1"}`
	b, _ := codec.Encode(jsonStr)
	conn.Write(b)
}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := codec.Decode(reader)
		fmt.Println(message)
		if err != nil {
			quitSemaphore <- true
			break
		}
		s := strings.Split(message, "|")
		if s[1] == "addPlacards" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|{"choose":1,"success":1,"objFail":[],"fail":0}|`+s[3]+``)
			conn.Write(str)
		} else if s[1] == "getTotalByServerZoneIdAndGameId" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|{"num":1}|`+s[3]+``)
			conn.Write(str)
		} else if s[1] == "getAllPlacards" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|[{"id":1,"serverZoneId":"1","gameId":"1","serverId":"fb_server_1","version":"1","contents":"1"}]|`+s[3]+``)
			conn.Write(str)
		}else if s[1] == "delPlacardById" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|{"message":"success"}|`+s[3]+``)
			conn.Write(str)
		}else if s[1] == "updatePlacards" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|{"choose":1,"success":1,"objFail":[],"fail":0}|`+s[3]+``)
			conn.Write(str)
		}

	}
}
