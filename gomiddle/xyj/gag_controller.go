package gomiddle

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
	"../../gomiddle"
	proto "../../tutorial/tcp"
)


type GagEntity struct{
	ServerZoneId string
	GameId string
	ServerId string
	Guid string
	GagTime string
	GagStart string
	GagEnd string
}


//禁言
func GagHandler() {
	http.HandleFunc("/xyjserver/gag/getAllGagAccount", GetAllGagAccount)
	http.HandleFunc("/xyjserver/gag/addGagAccount", AddGagAccount)
	http.HandleFunc("/xyjserver/gag/updateGagAccount", UpdateGagAccount)
	http.HandleFunc("/xyjserver/gag/delGagAccountById", DelGagAccountById)
	http.HandleFunc("/xyjserver/gag/getTotalByServerZoneIdAndGameId", TcpProtoIDXyjGagGetTotalByServerZoneIdAndGameId)	
}


func GetAllGagAccount(w http.ResponseWriter, r *http.Request){
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
			conn.Send(connid, makeNoticeMsg(JsonStr,proto.TcpProtoIDXyjGetAllGagAccount))	

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  getAllGagAccount  ",proto.TcpProtoIDXyjGetAllGagAccount)
				res = x[string(connid)+"_"+string(proto.TcpProtoIDXyjGetAllGagAccount)]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 4):
				fmt.Println(serverId, "  存在,超时客户端无返回值  getAllGagAccount ",proto.TcpProtoIDXyjGetAllGagAccount)
				res = `[]`
				bw := []byte(res)
				w.Write(bw)
			}
		} else {
			res = `[]`
			bw := []byte(res)
			w.Write(bw)
		}
	}
}

func AddGagAccount(w http.ResponseWriter, r *http.Request){
	AddOrUpdateGag(proto.TcpProtoIDXyjAddGagAccount, w, r)
}

func UpdateGagAccount(w http.ResponseWriter, r *http.Request){
	AddOrUpdateGag(proto.TcpProtoIDXyjUpdateGagAccount, w, r)
}

func DelGagAccountById(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		serverZoneId := r.FormValue("serverZoneId")
		gameId := r.FormValue("gameId")
		serverId := r.FormValue("serverId")
		guid := r.FormValue("guid")
		JsonStr := `{"serverZoneId":"` + serverZoneId + `","gameId":"` + gameId + `","serverId":"` + serverId + `","guid":"` + guid + `"}`
		conn, exists := gomiddle.ConnMap[serverId]
		var res string
		if exists {
			fmt.Println(r.FormValue("serverId"), "  存在   ", conn)
			connid, _ := gomiddle.ConnMa[serverId]
			conn.Send(connid, makeNoticeMsg(JsonStr,proto.TcpProtoIDXyjDelGagAccountById))	

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  delGagAccountById ",proto.TcpProtoIDXyjDelGagAccountById)
				res = x[string(connid)+"_"+string(proto.TcpProtoIDXyjDelGagAccountById)]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 4):
				fmt.Println(serverId, "  存在,超时客户端无返回值  delGagAccountById ",proto.TcpProtoIDXyjDelGagAccountById)
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

func AddOrUpdateGag(m uint16, w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		//结构已知，解析到结构体
		var s GagEntity
		json.Unmarshal([]byte(result), &s)
		ser := s.ServerId
		//判断serverid是否在ConnMap里
		fmt.Println("22222  "  + string(result))
		conn, exists := gomiddle.ConnMap[ser]
		var res string
		if exists {
			fmt.Println(ser, "  存在   ", conn)
			connid, _ := gomiddle.ConnMa[ser]
			conn.Send(connid, makeNoticeMsg(string(result),m))
			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(ser, "  存在,客户端有返回值  AddOrUpdate ",m)
				res = x[string(connid)+"_"+string(m)]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 4):
				fmt.Println(ser, "  存在,超时客户端无返回值  AddOrUpdate ",m)
				res = `{"message":"error"}`
				bw := []byte(res)
				w.Write(bw)
			}
		} else {
			fmt.Println(ser, "  不存在  ")
			res = `{"message":"error"}`
			bw := []byte(res)
			w.Write(bw)
		}

	}
}

func TcpProtoIDXyjGagGetTotalByServerZoneIdAndGameId(w http.ResponseWriter, r *http.Request) {
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
			conn.Send(connid, makeNoticeMsg(JsonStr,proto.TcpProtoIDXyjGagGetTotalByServerZoneIdAndGameId))	

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  TcpProtoIDXyjGagGetTotalByServerZoneIdAndGameId  ",proto.TcpProtoIDXyjGagGetTotalByServerZoneIdAndGameId)
				res = x[string(connid)+"_"+string(proto.TcpProtoIDXyjGagGetTotalByServerZoneIdAndGameId)]
				bw := []byte(res)
			    w.Write(bw)
			case <-time.After(time.Second * 4):
				fmt.Println(serverId, "  存在,超时客户端无返回值  TcpProtoIDXyjGagGetTotalByServerZoneIdAndGameId  ",proto.TcpProtoIDXyjGagGetTotalByServerZoneIdAndGameId)
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
