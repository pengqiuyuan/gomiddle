package gomiddle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"../../codec"
	"../../gomiddle"
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
	http.HandleFunc("/fbserver/gag/getAllGagAccount", GetAllGagAccount)
	http.HandleFunc("/fbserver/gag/addGagAccount", AddGagAccount)
	http.HandleFunc("/fbserver/gag/updateGagAccount", UpdateGagAccount)
	http.HandleFunc("/fbserver/gag/delGagAccountById", DelGagAccountById)
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
			b, err := codec.Encode(conn.RemoteAddr().String() + "|getAllGagAccount|" + string(JsonStr) + "|get")
			if err != nil {
				fmt.Println(err)
			}
			conn.Write(b)
			//x := <-gomiddle.Channel_c
			//res := x[conn.RemoteAddr().String()+"_getAllPlacards"]

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  getAllGagAccount")
				res = x[conn.RemoteAddr().String()+"_getAllGagAccount"]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 1):
				fmt.Println(serverId, "  存在,超时客户端无返回值  getAllGagAccount")
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
	AddOrUpdateGag("addGagAccount", w, r)
}

func UpdateGagAccount(w http.ResponseWriter, r *http.Request){
	AddOrUpdateGag("updateGagAccount", w, r)
}

func DelGagAccountById(w http.ResponseWriter, r *http.Request){
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
			b, err := codec.Encode(conn.RemoteAddr().String() + "|delGagAccountById|" + string(JsonStr) + "|delete")
			if err != nil {
				fmt.Println(err)
			}
			conn.Write(b)

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  delGagAccountById")
				res = x[conn.RemoteAddr().String()+"_delGagAccountById"]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 1):
				fmt.Println(serverId, "  存在,超时客户端无返回值  delGagAccountById")
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

func AddOrUpdateGag(m string, w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		//结构已知，解析到结构体
		var s GagEntity
		json.Unmarshal([]byte(result), &s)
		ser := s.ServerId
		fmt.Println(ser)
		//判断serverid是否在ConnMap里
		conn, exists := gomiddle.ConnMap[ser]
		var res string
		if exists {
			fmt.Println(ser, "  存在   ", conn)
			b,_ := codec.Encode(conn.RemoteAddr().String() + "|" + m + "|" + string(result) + "|post")
			conn.Write(b)
			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(ser, "  存在,客户端有返回值  AddOrUpdate")
				res = x[conn.RemoteAddr().String()+"_"+m]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 1):
				fmt.Println(ser, "  存在,超时客户端无返回值  AddOrUpdate")
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
