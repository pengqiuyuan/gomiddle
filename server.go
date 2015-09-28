package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	_ "github.com/go-sql-driver/mysql"

	"enlightgame/net/tcp"
	"enlightgame/transport"
	"os"
	"os/signal"

	hql "./gomiddle"
	fb "./gomiddle/fb"
	proto "./tutorial/tcp"
)



type ServerInfoJson struct {
	ServerZoneId int      `json:"serverZoneId"`
	PlatForm     []string `json:"platForm"`
	ServerId     string   `json:"serverId"`
	GameId       int      `json:"gameId"`
	Status       string   `json:"status"`
}


var (
	wg         sync.WaitGroup
	a          *tcp.Acceptor
	cnt        int                       = 0
	db *sql.DB
)

func shutdown() {
	wg.Done()
}

func handleConnect(id uint32) {
	log.Println("accept: ", id)
}

func handleMessage(id uint32, b []byte) {
	log.Println("on message: ", b)
	t := transport.TcpMessage{}

	// 从缓冲区获取消息内容
	err := t.Unpack(b)
	if err != nil {
		log.Fatal(err.Error())
	}

	m := t.Header
	//唯一游戏服务器发送的消息，服务器状态路由 TcpProtoIDStatus
	if m.Proto == proto.TcpProtoIDGmStatus {
		// 从消息payload部分获取正文内容
		s := proto.GetRootAsNotice(t.Payload, 0)
		var jsonServer ServerInfoJson
		if err := json.Unmarshal(s.Content(), &jsonServer); err == nil {
			// 新连接加入map
			hql.ConnMap[jsonServer.ServerId] = a
			hql.ConnMa[jsonServer.ServerId] = id
			hql.ConnM[id] = jsonServer.ServerId
			sip := strings.Split(a.RemoteAddr(id), ":")
			fmt.Printf("-->运营大区:%s 渠道:%s 服务器:%s 游戏:%s ip:%s 端口:%s 状态:%s\n", jsonServer.ServerZoneId, jsonServer.PlatForm, jsonServer.ServerId, int(jsonServer.GameId), sip[0], sip[1], jsonServer.Status)
			hql.Insert_serverZone(db, int(jsonServer.ServerZoneId))
			hql.Insert_gameId(db, int(jsonServer.GameId))
			for i := 0; i < len(jsonServer.PlatForm); i++ {
				hql.Insert_all_platform(db, int(jsonServer.ServerZoneId), int(jsonServer.GameId), jsonServer.PlatForm[i], jsonServer.ServerId)
			}
			hql.Select_all_server(db, int(jsonServer.ServerZoneId), int(jsonServer.GameId), jsonServer.ServerId, sip[0], sip[1], jsonServer.Status)
		}
	} else {
		// 从消息payload部分获取正文内容
		s := proto.GetRootAsNotice(t.Payload, 0)
		//   1_2   {"choose":1,"success":1,"objFail":["我是返回来的消息"],"fail":1}
		hql.ResponseMap[string(id)+"_"+string(m.Proto)] = string(s.Content())
		hql.Channel_c <- hql.ResponseMap
	}
	
}

func handleDisconnect(id uint32) {
	log.Println("disconnect id: ", id)
	delete(hql.ConnMap, hql.ConnM[id])
	delete(hql.ConnMa, hql.ConnM[id])
	delete(hql.ConnM, id)
	//移除断开客户端保存再ResponseMap里的消息
	for key, _ := range hql.ResponseMap {
		ke := strings.Split(key, "_")
		if ke[0] == string(id) {
			delete(hql.ResponseMap, key)
		}
	}
	sip := strings.Split(a.RemoteAddr(id), ":")
	//mysql删除对应保存的客户端信息
	hql.Delete_server(db, sip[0], sip[1])
}

func init() {
	// 根据协议类型创建acceptor
	p := tcp.ParseParam{}
	p.HeadSize = 23
	p.BodySizeOffset = 21
	p.BodySizeLen = 2
	p.NotifyWithHead = true
	a = tcp.NewAcceptor(":8898", p)	
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(10.0.29.251:3306)/game_server?charset=utf8")
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	
	log.SetFlags(log.Flags() | log.Lshortfile)
	a.HandleConnect(handleConnect)
	a.HandleMessage(handleMessage)
	a.HandleDisconnect(handleDisconnect)

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go a.Start()
	go Handle()
	<-ch
	a.Stop(shutdown)
	wg.Wait()

}

func Handle() {
	FbHandle()
	KdsHandle()
	KunHandle()
	err := http.ListenAndServe(":8899", nil)
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}

func FbHandle() {
	fb.ServerHandler()
	fb.GrayAccountHandler()
	fb.PlacardHandler()
	fb.GagHandler()
	fb.SealHandler()
	fb.EmailHandler()
	fb.ProductHandler()
}

func KdsHandle() {

}

func KunHandle() {

}
