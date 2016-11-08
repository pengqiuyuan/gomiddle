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

type SealEntity struct {
	ServerZoneId string
	GameId       string
	ServerId     string
	Guid         string
	SealTime      string
	SealStart     string
	SealEnd       string
}

//封号
func SealHandler() {
	http.HandleFunc("/xyjserver/seal/getAllSealAccount", GetAllSealAccount)
	http.HandleFunc("/xyjserver/seal/addSealAccount", AddSealAccount)
	http.HandleFunc("/xyjserver/seal/updateSealAccount", UpdateSealAccount)
	http.HandleFunc("/xyjserver/seal/delSealAccount", DelSealAccount)
	http.HandleFunc("/xyjserver/seal/getTotalByServerZoneIdAndGameId", TcpProtoIDXyjSealGetTotalByServerZoneIdAndGameId)
}

func GetAllSealAccount(w http.ResponseWriter, r *http.Request) {
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
			conn.Send(connid, makeNoticeMsg(JsonStr,proto.TcpProtoIDXyjGetAllSealAccount))	

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  getAllSealAccount ",proto.TcpProtoIDXyjGetAllSealAccount)
				res = x[string(connid)+"_"+string(proto.TcpProtoIDXyjGetAllSealAccount)]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 4):
				fmt.Println(serverId, "  存在,超时客户端无返回值  getAllSealAccount ",proto.TcpProtoIDXyjGetAllSealAccount)
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

func AddSealAccount(w http.ResponseWriter, r *http.Request) {
	AddOrUpdateSeal(proto.TcpProtoIDXyjAddSealAccount, w, r)
}

func UpdateSealAccount(w http.ResponseWriter, r *http.Request) {
	AddOrUpdateSeal(proto.TcpProtoIDXyjUpdateSealAccount, w, r)
}

func DelSealAccount(w http.ResponseWriter, r *http.Request) {
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
			conn.Send(connid, makeNoticeMsg(JsonStr,proto.TcpProtoIDXyjDelSealAccount))	

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  delSealAccount ",proto.TcpProtoIDXyjDelSealAccount)
				res = x[string(connid)+"_"+string(proto.TcpProtoIDXyjDelSealAccount)]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 4):
				fmt.Println(serverId, "  存在,超时客户端无返回值  delSealAccount ",proto.TcpProtoIDXyjDelSealAccount)
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

func AddOrUpdateSeal(m uint16, w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		//结构已知，解析到结构体
		var s SealEntity
		json.Unmarshal([]byte(result), &s)
		ser := s.ServerId
		fmt.Println(ser)
		//判断serverid是否在ConnMap里
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

func TcpProtoIDXyjSealGetTotalByServerZoneIdAndGameId(w http.ResponseWriter, r *http.Request) {
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
			conn.Send(connid, makeNoticeMsg(JsonStr,proto.TcpProtoIDXyjSealGetTotalByServerZoneIdAndGameId))	

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  TcpProtoIDXyjSealGetTotalByServerZoneIdAndGameId  ",proto.TcpProtoIDXyjSealGetTotalByServerZoneIdAndGameId)
				res = x[string(connid)+"_"+string(proto.TcpProtoIDXyjSealGetTotalByServerZoneIdAndGameId)]
				bw := []byte(res)
			    w.Write(bw)
			case <-time.After(time.Second * 4):
				fmt.Println(serverId, "  存在,超时客户端无返回值  TcpProtoIDXyjSealGetTotalByServerZoneIdAndGameId  ",proto.TcpProtoIDXyjSealGetTotalByServerZoneIdAndGameId)
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
