package gomiddle

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
	"strings"
	"../../gomiddle"
	entity "../../entity"
	proto "../../tutorial/tcp"
)

type Ann struct{
	ItemId string
	ItemNum string
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
	http.HandleFunc("/xyjserver/email/getAllEmails", GetAllEmails)
	http.HandleFunc("/xyjserver/email/addEmail", AddEmail)
	http.HandleFunc("/xyjserver/email/updateEmail", UpdateEmail)
	http.HandleFunc("/xyjserver/email/delEmailById", DelEmailById)
	http.HandleFunc("/xyjserver/email/getEmailById", GetEmailById)
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
			connid, _ := gomiddle.ConnMa[serverId]
			conn.Send(connid, makeNoticeMsg(JsonStr,proto.TcpProtoIDXyjGetAllEmails))	
			
			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  getAllEmails ",proto.TcpProtoIDXyjGetAllEmails)
				res = x[string(connid)+"_"+string(proto.TcpProtoIDXyjGetAllEmails)]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 3):
				fmt.Println(serverId, "  存在,超时客户端无返回值  getAllEmails ",proto.TcpProtoIDXyjGetAllEmails)
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
	AddOrUpdateEmail(proto.TcpProtoIDXyjAddEmail, w, r)
}

func UpdateEmail(w http.ResponseWriter, r *http.Request){
	AddOrUpdateEmail(proto.TcpProtoIDXyjUpdateEmail, w, r)
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
			connid, _ := gomiddle.ConnMa[serverId]
			conn.Send(connid, makeNoticeMsg(JsonStr,proto.TcpProtoIDXyjGetEmailById))	

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  getEmailById ",proto.TcpProtoIDXyjGetEmailById)
				res = x[string(connid)+"_"+string(proto.TcpProtoIDXyjGetEmailById)]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 3):
				fmt.Println(serverId, "  存在,超时客户端无返回值  getEmailById ",proto.TcpProtoIDXyjGetEmailById)
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
			connid, _ := gomiddle.ConnMa[serverId]
			conn.Send(connid, makeNoticeMsg(JsonStr,proto.TcpProtoIDXyjDelEmailById))	

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  delEmailById ",proto.TcpProtoIDXyjDelEmailById)
				res = x[string(connid)+"_"+string(proto.TcpProtoIDXyjDelEmailById)]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 3):
				fmt.Println(serverId, "  存在,超时客户端无返回值  delEmailById ",proto.TcpProtoIDXyjDelEmailById)
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


func AddOrUpdateEmail(m uint16, w http.ResponseWriter, r *http.Request) {
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
					
				case <-time.After(time.Second * 3):
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

