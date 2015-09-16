package gomiddle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"../../codec"
	"../../gomiddle"
	"strings"
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
			b, err := codec.Encode(conn.RemoteAddr().String() + "|getAllProducts|" + string(JsonStr) + "|get")
			if err != nil {
				fmt.Println(err)
			}
			conn.Write(b)
			//x := <-gomiddle.Channel_c
			//res := x[conn.RemoteAddr().String()+"_getAllPlacards"]

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  getAllProducts")
				res = x[conn.RemoteAddr().String()+"_getAllProducts"]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 1):
				fmt.Println(serverId, "  存在,超时客户端无返回值  getAllProducts")
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
	AddOrUpdateProduct("addProduct", w, r)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request){
	AddOrUpdateProduct("updateProduct", w, r)
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
			b, err := codec.Encode(conn.RemoteAddr().String() + "|delProductById|" + string(JsonStr) + "|delete")
			if err != nil {
				fmt.Println(err)
			}
			conn.Write(b)

			select {
			case x := <-gomiddle.Channel_c:
				fmt.Println(serverId, "  存在,客户端有返回值  delProductById")
				res = x[conn.RemoteAddr().String()+"_delProductById"]
				bw := []byte(res)
				w.Write(bw)
			case <-time.After(time.Second * 3):
				fmt.Println(serverId, "  存在,超时客户端无返回值  delProductById")
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


func AddOrUpdateProduct(m string, w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		//结构已知，解析到结构体
		var s ProductEntity
		json.Unmarshal([]byte(result), &s)
		ser := strings.Split(s.ServerId,",") 
		fmt.Println(ser)
				
		for _, key := range ser {
			//判断serverid是否在ConnMap里
			conn, exists := gomiddle.ConnMap[key]
			var res string
			if exists {
				fmt.Println(key, "  存在   ", conn)
				b,_ := codec.Encode(conn.RemoteAddr().String() + "|" + m + "|" + string(result) + "|post")
				conn.Write(b)
				select {
				case x := <-gomiddle.Channel_c:
					fmt.Println(key, "  存在,客户端有返回值  AddOrUpdate")
					res = x[conn.RemoteAddr().String()+"_"+m]
					bw := []byte(res)
					w.Write(bw)
				case <-time.After(time.Second * 1):
					fmt.Println(key, "  存在,超时客户端无返回值  AddOrUpdate")
					res = `{"message":"error"}`
					bw := []byte(res)
					w.Write(bw)
				}
			} else {
				fmt.Println(key, "  不存在  ")
				res = `{"message":"error"}`
				bw := []byte(res)
				w.Write(bw)
			}
		}

	}
}
