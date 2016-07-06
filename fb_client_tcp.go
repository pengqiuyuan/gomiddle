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

	if m.Proto == proto.TcpProtoIDFbUpdateServer {
		str := `{"choose":"1","success":"1","objFail":[],"fail":"0","status":"3"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	}

	if m.Proto == proto.TcpProtoIDFbSaveGrayAccount {
		str := `{"choose":"1","success":"1","objFail":[],"fail":"0"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDFbGetGrayAccountById {
		str := `{"id":1,"serverZoneId":"1","gameId":"1","serverId":"fb_server_1","platForm":"2","account":"2"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDFbGetAllGrayAccount {
		str := `[{"account": "88888888","platForm": "qq"},{"account": "555","platForm": "3434"}]`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDFbDelGrayAccountById {
		str := `{"message":"success"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDFbUpdateGrayAccount {
		str := `{"choose":"1","success":"1","objFail":[],"fail":"0"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	}

	if m.Proto == proto.TcpProtoIDFbSavePlacard {
		str := `{"choose":"1","success":"1","objFail":[],"fail":"0"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDFbGetTotalByServerZoneIdAndGameId {
		str := `{"num":1}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDFbGetAllPlacards {
		str := `[{"id":1,"serverZoneId":"1","gameId":"1","serverId":"xyj_server_test","version":"1","contents":"1"}]`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDFbDelPlacardById {
		str := `{"message":"success"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDFbUpdatePlacards {
		str := `{"choose":"1","success":"1","objFail":[],"fail":"0"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDXyjGetPlacardById {
		str := `{"id":1,"serverZoneId":"1","gameId":"1","serverId":"xyj_server_test","version":"1","contents":"我们来了"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	}

	if m.Proto == proto.TcpProtoIDFbGetAllEmails {
		str := `[{"id":1,"serverZoneId":"1","gameId":"1","serverId":"fb_server_1","sender":"1","title":"1","contents":"1","annex":[{"itemId":"1","itemNum":1},{"itemId":"1","itemNum":1},{"itemId":"112","itemNum":1123}] }]`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDFbAddEmail {
		str := `{"choose":"1","success":"1","objFail":[],"fail":"0"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDFbUpdateEmail {
		str := `{"choose":"1","success":"1","objFail":[],"fail":"0"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDFbDelEmailById {
		str := `{"message":"success"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDFbGetEmailById {
		str := `{"id":1,"serverZoneId":"1","gameId":"1","serverId":"fb_server_1","sender":"1","title":"1","contents":"1","annex":[{"itemId":"1","itemNum":1},{"itemId":"1","itemNum":1},{"itemId":"112","itemNum":11234444}] }`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	}


	if m.Proto == proto.TcpProtoIDFbGetAllGagAccount {
		str := `[{"id":1,"serverZoneId":"1","gameId":"1","serverId":"fb_server_1","guid":"111","name":"adc","account":"1","platForm":"1","gagTime":"1234","gagStart":"2014-12-11 16:55:15","gagEnd":"2014-12-11 16:55:15"}]`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDFbAddGagAccount {
		str := `{"message":"success"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDFbUpdateGagAccount {
		str := `{"message":"success"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDFbDelGagAccountById {
		str := `{"message":"success"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} 
	
	if m.Proto == proto.TcpProtoIDFbGetAllSealAccount {
		str := `[{"id":1,"serverZoneId":"1","gameId":"1","serverId":"fb_server_1","guid":"111","name":"adc","account":"1","platForm":"1","sealTime":"1234","sealStart":"2014-12-11 16:55:15","sealEnd":"2014-12-11 16:55:15"}]`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDFbAddSealAccount {
		str := `{"message":"success"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDFbUpdateSealAccount {
		str := `{"message":"success"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDFbDelSealAccount {
		str := `{"message":"success"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} 
	
			
	if m.Proto == proto.TcpProtoIDFbGetAllProducts {
		str := `[{"id":1,"serverZoneId":"1","gameId":"1","serverId":"fb_server_1","itemId":"1","num":"1","prodcutStoreId":"1","storeLocation":"1","isRandom":"1","randomProbability":"1","comsumeType":"1","comsumeNum":"1","discount":"1","levelLimit":"1","levelCap":"1","discountStartDate":"2014-12-11 16:55:15","discountContinueDate":"2014-12-11 16:55:15","discountCycleDate":"2014-12-11 16:55:15","productPostDate":"2014-12-11 16:55:15","productDownDate":"2014-12-11 16:55:15","showLevel":"1"}]`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDFbAddProduct {
		str := `{"choose":1,"success":1,"objFail":[],"fail":"0"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDFbUpdateProduct {
		str := `{"choose":"1","success":"1","objFail":[],"fail":"0"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
	} else if m.Proto == proto.TcpProtoIDFbDelProductById {
		str := `{"message":"success"}`
		m := makeNoticeMsg(str, m.Proto)
		c.Send(m)
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
	c := tcp.NewConnector(id, "10.0.29.111:8888", p)
	c.HandleMessage(clientHandleMessage)
	c.HandleConnect(handleConnect)
	c.HandleDisconnect(handleDisconnect)
	c.Start()
	return c
}

func main() {
	str := `{"serverZoneId":"1","platForm":["1","2"],"serverId":"fb_server_1","gameId":"1","status":"1"}`

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
