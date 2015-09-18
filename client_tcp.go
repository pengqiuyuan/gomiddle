package main

import (
	// "bytes"
	// "encoding/binary"
	"enlightgame/net/tcp"
	"enlightgame/transport"
	proto "./tutorial/tcp"
	flatbuffers "github.com/google/flatbuffers/go"
	"log"
	"time"
)

var ()

const (
	interval = 2
)

func clientHandleMessage(id uint32, b []byte) {
	t := transport.TcpMessage{}
	err := t.Unpack(b)
	if err != nil {
		log.Fatal(err.Error())
	}

	n := proto.GetRootAsNotice(t.Payload, 0)
	log.Printf("recv notice! %v", n.Content())
}

func makeStatusMsg() []byte {
	t := transport.TcpMessage{}
	builder := flatbuffers.NewBuilder(0)
	zone := builder.CreateString("zone")
	p1 := builder.CreateString("platform 1")
	p2 := builder.CreateString("platform 2")

	proto.StatusStartPlatFormVector(builder, 2)
	builder.PrependUOffsetT(p1)
	builder.PrependUOffsetT(p2)
	p := builder.EndVector(2)

	serverId := builder.CreateString("server id")
	game := builder.CreateString("game")
	ip := builder.CreateString("ip")

	proto.StatusStart(builder)
	proto.StatusAddServerZoneId(builder, zone)

	proto.StatusAddPlatForm(builder, p)

	proto.StatusAddServerId(builder, serverId)
	proto.StatusAddGameId(builder, game)
	proto.StatusAddIp(builder, ip)
	proto.StatusAddPort(builder, 12345)
	proto.StatusAddStatus(builder, 0)

	payload := proto.StatusEnd(builder)
	builder.Finish(payload)

	t.Payload = builder.Bytes[builder.Head():]

	// 填充头部信息
	t.Header.Flag = 0xdcba
	t.Header.Proto = proto.TcpProtoIDStatus
	t.Header.Size = uint16(len(t.Payload))

	ret, err := t.Pack()
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return ret
}

func makeGtomMsg() []byte {
	t := transport.TcpMessage{}

	builder := flatbuffers.NewBuilder(0)
	
	str:= `{"serverZoneId":1,"platForm":["1","2"],"serverId":"fb_server_1","gameId":1,"ip":"10.0.0.11","port":"1111","status":"1"}`
	
	ct := builder.CreateString(str)
	proto.NoticeStart(builder)
	proto.NoticeAddContent(builder, ct)
	payload := proto.NoticeEnd(builder)
	
	builder.Finish(payload)

	t.Payload = builder.Bytes[builder.Head():]

	// 填充协议头信息
	t.Header.Proto = proto.TcpProtoIDStatus
	t.Header.Flag = 0xdcba
	t.Header.Size = uint16(len(t.Payload))

	ret, err := t.Pack()
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return ret
}


func handleConnect(id uint32) {
	log.Println("on connected: ", id)
}

func handleDisconnect(id uint32) {
	log.Println("disconnect: ", id)
}

func client() *tcp.Connector {
	var id uint32 = 0
	log.Println("client start")
	p := tcp.ParseParam{}
	p.HeadSize = 23
	p.BodySizeOffset = 21
	p.BodySizeLen = 2
	p.NotifyWithHead = true
	c := tcp.NewConnector(id, "10.0.29.152:8898", p)
	c.HandleMessage(clientHandleMessage)
	c.HandleConnect(handleConnect)
	c.HandleDisconnect(handleDisconnect)
	c.Start()
	return c
}

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	c := client()

	t := time.NewTicker(interval * time.Second * 2)
	for {
		select {
		case _ = <-t.C:
			//m := makeStatusMsg()
			m :=makeGtomMsg()
			c.Send(m)
		}
	}

}
