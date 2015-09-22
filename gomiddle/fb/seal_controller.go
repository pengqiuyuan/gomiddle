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
	http.HandleFunc("/fbserver/seal/getAllSealAccount", GetAllSealAccount)
	http.HandleFunc("/fbserver/seal/addSealAccount", AddSealAccount)
	http.HandleFunc("/fbserver/seal/updateSealAccount", UpdateSealAccount)
	http.HandleFunc("/fbserver/seal/delSealAccount", DelSealAccount)
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
			b, err := codec.Encode(conn.RemoteAddr().String() + "|getAllSealAccount|" + string(JsonStr) + "|get")
			if err != nil {
				fmt.Println(err)
			}
			conn.Write(b)
			//x := <-gomiddle.Channel_c
			//res := x[conn.RemoteAddr().String()+"_getAllPlacards"]

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  getAllSealAccount")
				res = x[conn.RemoteAddr().String()+"_getAllSealAccount"]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 1):
				fmt.Println(serverId, "  存在,超时客户端无返回值  getAllSealAccount")
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
	AddOrUpdateSeal("addSealAccount", w, r)
}

func UpdateSealAccount(w http.ResponseWriter, r *http.Request) {
	AddOrUpdateSeal("updateSealAccount", w, r)
}

func DelSealAccount(w http.ResponseWriter, r *http.Request) {
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
			b, err := codec.Encode(conn.RemoteAddr().String() + "|delSealAccount|" + string(JsonStr) + "|delete")
			if err != nil {
				fmt.Println(err)
			}
			conn.Write(b)
			//x := <-gomiddle.Channel_c
			//res := x[conn.RemoteAddr().String()+"_delPlacardById"]

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  delSealAccount")
				res = x[conn.RemoteAddr().String()+"_delSealAccount"]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 1):
				fmt.Println(serverId, "  存在,超时客户端无返回值  delSealAccount")
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

func AddOrUpdateSeal(m string, w http.ResponseWriter, r *http.Request) {
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
