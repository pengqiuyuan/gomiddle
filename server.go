
package main

import (
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
    "log"
	"sync"
	"net/http"
	fb "./gomiddle/fb"
	"encoding/json"
	
	"enlightgame/net/tcp"
	"enlightgame/transport"
	proto "./tutorial/tcp"
	flatbuffers "github.com/google/flatbuffers/go"
	"os"
	"os/signal"
	"strconv"
)

type ServerInfoJson struct {
	ServerZoneId int      `json:"serverZoneId"`
	PlatForm     []string `json:"platForm"`
	ServerId     string   `json:"serverId"`
	GameId       int      `json:"gameId"`
	Ip           string   `json:"ip"`
	Port         string   `json:"port"`
	Status       string   `json:"status"`
}

type messageHandler func(uint32, *transport.TcpMessage)

var (
	wg         sync.WaitGroup
	a          *tcp.Acceptor
	cnt        int                       = 0
	dispatcher map[uint16]messageHandler = make(map[uint16]messageHandler)
)

func shutdown() {
	wg.Done()
}

func makeNoticeMsg() []byte {
	t := transport.TcpMessage{}

	builder := flatbuffers.NewBuilder(0)
	ct := builder.CreateString(strconv.Itoa(cnt))
	proto.NoticeStart(builder)
	cnt += 1
	proto.NoticeAddContent(builder, ct)
	payload := proto.NoticeEnd(builder)
	builder.Finish(payload)

	t.Payload = builder.Bytes[builder.Head():]

	// 填充协议头信息
	t.Header.Proto = proto.TcpProtoIDNotice
	t.Header.Flag = 0xdcba
	t.Header.Size = uint16(len(t.Payload))

	ret, err := t.Pack()
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return ret
}

func statusHandler(id uint32, t *transport.TcpMessage) {
	// 从消息payload部分获取正文内容
	s := proto.GetRootAsGtom(t.Payload, 0)
	//log.Printf("recv status! %v, %v, %v, %v, %v, %v, %v, %v", s.ServerZoneId(), s.PlatForm(0), s.PlatForm(1), s.ServerId(), s.GameId(), s.Ip(), s.Port(), s.Status())
	//log.Printf("recv status! %v, %v, %v, %v, %v, %v, %v, %v", string(s.ServerZoneId()), string(s.PlatForm(0)), string(s.PlatForm(1)), string(s.ServerId()), string(s.GameId()), string(s.Ip()), int(s.Port()), int(s.Status()) )
	

	log.Printf("222222222crecv status! %v",	a.RemoteAddr(id))
	var jsonServer ServerInfoJson
	if err := json.Unmarshal(s.Content(), &jsonServer); err == nil {
		log.Printf("-->运营大区:%s 渠道:%s 服务器:%s 游戏:%s ip:%s 端口:%s 状态:%s\n", jsonServer.ServerZoneId, jsonServer.PlatForm, jsonServer.ServerId, int(jsonServer.GameId), jsonServer.Ip, jsonServer.Port, jsonServer.Status)
	}
	
	// 发送notice消息
	//msg := makeNoticeMsg()
	//a.Send(id, msg)
}

func handleConnect(id uint32) {
	log.Println("accept: ", id)
}

func handleMessage(id uint32, b []byte) {
	log.Println("on message: ", b)
	rt := transport.TcpMessage{}

	// 从缓冲区获取消息内容
	err := rt.Unpack(b)
	if err != nil {
		log.Fatal(err.Error())
	}

	// 根据协议ID分发消息到指定的处理器
	handler, ok := dispatcher[rt.Header.Proto]
	if ok && handler != nil {
		handler(id, &rt)
	}
}

func handleDisconnect(id uint32) {
	log.Println("disconnect id: ", id)
}

func init() {
	// 根据协议类型创建acceptor
	p := tcp.ParseParam{}
	p.HeadSize = 23
	p.BodySizeOffset = 21
	p.BodySizeLen = 2
	p.NotifyWithHead = true
	a = tcp.NewAcceptor(":8898", p)

	// 注册消息处理器
	dispatcher[proto.TcpProtoIDStatus] = statusHandler
}


func main() {
	var wg sync.WaitGroup
	db, err := sql.Open("mysql", "root:123456@tcp(10.0.29.251:3306)/game_server?charset=utf8")
    if err != nil {
        log.Println("mysql init failed")
        return
    }else{
        log.Println("mysql init ok")
    }
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

func Handle(){
	FbHandle()
	KdsHandle()
	KunHandle()
	err := http.ListenAndServe(":8899", nil)
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}

func FbHandle(){
	fb.PlacardHandler()
	fb.GagHandler()
	fb.SealHandler()
	fb.EmailHandler()
	fb.ProductHandler()
}

func KdsHandle(){
	
}

func KunHandle(){
	
}



