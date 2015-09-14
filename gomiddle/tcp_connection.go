package gomiddle

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"../codec"
)

// 用来记录所有的客户端连接 ConnMap (key:fb_server_1 value:conn) ConnM(key:127.0.0.1:60604 value:fb_server_1)
var ConnMap map[string]*net.TCPConn
var ConnM   map[string]string  = make(map[string]string)
// 用来记录所有接收到游戏服务器发来的消息
var ResponseMap map[string]string = make(map[string]string)
// chan用来返回ResponseMap给httphandle
var Channel_c = make(chan map[string]string,1)

type ServerInfoJson struct {
	ServerZoneId int      `json:"serverZoneId"`
	PlatForm     []string `json:"platForm"`
	ServerId     string   `json:"serverId"`
	GameId      int      `json:"gameId"`
	Ip           string   `json:"ip"`
	Port         string   `json:"port"`
	Status       string   `json:"status"`
}

func TcpCon(db *sql.DB) {
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
		go TcpPipe(tcpConn, db)
	}
}

func TcpPipe(conn *net.TCPConn, db *sql.DB) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		//移除客户端的连接
		delete(ConnMap, ConnM[ipStr])
		delete(ConnM,ipStr)
		//移除断开客户端保存再ResponseMap里的消息
		for key,_ := range ResponseMap {
			ke := strings.Split(key,"_")
			if ke[0] == ipStr {
				delete(ResponseMap, key)
			}
		}
		sip := strings.Split(ipStr, ":")
		//mysql删除对应保存的客户端信息
		Delete_server(db, sip[0], sip[1])
		//断开
		conn.Close()
	}()
	reader := bufio.NewReader(conn)
	for {
		message, err := codec.Decode(reader)
		if err != nil {
			return
		}
		//fmt.Println(message)
		s := strings.Split(message, "|")
		//唯一游戏服务器发送的消息，服务器状态路由 serverStatus
		if s[1] == "serverStatus" {
			var jsonServer ServerInfoJson
			if err := json.Unmarshal([]byte(s[2]), &jsonServer); err == nil {
				// 新连接加入map
				ConnMap[jsonServer.ServerId] = conn
				ConnM[ipStr] = jsonServer.ServerId
				fmt.Printf("-->运营大区:%s 渠道:%s 服务器:%s 游戏:%s ip:%s 端口:%s 状态:%s\n", jsonServer.ServerZoneId, jsonServer.PlatForm, jsonServer.ServerId, int(jsonServer.GameId), jsonServer.Ip, jsonServer.Port, jsonServer.Status)
				Insert_serverZone(db, int(jsonServer.ServerZoneId))
				Insert_gameId(db, int(jsonServer.GameId))
				for i := 0; i < len(jsonServer.PlatForm); i++ {
					Insert_all_platform(db, int(jsonServer.ServerZoneId), int(jsonServer.GameId), jsonServer.PlatForm[i], jsonServer.ServerId)
				}
				Select_all_server(db, int(jsonServer.ServerZoneId), int(jsonServer.GameId), jsonServer.ServerId, jsonServer.Ip, jsonServer.Port, jsonServer.Status)
			}
		} else {
			//   127.0.0.1:53846_addPlacards   {"choose":1,"success":1,"objFail":["我是返回来的消息"],"fail":1}
			ResponseMap[s[0]+"_"+s[1]] = s[2]
			Channel_c <- ResponseMap
		}
	}
}
