
package main

import (
    "bufio"
    "fmt"
    "net" 
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "encoding/json"
    "net/http"    
    "io/ioutil"
	"strings"
	"./gomiddle"
)

type ServerInfoJson struct {
    ServerZoneId int `json:"serverZoneId"` 
    PlatForm []string `json:"platForm"` 
    ServerId string `json:"serverId"` 
    StoreId int `json:"storeId"` 
    Ip string `json:"ip"` 
    Port string `json:"port"` 
    Status string `json:"status"` 
}

   
// 用来记录所有的客户端连接
var ConnMap map[string]*net.TCPConn   

func main() {
    go savePlacard() 

    db, err := sql.Open("mysql", "root:123456@tcp(10.0.29.251:3306)/game_server?charset=utf8")
    if err != nil {
        fmt.Println("mysql init failed")
        return
    }else{
        fmt.Println("mysql init ok")
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    var tcpAddr *net.TCPAddr
    ConnMap = make(map[string]*net.TCPConn)
    tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

    tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

    defer tcpListener.Close()

    for {
        tcpConn, err := tcpListener.AcceptTCP()
        if err != nil {
            continue
        }

        fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
        go tcpPipe(tcpConn,db)
    }

}

func tcpPipe(conn *net.TCPConn,db *sql.DB) {
    ipStr := conn.RemoteAddr().String()
    defer func() {
        fmt.Println("disconnected :" + ipStr)
        conn.Close()
    }()
    reader := bufio.NewReader(conn)

    for {
        message, err := reader.ReadString('\n')
        if err != nil {
            return
        }
        fmt.Println(string(message))
        var jsonServer ServerInfoJson
        if err := json.Unmarshal([]byte(message), &jsonServer); err == nil {  
		     // 新连接加入map
       		 ConnMap[jsonServer.ServerId] = conn
            fmt.Printf("-->运营大区:%s 渠道:%s 服务器:%s 游戏:%s ip:%s 端口:%s 状态:%s\n", jsonServer.ServerZoneId,jsonServer.PlatForm,jsonServer.ServerId,int(jsonServer.StoreId),jsonServer.Ip,jsonServer.Port,jsonServer.Status)
            gomiddle.Insert_serverZone(db,int(jsonServer.ServerZoneId))
            gomiddle.Insert_storeId(db,int(jsonServer.StoreId))
            for  i:=0;i<len(jsonServer.PlatForm);i++ {
                 gomiddle.Insert_all_platform(db,int(jsonServer.ServerZoneId),int(jsonServer.StoreId),jsonServer.PlatForm[i],jsonServer.ServerId)
            }
            gomiddle.Select_all_server(db,int(jsonServer.ServerZoneId),int(jsonServer.StoreId),jsonServer.ServerId,jsonServer.Ip,jsonServer.Port,jsonServer.Status)
        }  
    }
}

/**
 * 保存公告
 */
func savePlacard() {
     http.HandleFunc("/fbserver/placard/addPlacards", SavePlacardHandler)
     http.ListenAndServe(":8899", nil)
}

type PlacardEntity struct {
	GameId string
 	ServerZoneId  string
	ServerId string
	Version string	
  	Contents  string
}

func SavePlacardHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        r.ParseForm()
        result, _:= ioutil.ReadAll(r.Body)
        r.Body.Close()
         //结构已知，解析到结构体
        var s PlacardEntity;
        json.Unmarshal([]byte(result), &s)
        fmt.Println(s);

		//多个serverId按，切分
		ser:= strings.Split(s.ServerId, ",")
        for _,key := range ser {
			//判断serverid是否在ConnMap里
	        value, exists := ConnMap[key]
			if exists {
			  fmt.Println(key,"  存在   ",value)
			
			}else{
			  fmt.Println(key,"  不存在  ")
			}
	    }
		
        jsonStr := `{"choose":1,"success":1,"objFail":["Sample text"],"fail":1}`
        b := []byte(jsonStr+"\n")        
        w.Write(b)
    }
}


