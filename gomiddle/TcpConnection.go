package gomiddle

import (
    "bufio"
    "fmt"
    "net" 
    "database/sql"
    "encoding/json" 
	"time"

)

// 用来记录所有的客户端连接
var ConnMap map[string]*net.TCPConn   

type ServerInfoJson struct {
    ServerZoneId int `json:"serverZoneId"` 
    PlatForm []string `json:"platForm"` 
    ServerId string `json:"serverId"` 
    StoreId int `json:"storeId"` 
    Ip string `json:"ip"` 
    Port string `json:"port"` 
    Status string `json:"status"` 
}

func TcpCon(db *sql.DB){
	var tcpAddr *net.TCPAddr
    ConnMap = make(map[string]*net.TCPConn)
    tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

    tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

    defer tcpListener.Close()

    for {
        tcpConn, err := tcpListener.AcceptTCP()
        if err != nil {
            continue
        }

        fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
        go TcpPipe(tcpConn,db)
    }
}

func TcpPipe(conn *net.TCPConn,db *sql.DB) {
    ipStr := conn.RemoteAddr().String()
    defer func() {
        fmt.Println("disconnected :" + ipStr)
        conn.Close()
    }()
    reader := bufio.NewReader(conn)

    for {
        message, err := reader.ReadString('\n')
        if err != nil {
            return
        }
        fmt.Println(string(message))
        var jsonServer ServerInfoJson
        if err := json.Unmarshal([]byte(message), &jsonServer); err == nil {  
		     // 新连接加入map
       		 ConnMap[jsonServer.ServerId] = conn
            fmt.Printf("-->运营大区:%s 渠道:%s 服务器:%s 游戏:%s ip:%s 端口:%s 状态:%s\n", jsonServer.ServerZoneId,jsonServer.PlatForm,jsonServer.ServerId,int(jsonServer.StoreId),jsonServer.Ip,jsonServer.Port,jsonServer.Status)
            Insert_serverZone(db,int(jsonServer.ServerZoneId))
            Insert_storeId(db,int(jsonServer.StoreId))
            for  i:=0;i<len(jsonServer.PlatForm);i++ {
                 Insert_all_platform(db,int(jsonServer.ServerZoneId),int(jsonServer.StoreId),jsonServer.PlatForm[i],jsonServer.ServerId)
            }
            Select_all_server(db,int(jsonServer.ServerZoneId),int(jsonServer.StoreId),jsonServer.ServerId,jsonServer.Ip,jsonServer.Port,jsonServer.Status)
        }  
		msg := time.Now().String() + "\n"
        b := []byte(msg)
        conn.Write(b)
		fmt.Println(msg)
    }
}

