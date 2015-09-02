
package main

import (
    "bufio"
    "fmt"
    "net"
	"strings"
)

var quitSemaphore chan bool

func main() {
    var tcpAddr *net.TCPAddr
    tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

    conn, _ := net.DialTCP("tcp", nil, tcpAddr)
    defer conn.Close()
    fmt.Println("connected!")


	//go onMessageRecived(conn)
	/*
    c := time.Tick(10 * time.Second)
    for now := range c {

    }
	*/
		s := strings.Split(conn.LocalAddr().String(), ":")
        //fmt.Printf("%v \n", now)
        jsonStr := `{"serverZoneId":1,"platForm":["1","2"],"serverId":"fb_server_1","storeId":1,"ip":"`+s[0]+`","port":"`+s[1]+`","status":"1"}`
        b := []byte(jsonStr+"\n")        
        conn.Write(b)
		
    <-quitSemaphore
}

func onMessageRecived(conn *net.TCPConn) {
    reader := bufio.NewReader(conn)
    for {
        msg, err := reader.ReadString('\n')
        fmt.Println(msg)
        if err != nil {
            quitSemaphore <- true
            break
        }
        b := []byte(msg)
        conn.Write(b)
    }
}