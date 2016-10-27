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

type EventPrototype struct{
	Id string `json:"id"`
	ServerZoneId string `json:"serverZoneId"`
	GameId string `json:"gameId"`
	ServerId string  `json:"serverId"`
	EventType string    `json:"eventType"`
	ActiveType string `json:"activeType"`
	ActiveData string  `json:"activeData"`
	ActiveDay string   `json:"activeDay"`

	ActiveDelay string   `json:"activeDelay"`
	Times string `json:"times"`
	EventRepeatInterval string  `json:"eventRepeatInterval"`
	FollowingEvent string   `json:"followingEvent"`
	
	RoleLevelMin string   `json:"roleLevelMin"`
	RoleLevelMax string `json:"roleLevelMax"`
	MainUiPosition string  `json:"mainUiPosition"`
	EventTitle string   `json:"eventTitle"`
	
	EventName string   `json:"eventName"`
	EventDes string `json:"eventDes"`
	EventPic string  `json:"eventPic"`
	EventIcon string   `json:"eventIcon"`
	
	ListPriority string   `json:"listPriority"`
	DoneHiding string `json:"doneHiding"`
	EventShow string  `json:"eventShow"`
}

type EventDataPrototype struct{
	ServerId string `json:"serverId"`
	EventDataId string `json:"eventDataId"`
	EventId string `json:"eventId"`
	Group string `json:"group"`
	EventDataName string   `json:"eventDataName"`
	EventDataDes string `json:"eventDataDes"`
	EventDataDelay string  `json:"eventDataDelay"`
	EventDataTimes string   `json:"eventDataTimes"`
	VipMin string   `json:"vipMin"`
	VipMax string `json:"vipMax"`
	EventConditionType string  `json:"eventConditionType"`
	EventCondition string   `json:"eventCondition"`
	ConditionValue1 string   `json:"conditionValue1"`
	ConditionValue2 string `json:"conditionValue2"`
	EventRewards string  `json:"eventRewards"`
	EventRewardsNum string   `json:"eventRewardsNum"`
}

type EventDataPrototypeInstruction struct{
	Id string 
	EventCondition string
	EventConditionType string
	ConditionValue1 string   
	ConditionName1 string 
	ConditionValue2 string  
	ConditionName2 string   
}

func EventHandler() {
	http.HandleFunc("/xyjserver/eventPrototype/addEventPrototype", AddEventPrototype)
	http.HandleFunc("/xyjserver/eventPrototype/updateEventPrototype", UpdateEventPrototype)
	http.HandleFunc("/xyjserver/eventPrototype/addEventDataPrototype", AddEventDataPrototype)
	http.HandleFunc("/xyjserver/eventPrototype/updateEventDataPrototype", UpdateEventDataPrototype)
	http.HandleFunc("/xyjserver/eventPrototype/closeEventPrototype", CloseEventPrototype)
}

func AddEventPrototype(w http.ResponseWriter, r *http.Request){
	AddOrUpdateEventPrototype(proto.TcpProtoIDXyjAddEventPrototype, w, r)
}

func UpdateEventPrototype(w http.ResponseWriter, r *http.Request){
	AddOrUpdateEventPrototype(proto.TcpProtoIDXyjUpdateEventPrototype, w, r)
}

func AddEventDataPrototype(w http.ResponseWriter, r *http.Request){
	AddOrUpdateEventDataPrototype(proto.TcpProtoIDXyjAddEventDataPrototype, w, r)
}

func UpdateEventDataPrototype(w http.ResponseWriter, r *http.Request){
	AddOrUpdateEventDataPrototype(proto.TcpProtoIDXyjUpdateEventDataPrototype, w, r)
}

