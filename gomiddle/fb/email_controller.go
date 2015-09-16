package gomiddle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"strings"
	"../../codec"
	entity "../../entity"
	"../../gomiddle"
)

type Ann struct{
	ItemId string
	ItemNuM int
}

type EmailEntity struct{
	ServerZoneId string 
	GameId string
	ServerId string
	PlatForm string   
	Sender string 
	Title string  
	Contents string   
	Annex []Ann
}

func EmailHandler() {
	http.HandleFunc("/fbserver/email/getAllEmails", GetAllEmails)
	http.HandleFunc("/fbserver/email/addEmail", AddEmail)
	http.HandleFunc("/fbserver/email/updateEmail", UpdateEmail)
	http.HandleFunc("/fbserver/email/delEmailById", DelEmailById)
	http.HandleFunc("/fbserver/email/getEmailById", GetEmailById)
}


func GetAllEmails(w http.ResponseWriter, r *http.Request){
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
			b, err := codec.Encode(conn.RemoteAddr().String() + "|getAllEmails|" + string(JsonStr) + "|get")
			if err != nil {
				fmt.Println(err)
			}
			conn.Write(b)
			
			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  getAllEmails")
				res = x[conn.RemoteAddr().String()+"_getAllEmails"]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 3):
				fmt.Println(serverId, "  存在,超时客户端无返回值  getAllEmails")
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

func AddEmail(w http.ResponseWriter, r *http.Request){
	AddOrUpdateEmail("addEmail", w, r)
}

func UpdateEmail(w http.ResponseWriter, r *http.Request){
	AddOrUpdateEmail("updateEmail", w, r)
}

func GetEmailById(w http.ResponseWriter, r *http.Request){
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
			b, err := codec.Encode(conn.RemoteAddr().String() + "|getEmailById|" + string(JsonStr) + "|get")
			if err != nil {
				fmt.Println(err)
			}
			conn.Write(b)

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  getEmailById")
				res = x[conn.RemoteAddr().String()+"_getEmailById"]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 3):
				fmt.Println(serverId, "  存在,超时客户端无返回值  getEmailById")
				res = ``
				bw := []byte(res)
				w.Write(bw)
			}
		} else {
			res = ``
			bw := []byte(res)
			w.Write(bw)
		}
	}
}

func DelEmailById(w http.ResponseWriter, r *http.Request){
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
			b, err := codec.Encode(conn.RemoteAddr().String() + "|delEmailById|" + string(JsonStr) + "|delete")
			if err != nil {
				fmt.Println(err)
			}
			conn.Write(b)

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  delEmailById")
				res = x[conn.RemoteAddr().String()+"_delEmailById"]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 3):
				fmt.Println(serverId, "  存在,超时客户端无返回值  delEmailById")
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


func AddOrUpdateEmail(m string, w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		//结构已知，解析到结构体
		var s EmailEntity
		json.Unmarshal([]byte(result), &s)
		fmt.Println(s)
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
				b, err := codec.Encode(conn.RemoteAddr().String() + "|" + m + "|" + string(result) + "|post")
				if err != nil {
					continue
				}
				conn.Write(b)

				select {
				case x := <-gomiddle.Channel_c:
					fmt.Println(key, "  存在,客户端有返回值  AddOrUpdate")
					var responseList entity.ResponseList
					if err := json.Unmarshal([]byte(x[conn.RemoteAddr().String()+"_"+m]), &responseList); err == nil {
						success = success + responseList.Success
						fail = fail + responseList.Fail
						if len(responseList.ObjFail) != 0 {
							objFail = append(objFail, responseList.ObjFail[0])
						}
					}
				case <-time.After(time.Second * 3):
					fmt.Println(key, "  存在,超时客户端无返回值  AddOrUpdate")
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

