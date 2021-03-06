package main

import (
	// "bytes"
	// "encoding/binary"
	"enlightgame/net/tcp"
	"enlightgame/transport"
	"log"
	"os"
	"os/signal"
	"time"
	"fmt"
	proto "./tutorial/tcp"
	flatbuffers "github.com/google/flatbuffers/go"
)

var (
	c *tcp.Connector
)

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
	m := t.Header
	log.Printf("recv notice! %v", n.Content())
	log.Printf("recv notice! %v", string(n.Content()))
	log.Printf("recv notice! %v %v %v", m.Flag, m.Proto, m.Size)

	if m.Proto == proto.TcpProtoIDXyjUpdateServer {
		str := `{"choose":"1","success":"1","objFail":[],"fail":"0","status":"3"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	}

	if m.Proto == proto.TcpProtoIDXyjSaveGrayAccount {
		str := `{"choose":"1","success":"1","objFail":[],"fail":"0"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjGetGrayAccountById {
		str := `{"id":1,"serverZoneId":"1","gameId":"1","serverId":"kds_server_1","platForm":"2","account":"2"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjGetAllGrayAccount {
		str := `[{"account": "88888888","platForm": "qq"},{"account": "555","platForm": "3434"}]`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjDelGrayAccountById {
		str := `{"message":"success"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjUpdateGrayAccount {
		str := `{"choose":"1","success":"1","objFail":[],"fail":"0"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjGrayGetTotalByServerZoneIdAndGameId {
		str := `{"num":1}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} 

	if m.Proto == proto.TcpProtoIDXyjSavePlacard {
		str := `{"choose":"1","success":"1","objFail":[],"fail":"0"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjPlacardGetTotalByServerZoneIdAndGameId {
		str := `{"num":1}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjGetAllPlacards {
		str := `[{"id":1,"serverZoneId":"1","gameId":"1","serverId":"xyj_server_test","version":"1","contents":"1"}]`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjDelPlacardById {
		str := `{"message":"success"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjUpdatePlacards {
		str := `{"choose":"1","success":"1","objFail":[],"fail":"0"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjGetPlacardById {
		str := `{"id":1,"serverZoneId":"1","gameId":"1","serverId":"xyj_server_test","version":"1","contents":"我们来了"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	}

	if m.Proto == proto.TcpProtoIDXyjGetAllEmails {
		str := `[{"id":1,"serverZoneId":"1","gameId":"1","serverId":"xyj_server_test","sender":"1","title":"1","contents":"1","annex":[{"itemId":"1","itemNum":1},{"itemId":"1","itemNum":1},{"itemId":"112","itemNum":1123}] }]`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjAddEmail {
		str := `{"choose":"1","success":"0","objFail":["我是一个测试"],"fail":"1"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjUpdateEmail {
		str := `{"choose":"1","success":"1","objFail":[],"fail":"0"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjDelEmailById {
		str := `{"message":"success"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjGetEmailById {
		str := `{"id":1,"serverZoneId":"1","gameId":"1","serverId":"xyj_server_1","sender":"1","title":"1","contents":"1","annex":[{"itemId":"1","itemNum":1},{"itemId":"1","itemNum":1},{"itemId":"112","itemNum":11234444}] }`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjEmailGetTotalByServerZoneIdAndGameId {
		str := `{"num":1}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	}


	if m.Proto == proto.TcpProtoIDXyjGetAllGagAccount {
		str := `[{"guid":"1173993848902","name":"player_2","account":"sk2","platForm":"win32","gagTime":"-1","gagStart":"1467874079","gagEnd":"-1"},{"guid":"1173992487619","name":"player_6","account":"xc2203","platForm":"win32","gagTime":"43200","gagStart":"1467874200","gagEnd":"1467917400"}]`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjAddGagAccount {
		str := `{"message":"success"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjUpdateGagAccount {
		str := `{"message":"success"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjDelGagAccountById {
		str := `{"message":"success"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjGagGetTotalByServerZoneIdAndGameId {
		str := `{"num":1}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} 
	
	if m.Proto == proto.TcpProtoIDXyjGetAllSealAccount {
		str := `[{"guid":"111","name":"adc","account":"1","platForm":"1","sealTime":"1234","sealStart":"1467871730","sealEnd":"1467871730"},{"guid":"111","name":"adc","account":"1","platForm":"1","sealTime":"1234","sealStart":"1467871730","sealEnd":"1467871730"},{"guid":"111","name":"adc","account":"1","platForm":"1","sealTime":"1234","sealStart":"1467871730","sealEnd":"1467871730"}]`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjAddSealAccount {
		str := `{"message":"success"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjUpdateSealAccount {
		str := `{"message":"success"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjDelSealAccount {
		str := `{"message":"success"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	}  else if m.Proto == proto.TcpProtoIDXyjSealGetTotalByServerZoneIdAndGameId {
		str := `{"num":3}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	}
	
	if m.Proto == proto.TcpProtoIDXyjAddEventPrototype {
		
		// 从消息payload部分获取正文内容
		s := proto.GetRootAsNotice(t.Payload, 0)
		fmt.Println("  测试1111   ", string(s.Content()))
		
		str := `{"choose":"1","success":"1","objFail":[],"fail":"0"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjUpdateEventPrototype {
				
		// 从消息payload部分获取正文内容
		s := proto.GetRootAsNotice(t.Payload, 0)
		fmt.Println("  测试2222   ", string(s.Content()))
	
		str := `{"choose":"1","success":"1","objFail":[],"fail":"0"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjAddEventDataPrototype {
				
		// 从消息payload部分获取正文内容
		s := proto.GetRootAsNotice(t.Payload, 0)
		fmt.Println("  测试3333   ", string(s.Content()))
	
		str := `{"choose":"1","success":"1","objFail":[],"fail":"0"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjUpdateEventDataPrototype {
				
		// 从消息payload部分获取正文内容
		s := proto.GetRootAsNotice(t.Payload, 0)
		fmt.Println("  测试4444   ", string(s.Content()))
	
		str := `{"choose":"1","success":"1","objFail":[],"fail":"0"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjCloseEventPrototype {
				
		// 从消息payload部分获取正文内容
		s := proto.GetRootAsNotice(t.Payload, 0)
		fmt.Println("  测试5555   ", string(s.Content()))
	
		str := `{"message":"success"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} 
	
	if m.Proto == proto.TcpProtoIDGmStatus { //获取 gomiddle 返回的活动列表
		// 从消息payload部分获取正文内容
		s := proto.GetRootAsNotice(t.Payload, 0)
		fmt.Println("  测试   ", string(s.Content()))
	}
	
}

func makeNoticeMsg(str string, p uint16) []byte {
	t := transport.TcpMessage{}

	builder := flatbuffers.NewBuilder(0)

	ct := builder.CreateString(str)
	proto.NoticeStart(builder)
	proto.NoticeAddContent(builder, ct)
	payload := proto.NoticeEnd(builder)

	builder.Finish(payload)

	t.Payload = builder.Bytes[builder.Head():]

	// 填充协议头信息
	t.Header.Proto = p
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
	c := tcp.NewConnector(id, "127.0.0.1:8888", p)
	c.HandleMessage(clientHandleMessage)
	c.HandleConnect(handleConnect)
	c.HandleDisconnect(handleDisconnect)
	c.Start()
	return c
}

func main() {
	str := `{"serverZoneId":"1","platForm":["2","3"],"serverId":"xyj_server_test","gameId":"4","status":"1"}`

	log.SetFlags(log.Flags() | log.Lshortfile)
	c = client()

	t := time.NewTicker(interval * time.Second)
	select {
	case _ = <-t.C:
		m := makeNoticeMsg(str, proto.TcpProtoIDGmStatus)
		c.Send(m)
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch

}
