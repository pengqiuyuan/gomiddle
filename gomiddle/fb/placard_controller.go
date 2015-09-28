package gomiddle

import (
	"fmt"
	"net/http"
	"time"
	"log"
	"io/ioutil"
	"strings"
	"encoding/json"
	"../../gomiddle"
	"enlightgame/transport"
	entity "../../entity"
	proto "../../tutorial/tcp"
	flatbuffers "github.com/google/flatbuffers/go"
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
	http.HandleFunc("/fbserver/placard/addPlacards", SavePlacard)
	http.HandleFunc("/fbserver/placard/updatePlacards", UpdatePlacards)
	http.HandleFunc("/fbserver/placard/delPlacardById", DelPlacardById)
}

func makeNoticeMsg(str string,p uint16) []byte {
	t := transport.TcpMessage{}

	builder := flatbuffers.NewBuilder(0)
	
	ct := builder.CreateString(str)
	proto.NoticeStart(builder)
	proto.NoticeAddContent(builder, ct)
	payload := proto.NoticeEnd(builder)
	
	builder.Finish(payload)

	t.Payload = builder.Bytes[builder.Head():]

	// 填充协议头信息
	t.Header.Proto = p
	t.Header.Flag = 0xdcba
	t.Header.Size = uint16(len(t.Payload))

	ret, err := t.Pack()
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return ret
}

/**
 * 查询运营大区、游戏下所有的公告
 * 参数 localhost:8899/fbserver/placard/getAllPlacards?serverZoneId=1&gameId=1&pageNumber=1&serverId=fb_server_1&pageSize=1
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
			connid, _ := gomiddle.ConnMa[serverId]

			conn.Send(connid, makeNoticeMsg(JsonStr,proto.TcpProtoIDFbGetAllPlacards))	
			
	
			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  getAllPlacards  " , proto.TcpProtoIDFbGetAllPlacards)
				res = x[string(connid)+"_"+string(proto.TcpProtoIDFbGetAllPlacards)]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 1):
				fmt.Println(serverId, "  存在,超时客户端无返回值  getAllPlacards  " , proto.TcpProtoIDFbGetAllPlacards)
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
			
			connid, _ := gomiddle.ConnMa[serverId]
			conn.Send(connid, makeNoticeMsg(JsonStr,proto.TcpProtoIDFbGetTotalByServerZoneIdAndGameId))	

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  GetTotalByServerZoneIdAndGameId  ",proto.TcpProtoIDFbGetTotalByServerZoneIdAndGameId)
				res = x[string(connid)+"_"+string(proto.TcpProtoIDFbGetTotalByServerZoneIdAndGameId)]
				bw := []byte(res)
			    w.Write(bw)
			case <-time.After(time.Second * 1):
				fmt.Println(serverId, "  存在,超时客户端无返回值  GetTotalByServerZoneIdAndGameId  ",proto.TcpProtoIDFbGetTotalByServerZoneIdAndGameId)
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
 */
func SavePlacard(w http.ResponseWriter, r *http.Request) {
	AddOrUpdatePlacard(proto.TcpProtoIDFbSavePlacard, w, r)
}

func UpdatePlacards(w http.ResponseWriter, r *http.Request) {
	AddOrUpdatePlacard(proto.TcpProtoIDFbUpdatePlacards, w, r)
}

/**
 * 根据id删除公告
 * 参数 localhost:8899/fbserver/placard/delPlacardById?id=1&serverZoneId=1&gameId=1&serverId=fb_server_1
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
			connid, _ := gomiddle.ConnMa[serverId]
			conn.Send(connid, makeNoticeMsg(JsonStr,proto.TcpProtoIDFbDelPlacardById))	
			
			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  delPlacardById  ",proto.TcpProtoIDFbDelPlacardById)
				res = x[string(connid)+"_"+string(proto.TcpProtoIDFbDelPlacardById)]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 1):
				fmt.Println(serverId, "  存在,超时客户端无返回值  delPlacardById  ",proto.TcpProtoIDFbDelPlacardById)
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

func AddOrUpdatePlacard(m uint16, w http.ResponseWriter, r *http.Request) {
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

		//ser := s.ServerIds
		choose := len(ser)
		success := 0
		fail := 0
		var objFail []string

		for _, key := range ser {
			//判断serverid是否在ConnMap里
			conn, exists := gomiddle.ConnMap[key]
			if exists {
				fmt.Println(key, "  存在   ", conn)
				connid, _ := gomiddle.ConnMa[key]
				conn.Send(connid, makeNoticeMsg(string(result),m))	

				select {
				case x := <-gomiddle.Channel_c:
					fmt.Println(key, "  存在,客户端有返回值  AddOrUpdate  ",m)
					var responseList entity.ResponseList
					if err := json.Unmarshal([]byte(x[string(connid)+"_"+string(m)]), &responseList); err == nil {
						success = success + responseList.Success
						fail = fail + responseList.Fail
						if len(responseList.ObjFail) != 0 {
							objFail = append(objFail, responseList.ObjFail[0])
						}
					}
				case <-time.After(time.Second * 1):
					fmt.Println(key, "  存在,超时客户端无返回值  AddOrUpdate  ",m)
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



