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

type GrayAccountEntity struct {
	ServerZoneId string 
	GameId string
	ServerId string
	PlatForm string
	Account string
}

func GrayAccountHandler() {
	http.HandleFunc("/fbserver/server/getAllGrayAccount", GetAllGrayAccount)
	http.HandleFunc("/fbserver/server/addGrayAccount", AddGrayAccount)
	http.HandleFunc("/fbserver/server/updateGrayAccount", UpdateGrayAccount)
	http.HandleFunc("/fbserver/server/delGrayAccountById", DelGrayAccountById)
	http.HandleFunc("/fbserver/server/getGrayAccountByAccountId", GetGrayAccountByAccountId)
	http.HandleFunc("/fbserver/server/getTotalByServerZoneIdAndGameId", TcpProtoIDFbGrayGetTotalByServerZoneIdAndGameId)
}

func GetAllGrayAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		serverZoneId := r.FormValue("serverZoneId")
		gameId := r.FormValue("gameId")
		serverId := r.FormValue("serverId")
		pageNumber := r.FormValue("pageNumber")
		pageSize := r.FormValue("pageSize")
		JsonStr := `{"serverZoneId":"` + serverZoneId + `","gameId":"` + gameId + `","serverId":"` + serverId + `","pageNumber":"` + pageNumber + `","pageSize":"` + pageSize + `"}`
		conn, exists := gomiddle.ConnMap[serverId]
		var res string
		if exists {
			fmt.Println(r.FormValue("serverId"), "  存在   ", conn)
			connid, _ := gomiddle.ConnMa[serverId]

			conn.Send(connid, makeNoticeMsg(JsonStr,proto.TcpProtoIDFbGetAllGrayAccount))	
			
	
			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  GetAllGrayAccount  " , proto.TcpProtoIDFbGetAllGrayAccount)
				res = x[string(connid)+"_"+string(proto.TcpProtoIDFbGetAllGrayAccount)]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 1):
				fmt.Println(serverId, "  存在,超时客户端无返回值  GetAllGrayAccount  " , proto.TcpProtoIDFbGetAllGrayAccount)
				res = `[]`
				bw := []byte(res)
				w.Write(bw)
			}
		}else {
			res = `[]`
			bw := []byte(res)
			w.Write(bw)
		}
	}
}

func DelGrayAccountById(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		serverZoneId := r.FormValue("serverZoneId")
		gameId := r.FormValue("gameId")
		serverId := r.FormValue("serverId")
		platForm := r.FormValue("platForm")
		account := r.FormValue("account")
		JsonStr := `{"serverZoneId":"` + serverZoneId + `","gameId":"` + gameId + `","serverId":"` + serverId + `","platForm":"` + platForm + `","account":"` + account + `"}`
		conn, exists := gomiddle.ConnMap[serverId]
		var res string
		if exists {
			fmt.Println(r.FormValue("serverId"), "  存在   ", conn)
			connid, _ := gomiddle.ConnMa[serverId]
			conn.Send(connid, makeNoticeMsg(JsonStr,proto.TcpProtoIDFbDelGrayAccountById))	
			
			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  DelGrayAccountById  ",proto.TcpProtoIDFbDelGrayAccountById)
				res = x[string(connid)+"_"+string(proto.TcpProtoIDFbDelGrayAccountById)]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 1):
				fmt.Println(serverId, "  存在,超时客户端无返回值  DelGrayAccountById  ",proto.TcpProtoIDFbDelGrayAccountById)
				res = `{"message":"error"}`
				bw := []byte(res)
				w.Write(bw)
			}
		} else {
			res = `{"message":"error"}`
			bw := []byte(res)
			w.Write(bw)
		}
	}
}

