package gomiddle

import (
	"enlightgame/net/tcp"
)

var (
	ConnMap map[string]*tcp.Acceptor = make(map[string]*tcp.Acceptor)  // 用来记录所有的客户端连接 ConnMap (key:fb_server_1 value:a) ConnM(key:fb_server_1 value:1) ConnM(key:1 value:fb_server_1)
	ConnMa map[string]uint32 = make(map[string]uint32)
	ConnM map[uint32]string = make(map[uint32]string)
	ResponseMap map[string]string = make(map[string]string)  // 用来记录所有接收到游戏服务器发来的消息
	Channel_c = make(chan map[string]string, 1) // chan用来返回ResponseMap给httphandle
)
