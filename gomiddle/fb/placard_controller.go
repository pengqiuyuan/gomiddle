package gomiddle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"../../codec"
	entity "../../entity"
	"../../gomiddle"
)

type PlacardEntity struct {
	GameId       string
	ServerZoneId string
	ServerId     string
	Version      string
	Contents     string
}

func PlacardHandler() {
	http.HandleFunc("/fbserver/placard/addPlacards", SavePlacardHandler)
	http.HandleFunc("/fbserver/getTotalByServerZoneIdAndGameId", GetTotalByServerZoneIdAndGameId)
	err := http.ListenAndServe(":8899", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}

/**
 * 查询运营大区、游戏下 公告的总数
 * 参数 localhost:8899/fbserver/getTotalByServerZoneIdAndGameId?serverZoneId=1&storeId=1&category=placard&serverId=fb_server_1
 * 传入格式 127.0.0.1:54726|getTotalByServerZoneIdAndGameId|{"serverZoneId":"1","storeId":"1","category":"placard","serverId":"fb_server_1"}|get
 * 返回格式 127.0.0.1:54726|getTotalByServerZoneIdAndGameId|{}|get
 */
func GetTotalByServerZoneIdAndGameId(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		serverZoneId := r.FormValue("serverZoneId")
		storeId := r.FormValue("storeId")
		category := r.FormValue("category")
		serverId := r.FormValue("serverId")
		JsonStr := `{"serverZoneId":"` + serverZoneId + `","storeId":"` + storeId + `","category":"` + category + `","serverId":"` + serverId + `"}`
		
		conn, exists := gomiddle.ConnMap[r.FormValue("serverId")]
		if exists {
			fmt.Println(r.FormValue("serverId"), "  存在   ", conn)
			b, err := codec.Encode(conn.RemoteAddr().String() + "|getTotalByServerZoneIdAndGameId|" + string(JsonStr) + "|get")
			if err != nil {
				fmt.Println(err)
			}
			conn.Write(b)
		}
	}
}

/**
 * 保存公告
 * 传入数据格式 127.0.0.1:54726|addPlacards|{"id":2,"serverZoneId":"1","contents":"<p>\r\n\t1<\/p>\r\n","gameId":"1","serverId":"fb_server_1,fb_server_2,fb_server_3","serverIds":null,"version":"1"}|post
 * 返回数据格式 
 */
func SavePlacardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		//结构已知，解析到结构体
		var s PlacardEntity
		json.Unmarshal([]byte(result), &s)
		//fmt.Println(s)
		//多个serverId按，切分
		ser := strings.Split(s.ServerId, ",")

		choose := len(ser)
		success := 0
		fail := 0
		var objFail []string

		for _, key := range ser {
			//判断serverid是否在ConnMap里
			conn, exists := gomiddle.ConnMap[key]
			if exists {
				fmt.Println(key, "  存在   ", conn)
				b, err := codec.Encode(conn.RemoteAddr().String() + "|addPlacards|" + string(result) + "|post")
				if err != nil {
					continue
				}
				conn.Write(b)
				x := <-gomiddle.Channel_c

				var responseList entity.ResponseList
				if err := json.Unmarshal([]byte(x[conn.RemoteAddr().String()+"_addPlacards"]), &responseList); err == nil {
					success = success + responseList.Success
					fail = fail + responseList.Fail
					if len(responseList.ObjFail) != 0 {
						objFail = append(objFail, responseList.ObjFail[0])
					}

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
