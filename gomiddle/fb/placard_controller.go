package gomiddle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"../../codec"
	entity "../../entity"
	"../../gomiddle"
)

type PlacardEntity struct {
	GameId       string
	ServerZoneId string
	ServerId     string
	ServerIds    []string
	Version      string
	Contents     string
}

func PlacardHandler() {
	http.HandleFunc("/fbserver/placard/getAllPlacards", GetAllPlacards)
	http.HandleFunc("/fbserver/getTotalByServerZoneIdAndGameId", GetTotalByServerZoneIdAndGameId)
	http.HandleFunc("/fbserver/placard/addPlacards", SavePlacardHandler)
	http.HandleFunc("/fbserver/placard/updatePlacards", UpdatePlacards)
	http.HandleFunc("/fbserver/placard/delPlacardById", DelPlacardById)
}

/**
 * 查询运营大区、游戏下所有的公告
 * 参数 localhost:8899/fbserver/placard/getAllPlacards?serverZoneId=1&gameId=1&pageNumber=1&serverId=fb_server_1&pageSize=1
 * 传入格式 127.0.0.1:53038|getAllPlacards|{"serverZoneId":"1","gameId":"1","serverId":"fb_server_1","pageNumber":"1","pageSize":"1"}|get
 */
func GetAllPlacards(w http.ResponseWriter, r *http.Request) {
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
			b, err := codec.Encode(conn.RemoteAddr().String() + "|getAllPlacards|" + string(JsonStr) + "|get")
			if err != nil {
				fmt.Println(err)
			}
			conn.Write(b)
			//x := <-gomiddle.Channel_c
			//res := x[conn.RemoteAddr().String()+"_getAllPlacards"]
			
			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  GetAllPlacards")
				res = x[conn.RemoteAddr().String()+"_getAllPlacards"]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 1):
				fmt.Println(serverId, "  存在,超时客户端无返回值  GetAllPlacards")
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

/**
 * 查询运营大区、游戏下 公告的总数
 * 参数 localhost:8899/fbserver/getTotalByServerZoneIdAndGameId?serverZoneId=1&gameId=1&category=placard&serverId=fb_server_1
 * 传入格式 127.0.0.1:54726|getTotalByServerZoneIdAndGameId|{"serverZoneId":"1","gameId":"1","category":"placard","serverId":"fb_server_1"}|get
 * 返回格式 map key 127.0.0.1:54726_getTotalByServerZoneIdAndGameId    value {"num":1}
 */
func GetTotalByServerZoneIdAndGameId(w http.ResponseWriter, r *http.Request) {
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
			b, err := codec.Encode(conn.RemoteAddr().String() + "|getTotalByServerZoneIdAndGameId|" + string(JsonStr) + "|get")
			if err != nil {
				fmt.Println(err)
			}
			conn.Write(b)
			//x := <-gomiddle.Channel_c
			//res := x[conn.RemoteAddr().String()+"_getTotalByServerZoneIdAndGameId"]

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  GetTotalByServerZoneIdAndGameId")
				res = x[conn.RemoteAddr().String()+"_getTotalByServerZoneIdAndGameId"]
				bw := []byte(res)
			    w.Write(bw)
			case <-time.After(time.Second * 1):
				fmt.Println(serverId, "  存在,超时客户端无返回值  GetTotalByServerZoneIdAndGameId")
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

/**
 * 保存公告
 * 传入数据格式 127.0.0.1:54726|addPlacards|{"id":2,"serverZoneId":"1","contents":"<p>\r\n\t1<\/p>\r\n","gameId":"1","serverId":"fb_server_1,fb_server_2,fb_server_3","serverIds":null,"version":"1"}|post
 * 返回数据格式 127.0.0.1:54813|addPlacards|{"choose":1,"success":1,"objFail":[],"fail":0}|post
 */
func SavePlacardHandler(w http.ResponseWriter, r *http.Request) {
	AddOrUpdate("addPlacards", w, r)
}

func UpdatePlacards(w http.ResponseWriter, r *http.Request) {
	AddOrUpdate("updatePlacards", w, r)
}

/**
 * 根据id删除公告
 * 参数 localhost:8899/fbserver/placard/delPlacardById?id=1&serverZoneId=1&gameId=1&serverId=fb_server_1
 * 传入格式 127.0.0.1:53340|delPlacardById|{"serverZoneId":"1","gameId":"1","serverId":"fb_server_1","id":"1"}|delete
 */
func DelPlacardById(w http.ResponseWriter, r *http.Request) {
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
			b, err := codec.Encode(conn.RemoteAddr().String() + "|delPlacardById|" + string(JsonStr) + "|delete")
			if err != nil {
				fmt.Println(err)
			}
			conn.Write(b)
			//x := <-gomiddle.Channel_c
			//res := x[conn.RemoteAddr().String()+"_delPlacardById"]

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  DelPlacardById")
				res = x[conn.RemoteAddr().String()+"_delPlacardById"]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 1):
				fmt.Println(serverId, "  存在,超时客户端无返回值  DelPlacardById")
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

func AddOrUpdate(m string, w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		//结构已知，解析到结构体
		var s PlacardEntity
		json.Unmarshal([]byte(result), &s)
		//fmt.Println(s)
		//多个serverId按，切分
		//ser := strings.Split(s.ServerId, ",")

		ser := s.ServerIds
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
				case <-time.After(time.Second * 1):
					fmt.Println(key, "  存在,超时客户端无返回值  AddOrUpdate")
					fail = fail + 1
					objFail = append(objFail, key)
					/*
						default:
							fmt.Println(key, "  存在,客户端无返回值  ")
							fail = fail + 1
							objFail = append(objFail, key)
					*/
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