/*新增、修改活动*/
func AddOrUpdateEventPrototype(m uint16, w http.ResponseWriter, r *http.Request) {
	fmt.Println(" ---------------------------- 开始-----------------------------")
	if r.Method == "POST" {
		r.ParseForm()
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		//结构已知，解析到结构体
		var s EventPrototype
		err := json.Unmarshal([]byte(result), &s)
		if err != nil {
			fmt.Printf(err.Error())
		}
		fmt.Println(s," 新增、修改活动")
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
					fmt.Println(key, "  存在,客户端有返回值  AddOrUpdate 新增、修改活动",m)		
					var responseList entity.ResponseList
			
					if err := json.Unmarshal([]byte(x[string(connid)+"_"+string(m)]), &responseList); err == nil {
						if len(responseList.ObjFail) != 0 {
							objF = responseList.ObjFail[0]
						}
						res = `{"choose":"` + responseList.Choose + `","success":"` + responseList.Success + `","objFail":"` + objF + `","fail":"` + responseList.Fail + `"}`
					}else{
						fmt.Println(err," 存在,客户端有返回值  AddOrUpdate 新增、修改活动，出错了")
						res = `{"choose":"1","success":"0","objFail":"` + key + `","fail":"1"}`
					}
					
				case <-time.After(time.Second * 1):
					fmt.Println(key, "  存在,超时客户端无返回值  AddOrUpdate 新增、修改活动",m)					
					res = `{"choose":"1","success":"0","objFail":"` + key + `","fail":"1"}`
					
				}
			} else {
				fmt.Println(key, "  不存在  ")
				res = `{"choose":"1","success":"0","objFail":"` + key + `","fail":"1"}`
			}
		}
		b := []byte(res)
		w.Write(b)
		fmt.Println(" ---------------------------- 结束-----------------------------")
	}
}

/*新增、修改活动下得条目*/
func AddOrUpdateEventDataPrototype(m uint16, w http.ResponseWriter, r *http.Request) {
	fmt.Println(" ---------------------------- 开始-------------------------------------------------------------")
	if r.Method == "POST" {
		r.ParseForm()
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		//结构已知，解析到结构体
		var s EventDataPrototype
		err := json.Unmarshal([]byte(result), &s)
		if err != nil {
			fmt.Printf(err.Error())
		}
		fmt.Println(s," 新增、修改活动下得条目")
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
					fmt.Println(key, "  存在,客户端有返回值  AddOrUpdate 新增、修改活动下得条目",m)		
					var responseList entity.ResponseList
			
					if err := json.Unmarshal([]byte(x[string(connid)+"_"+string(m)]), &responseList); err == nil {
						if len(responseList.ObjFail) != 0 {
							objF = responseList.ObjFail[0]
						}
						res = `{"choose":"` + responseList.Choose + `","success":"` + responseList.Success + `","objFail":"` + objF + `","fail":"` + responseList.Fail + `"}`
					}else{
						fmt.Println(err," 存在,客户端有返回值  AddOrUpdate 新增、修改活动下得条目，出错了")
						res = `{"choose":"1","success":"0","objFail":"` + key + `","fail":"1"}`
					}
					
				case <-time.After(time.Second * 1):
					fmt.Println(key, "  存在,超时客户端无返回值  AddOrUpdate 新增、修改活动下得条目",m)					
					res = `{"choose":"1","success":"0","objFail":"` + key + `","fail":"1"}`
					
				}
			} else {
				fmt.Println(key, "  不存在  ")
				res = `{"choose":"1","success":"0","objFail":"` + key + `","fail":"1"}`
			}
		}
		b := []byte(res)
		w.Write(b)
		fmt.Println(" ---------------------------- 结束-------------------------------------------------------------")
	}
}

/*关闭活动*/
func CloseEventPrototype(w http.ResponseWriter, r *http.Request){
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
			conn.Send(connid, makeNoticeMsg(JsonStr,proto.TcpProtoIDXyjCloseEventPrototype))	

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  CloseEventPrototype 关闭活动",proto.TcpProtoIDXyjCloseEventPrototype)
				res = x[string(connid)+"_"+string(proto.TcpProtoIDXyjCloseEventPrototype)]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 1):
				fmt.Println(serverId, "  存在,超时客户端无返回值  CloseEventPrototype 关闭活动",proto.TcpProtoIDXyjCloseEventPrototype)
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
