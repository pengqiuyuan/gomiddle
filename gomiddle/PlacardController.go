package gomiddle

import (
    "fmt"
    "encoding/json"
    "net/http"    
    "io/ioutil"
	"strings"
	"../codec"
)

/**
 * 保存公告
 */
func SavePlacard() {
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
			  tHeader := &codec.TcpMessageHeader{Flag:uint16(0xdcba),Proto:uint16(60001),Size:uint16(len(result))}			
			  tMessage := &codec.TcpMessage{Header:*tHeader,Payload:[]byte(result)} 		
			  msg, err := codec.Encode(tMessage);
			  if err != nil {
		         return
		      }
			  value.Write(msg)				
			}else{
			  fmt.Println(key,"  不存在  ")
			
			
			
			}
	    }
		
        jsonStr := `{"choose":1,"success":1,"objFail":["Sample text"],"fail":1}`
        b := []byte(jsonStr+"\n")        
        w.Write(b)
    }
}