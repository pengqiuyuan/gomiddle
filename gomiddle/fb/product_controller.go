package gomiddle

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
	"strings"
	"../../gomiddle"
	proto "../../tutorial/tcp"
	entity "../../entity"
)

type ProductEntity struct{
	ServerZoneId string 
	GameId string
	ServerId string
	ItemId string   //道具ID
	Num string  //个数
	ProdcutStoreId string  //商店ID
	StoreLocation string   //商品出现位置
	IsRandom string   //出现是否随机
	RandomProbability string  //随机概率
	ComsumeType string  //消费类型
	ComsumeNum string    //消费数量
	Discount string   //折扣率
	LevelLimit string   //玩家获取该商品的等级下限
	LevelCap string     //玩家获取该商品的等级上限
	DiscountStartDate string   //折扣生效时间
	DiscountContinueDate string   //折扣持续时间
	DiscountCycleDate string   //折扣循环时间
	ProductPostDate string   //商品上架时间
	ProductDownDate string   //商品下架时间
	ShowLevel string   //显示优先级
}

func ProductHandler() {
	http.HandleFunc("/fbserver/product/getAllProducts", GetAllProducts)
	http.HandleFunc("/fbserver/product/addProduct", AddProduct)
	http.HandleFunc("/fbserver/product/updateProduct", UpdateProduct)
	http.HandleFunc("/fbserver/product/delProductById", DelProductById)
	http.HandleFunc("/fbserver/product/getProduct", GetProduct)
}


func GetAllProducts(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		serverZoneId := r.FormValue("serverZoneId")
		gameId := r.FormValue("gameId")
		serverId := r.FormValue("serverId")
		pageNumber := r.FormValue("pageNumber")
		pageSize := r.FormValue("pageSize")
		itemId := r.FormValue("itemId")
		JsonStr := `{
					"serverZoneId":"` + serverZoneId + `"
					,"gameId":"` + gameId + `"
					,"serverId":"` + serverId + `"
					,"pageNumber":"` + pageNumber + `"
					,"pageSize":"` + pageSize + `"
					,"itemId":"` + itemId + `"
					}`
		conn, exists := gomiddle.ConnMap[serverId]
		var res string
		if exists {
			fmt.Println(r.FormValue("serverId"), "  存在   ", conn)
			connid, _ := gomiddle.ConnMa[serverId]
			conn.Send(connid, makeNoticeMsg(JsonStr,proto.TcpProtoIDFbGetAllProducts))	

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  getAllProducts ",proto.TcpProtoIDFbGetAllProducts)
				res = x[string(connid)+"_"+string(proto.TcpProtoIDFbGetAllProducts)]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 1):
				fmt.Println(serverId, "  存在,超时客户端无返回值  getAllProducts ",proto.TcpProtoIDFbGetAllProducts)
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

func AddProduct(w http.ResponseWriter, r *http.Request){
	AddOrUpdateProduct(proto.TcpProtoIDFbAddProduct, w, r)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request){
	AddOrUpdateProduct(proto.TcpProtoIDFbUpdateProduct, w, r)
}

func DelProductById(w http.ResponseWriter, r *http.Request){
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
			conn.Send(connid, makeNoticeMsg(JsonStr,proto.TcpProtoIDFbDelProductById))	

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  delProductById ",proto.TcpProtoIDFbDelProductById)
				res = x[string(connid)+"_"+string(proto.TcpProtoIDFbDelProductById)]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 3):
				fmt.Println(serverId, "  存在,超时客户端无返回值  delProductById ",proto.TcpProtoIDFbDelProductById)
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

func GetProduct(w http.ResponseWriter, r *http.Request){
	
}


func AddOrUpdateProduct(m uint16, w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		//结构已知，解析到结构体
		var s ProductEntity
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
