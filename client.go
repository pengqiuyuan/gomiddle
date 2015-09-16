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
		
		if s[1] == "addSealAccount" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|{"message":"success"}|`+s[3]+``)
			conn.Write(str)
		}else if s[1] == "updateSealAccount" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|{"message":"success"}|`+s[3]+``)
			conn.Write(str)
		}else if s[1] == "getAllSealAccount" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|[{"id":1,"serverZoneId":"1","gameId":"1","serverId":"fb_server_1","guid":"111","name":"adc","account":"1","platForm":"1","sealTime":"1234","sealStart":"2014-12-11 16:55:15","sealEnd":"2014-12-11 16:55:15"}]|`+s[3]+``)
			conn.Write(str)
		}else if s[1] == "delSealAccount" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|{"message":"success"}|`+s[3]+``)
			conn.Write(str)
		}
		
		if s[1] == "addGagAccount" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|{"message":"success"}|`+s[3]+``)
			conn.Write(str)
		}else if s[1] == "updateGagAccount" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|{"message":"success"}|`+s[3]+``)
			conn.Write(str)
		}else if s[1] == "getAllGagAccount" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|[{"id":1,"serverZoneId":"1","gameId":"1","serverId":"fb_server_1","guid":"111","name":"adc","account":"1","platForm":"1","gagTime":"1234","gagStart":"2014-12-11 16:55:15","gagEnd":"2014-12-11 16:55:15"}]|`+s[3]+``)																
			conn.Write(str)
		}else if s[1] == "delGagAccountById" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|{"message":"success"}|`+s[3]+``)
			conn.Write(str)
		}
		
		if s[1] == "addProduct" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|{"message":"success"}|`+s[3]+``)
			conn.Write(str)
		}else if s[1] == "updateProduct" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|{"message":"success"}|`+s[3]+``)
			conn.Write(str)
		}else if s[1] == "getAllProducts" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|[{"id":1,"serverZoneId":"1","gameId":"1","serverId":"fb_server_1","itemId":"1","num":"1","prodcutStoreId":"1","storeLocation":"1","isRandom":"1","randomProbability":"1","comsumeType":"1","comsumeNum":"1","discount":"1","levelLimit":"1","levelCap":"1","discountStartDate":"2014-12-11 16:55:15","discountContinueDate":"2014-12-11 16:55:15","discountCycleDate":"2014-12-11 16:55:15","productPostDate":"2014-12-11 16:55:15","productDownDate":"2014-12-11 16:55:15","showLevel":"1"}]|`+s[3]+``)																
			conn.Write(str)
		}else if s[1] == "delProductById" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|{"message":"success"}|`+s[3]+``)
			conn.Write(str)
		}
		
		if s[1] == "addEmail" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|{"choose":1,"success":1,"objFail":[],"fail":0}|`+s[3]+``)
			conn.Write(str)
		}else if s[1] == "updateEmail" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|{"choose":1,"success":1,"objFail":[],"fail":0}|`+s[3]+``)
			conn.Write(str)
		}else if s[1] == "getAllEmails" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|[{"id":1,"serverZoneId":"1","gameId":"1","serverId":"fb_server_1","sender":"1","title":"1","contents":"1","annex":[{"itemId":"1","itemNum":1},{"itemId":"1","itemNum":1},{"itemId":"112","itemNum":1123}] }]|`+s[3]+``)																
			conn.Write(str)
		}else if s[1] == "delEmailById" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|{"message":"success"}|`+s[3]+``)
			conn.Write(str)
		}else if s[1] == "getEmailById" {
			str, _ := codec.Encode(conn.LocalAddr().String() + `|`+s[1]+`|{"id":1,"serverZoneId":"1","gameId":"1","serverId":"fb_server_1","sender":"1","title":"1","contents":"1","annex":[{"itemId":"1","itemNum":1},{"itemId":"1","itemNum":1},{"itemId":"112","itemNum":11234444}] }|`+s[3]+``)
			conn.Write(str)
		}

	}
}
