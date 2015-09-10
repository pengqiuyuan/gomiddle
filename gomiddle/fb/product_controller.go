package gomiddle

import (
	"net/http"
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
	http.HandleFunc("/fbserver/seal/getAllProducts", GetAllProducts)
	http.HandleFunc("/fbserver/seal/addProduct", AddProduct)
	http.HandleFunc("/fbserver/seal/updateProduct", UpdateProduct)
	http.HandleFunc("/fbserver/seal/delProductById", DelProductById)
	http.HandleFunc("/fbserver/seal/getProduct", GetProduct)
}


func GetAllProducts(w http.ResponseWriter, r *http.Request){
	
}

func AddProduct(w http.ResponseWriter, r *http.Request){
	
}

func UpdateProduct(w http.ResponseWriter, r *http.Request){
	
}

func DelProductById(w http.ResponseWriter, r *http.Request){
	
}

func GetProduct(w http.ResponseWriter, r *http.Request){
	
}