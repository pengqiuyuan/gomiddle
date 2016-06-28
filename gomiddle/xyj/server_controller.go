package gomiddle

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"strings"
	"encoding/json"
	"../../gomiddle"
	proto "../../tutorial/tcp"
)

type ServerEntity struct {
	GameId       string
	ServerZoneId string
	ServerId     string
	Status       string
}

type ResponseList struct {
	Choose   string   `json:"choose"`
	Success  string   `json:"success"`
	ObjFail  []string `json:"objFail"`
	Fail     string   `json:"fail"`
	Status   string   `json:"status"`
}

func ServerHandler() {
	http.HandleFunc("/xyjserver/server/updateServers", UpdateServers)
}


func UpdateServers(w http.ResponseWriter, r *http.Request) {
	AddOrUpdateServer(proto.TcpProtoIDXyjUpdateServer, w, r)
}


func AddOrUpdateServer(m uint16, w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		//结构已知，解析到结构体
		var s ServerEntity
		json.Unmarshal([]byte(result), &s)
		fmt.Println(s)
		//多个serverId按，切分
		ser := strings.Split(s.ServerId, ",")
		var res,objF string
		for _, key := range ser {
			//判断serverid是否在ConnMap里
			conn, exists := gomiddle.ConnMap[key]
			if exists {
				fmt.Println(key, "  存在   ", conn)
				connid, _ := gomiddle.ConnMa[key]
				conn.Send(connid, makeNoticeMsg(string(result),m))

				select {
				case x := <-gomiddle.Channel_c:
					fmt.Println(key, "  存在,客户端有返回值  AddOrUpdate ",m)		
					var responseList ResponseList
			
					if err := json.Unmarshal([]byte(x[string(connid)+"_"+string(m)]), &responseList); err == nil {
						if len(responseList.ObjFail) != 0 {
							objF = responseList.ObjFail[0]
						}
						res = `{"choose":"` + responseList.Choose + `","success":"` + responseList.Success + `","objFail":"` + objF + `","fail":"` + responseList.Fail + `","status":"` + responseList.Status+ `"}`
					}
					
				case <-time.After(time.Second * 1):
					fmt.Println(key, "  存在,超时客户端无返回值  AddOrUpdate ",m)	
					//web server 修改服务器状态，游戏服务器存在但没有响应，返回-1				
					res = `{"choose":"1","success":"0","objFail":"` + key + `","fail":"1","status":"-1"}`
					
				}
			} else {
				fmt.Println(key, "  不存在  ")
				//web server 修改服务器状态，gomiddle服务器与游戏服务器断连，返回-2
				res = `{"choose":"1","success":"0","objFail":"` + key + `","fail":"1","status":"-2"}`
			}
		}
		b := []byte(res)
		w.Write(b)
	}
}