func GetGrayAccountByAccountId(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		serverZoneId := r.FormValue("serverZoneId")
		gameId := r.FormValue("gameId")
		serverId := r.FormValue("serverId")
		id := r.FormValue("id")
		JsonStr := `{"serverZoneId":"` + serverZoneId + `","gameId":"` + gameId + `","serverId":"` + serverId + `","id":"` + id + `"}`
		conn, exists := gomiddle.ConnMap[serverId]
		var res string
		if exists {
			fmt.Println(r.FormValue("serverId"), "  存在   ", conn)
			connid, _ := gomiddle.ConnMa[serverId]
			conn.Send(connid, makeNoticeMsg(JsonStr,proto.TcpProtoIDFbGetGrayAccountById))	

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  GetGrayAccountById ",proto.TcpProtoIDFbGetGrayAccountById)
				res = x[string(connid)+"_"+string(proto.TcpProtoIDFbGetGrayAccountById)]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 1):
				fmt.Println(serverId, "  存在,超时客户端无返回值  GetGrayAccountById ",proto.TcpProtoIDFbGetGrayAccountById)
				res = `{}`
				bw := []byte(res)
				w.Write(bw)
			}
		} else {
			res = `{}`
			bw := []byte(res)
			w.Write(bw)
		}
	}
}

func AddGrayAccount(w http.ResponseWriter, r *http.Request) {
	AddOrUpdateGrayAccount(proto.TcpProtoIDFbSaveGrayAccount, w, r)
}

func UpdateGrayAccount(w http.ResponseWriter, r *http.Request) {
	AddOrUpdateGrayAccount(proto.TcpProtoIDFbUpdateGrayAccount, w, r)
}


func AddOrUpdateGrayAccount(m uint16, w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		//结构已知，解析到结构体
		var s GrayAccountEntity
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
					var responseList entity.ResponseList
			
					if err := json.Unmarshal([]byte(x[string(connid)+"_"+string(m)]), &responseList); err == nil {
						if len(responseList.ObjFail) != 0 {
							objF = responseList.ObjFail[0]
						}
						res = `{"choose":"` + responseList.Choose + `","success":"` + responseList.Success + `","objFail":"` + objF + `","fail":"` + responseList.Fail + `"}`
					}
					
				case <-time.After(time.Second * 1):
					fmt.Println(key, "  存在,超时客户端无返回值  AddOrUpdate ",m)					
					res = `{"choose":"1","success":"0","objFail":"` + key + `","fail":"1"}`
					
				}
			} else {
				fmt.Println(key, "  不存在  ")
				res = `{"choose":"1","success":"0","objFail":"` + key + `","fail":"1"}`
			}
		}
		b := []byte(res)
		w.Write(b)
	}
}

func TcpProtoIDFbGrayGetTotalByServerZoneIdAndGameId(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		serverZoneId := r.FormValue("serverZoneId")
		gameId := r.FormValue("gameId")
		category := r.FormValue("category")
		serverId := r.FormValue("serverId")
		JsonStr := `{"serverZoneId":"` + serverZoneId + `","gameId":"` + gameId + `","category":"` + category + `","serverId":"` + serverId + `"}`

		conn, exists := gomiddle.ConnMap[serverId]
		var res string
		if exists {
			fmt.Println(r.FormValue("serverId"), "  存在   ", conn)
			
			connid, _ := gomiddle.ConnMa[serverId]
			conn.Send(connid, makeNoticeMsg(JsonStr,proto.TcpProtoIDFbGrayGetTotalByServerZoneIdAndGameId))	

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  TcpProtoIDFbGrayGetTotalByServerZoneIdAndGameId  ",proto.TcpProtoIDFbGrayGetTotalByServerZoneIdAndGameId)
				res = x[string(connid)+"_"+string(proto.TcpProtoIDFbGrayGetTotalByServerZoneIdAndGameId)]
				bw := []byte(res)
			    w.Write(bw)
			case <-time.After(time.Second * 1):
				fmt.Println(serverId, "  存在,超时客户端无返回值  TcpProtoIDFbGrayGetTotalByServerZoneIdAndGameId  ",proto.TcpProtoIDFbGrayGetTotalByServerZoneIdAndGameId)
				res = `{"num":0}`
				bw := []byte(res)
				w.Write(bw)
			}
		} else {
			res = `{"num":0}`
			bw := []byte(res)
			w.Write(bw)
		}
	}
}
