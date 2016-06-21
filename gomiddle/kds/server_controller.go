package gomiddle

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"strings"
	"encoding/json"
	"../../gomiddle"
	entity "../../entity"
	proto "../../tutorial/tcp"
)

type ServerEntity struct {
	ServerId   string
	Status     string
}

func ServerHandler() {
	http.HandleFunc("/kdsserver/server/updateServers", UpdateServers)
}


func UpdateServers(w http.ResponseWriter, r *http.Request) {
	AddOrUpdateServer(proto.TcpProtoIDKdsUpdateServer, w, r)
}


func AddOrUpdateServer(m uint16, w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		//结构已知，解析到结构体
		var s ServerEntity
		json.Unmarshal([]byte(result), &s)
		//fmt.Println(s)
		//多个serverId按，切分
		ser := strings.Split(s.ServerId, ",")

		//ser := s.ServerIds
		choose := len(ser)
		success := 0
		fail := 0
		var objFail []string

		for _, key := range ser {
			//判断serverid是否在ConnMap里
			conn, exists := gomiddle.ConnMap[key]
			if exists {
				fmt.Println(key, "  存在   ", conn)
				connid, _ := gomiddle.ConnMa[key]
				conn.Send(connid, makeNoticeMsg(string(result),m))	
				
				select {
				case x := <-gomiddle.Channel_c:
					fmt.Println(key, "  存在,客户端有返回值  AddOrUpdate  ",m)
					var responseList entity.ResponseList
					if err := json.Unmarshal([]byte(x[string(connid)+"_"+string(m)]), &responseList); err == nil {
						success = success + responseList.Success
						fail = fail + responseList.Fail
						if len(responseList.ObjFail) != 0 {
							objFail = append(objFail, responseList.ObjFail[0])
						}
					}
				case <-time.After(time.Second * 1):
					fmt.Println(key, "  存在,超时客户端无返回值  AddOrUpdate  ",m)
					fail = fail + 1
					objFail = append(objFail, key)
				}
			} else {
				fmt.Println(key, "  不存在  ")
				fail = fail + 1
				objFail = append(objFail, key)
			}
		}
		respons := entity.ResponseList{Choose: choose, Success: success, ObjFail: objFail, Fail: fail}
		res, _ := json.Marshal(respons)
		b := []byte(res)
		w.Write(b)
	}
}

