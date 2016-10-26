package gomiddle

import (
    "fmt"
    "database/sql"
    "log"
	"encoding/json"
)



func Test(db *sql.DB) {
    tx, err := db.Begin()
    if err != nil {
        log.Fatal(err)
    }
    defer tx.Rollback()
    stmt, err := tx.Prepare("INSERT INTO game_go_serverzone(server_zone_id) VALUES(?)")
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close() 
    for i := 0; i < 10; i++ {
        _, err = stmt.Exec(i)
        if err != nil {
            log.Fatal(err)
        }
    }
    err = tx.Commit()
    if err != nil {
        log.Fatal(err)
    }
}

func Insert_serverZone(db *sql.DB,serverZoneId int,gameId int) {
    tx, err := db.Begin()
    if err != nil {
        log.Fatal(err)
    }
    defer tx.Rollback()

    stmt, err := tx.Prepare("INSERT INTO game_go_serverzone(server_zone_id,store_id) VALUES(?,?)")
    defer stmt.Close()
    if err != nil {
        log.Fatal(err)
        return
    }
    _,err = stmt.Exec(serverZoneId,gameId)
    err = tx.Commit()

}

func Insert_gameId(db *sql.DB,gameId int) {
    tx, err := db.Begin()
    if err != nil {
        log.Fatal(err)
    }
    defer tx.Rollback()

    stmt, err := tx.Prepare("INSERT INTO game_go_store(store_id) VALUES(?)")
    defer stmt.Close()
    if err != nil {
        log.Fatal(err)
        return
    }

    _,err = stmt.Exec(gameId)
    err = tx.Commit()
}

func Delete_server(db *sql.DB,ip string,port string) {
    tx, err := db.Begin()
    if err != nil {
        log.Fatal(err)
    }
    defer tx.Rollback()

    stmt, err := tx.Prepare("DELETE FROM game_go_all_server WHERE ip=? and port=?")
    defer stmt.Close()
    if err != nil {
        log.Fatal(err)
        return
    }

    _,err = stmt.Exec(ip,port)
    err = tx.Commit()
}

func Delete_platform(db *sql.DB,ip string,port string) {
    tx, err := db.Begin()
    if err != nil {
        log.Fatal(err)
    }
    defer tx.Rollback()

    stmt, err := tx.Prepare("DELETE FROM game_go_all_platform WHERE ip=? and port=?")
    defer stmt.Close()
    if err != nil {
        log.Fatal(err)
        return
    }

    _,err = stmt.Exec(ip,port)
    err = tx.Commit()
}

func Truncate_server(db *sql.DB) {
    tx, err := db.Begin()
    if err != nil {
        log.Fatal(err)
    }
    defer tx.Rollback()

    stmt, err := tx.Prepare("TRUNCATE TABLE game_go_all_server")
    defer stmt.Close()
    if err != nil {
        log.Fatal(err)
        return
    }

    _,err = stmt.Exec()
    err = tx.Commit()
}

func Truncate_platform(db *sql.DB) {
    tx, err := db.Begin()
    if err != nil {
        log.Fatal(err)
    }
    defer tx.Rollback()

    stmt, err := tx.Prepare("TRUNCATE TABLE game_go_all_platform")
    defer stmt.Close()
    if err != nil {
        log.Fatal(err)
        return
    }

    _,err = stmt.Exec()
    err = tx.Commit()
}


func Select_all_server(db *sql.DB,serverZoneId int,gameId int,serverId string,ip string, port string,status string){
    stmt,err :=  db.Prepare("SELECT server_id FROM game_go_all_server WHERE server_id = ?")
    defer stmt.Close()
    if err != nil{
        log.Println(err)
        return
    }

    err = stmt.QueryRow(serverId).Scan(&serverId) // WHERE serverId = 1
    if err != nil {
        fmt.Println("----------查询不存在-------------------------")
        tx, err := db.Begin()
        if err != nil {
            log.Fatal(err)
        }
        defer tx.Rollback()

        stmt, err := tx.Prepare("INSERT INTO game_go_all_server(server_zone_id,store_id,server_id,ip,port,status) VALUES(?,?,?,?,?,?)")
        defer stmt.Close()
        if err != nil {
            log.Println(err)
            return
        }
        stmt.Exec(serverZoneId,gameId,serverId,ip,port,status)
        err = tx.Commit()
    }else{
        tx, err := db.Begin()
        if err != nil {
            log.Fatal(err)
        }
        defer tx.Rollback()

        stmt,err :=  tx.Prepare("UPDATE game_go_all_server g SET g.server_zone_id=?,g.store_id=?,g.ip=?,g.port=?,g.status=?  WHERE g.server_id=?")
        defer stmt.Close()
        if err != nil{
            log.Println(err)
            return
        }

        res, err := stmt.Exec(serverZoneId,gameId,ip,port,status,serverId)
        if err != nil{
            log.Println(res)
            return
        }
        err = tx.Commit()   

    }

}

func Insert_all_platform(db *sql.DB,serverZoneId int,gameId int,platFormId string,serverId string,ip string, port string) {
    tx, err := db.Begin()
    if err != nil {
        log.Fatal(err)
    }
    defer tx.Rollback()

    stmt, err := tx.Prepare("INSERT INTO game_go_all_platform(server_zone_id,store_id,plat_form_id,server_id,ip,port) VALUES(?,?,?,?,?,?)")
    defer stmt.Close()
    if err != nil {
        log.Println(err)
        return
    }
    stmt.Exec(serverZoneId,gameId,platFormId,serverId,ip,port)
    err = tx.Commit()
}

func GetEventJSON(db *sql.DB,serverZoneId int,gameId int) (string, error) {  
    stmt, err := db.Prepare("SELECT CAST(id AS CHAR) id,server_zone_id serverZoneId,game_id gameId,event_type eventType,main_ui_position mainUiPosition,event_pic eventPic,active_type activeType,active_data activeData,role_level_min roleLevelMin,role_level_max roleLevelMax,times times,active_delay activeDelay,active_day activeDay,event_repeat_interval eventRepeatInterval,event_name eventName,event_title eventTitle,event_des eventDes,event_icon eventIcon,list_priority listPriority,done_hiding doneHiding,following_event followingEvent,event_show eventShow  from game_gm_event_prototype  WHERE game_id = ? and server_zone_id = ?")  
    if err != nil {  
        return "", err  
    }  
    defer stmt.Close()  
    rows, err := stmt.Query(gameId,serverZoneId)  
    if err != nil {  
        return "", err  
    }  
    defer rows.Close()  
    columns, err := rows.Columns()  
    if err != nil {  
      return "", err  
    }  
    count := len(columns)  
    tableData := make([]map[string]interface{}, 0)  
    values := make([]interface{}, count)  
    valuePtrs := make([]interface{}, count)  
    for rows.Next() {  
      for i := 0; i < count; i++ {  
          valuePtrs[i] = &values[i]  
      }  
      rows.Scan(valuePtrs...)  
      entry := make(map[string]interface{})  
      for i, col := range columns {  
          var v interface{}  
          val := values[i]  
          b, ok := val.([]byte)  
          if ok {  
              v = string(b)  
          } else {   
              v = val  
          }  
          entry[col] = v
      }  
      tableData = append(tableData, entry)  
    }  
    jsonData, err := json.Marshal(tableData)  
    if err != nil {  
      return "", err  
    }  
    //fmt.Println(string(jsonData))  
    return string(jsonData), nil   
}  

func GetEventDataJSON(db *sql.DB,eventId string) (string, error) {  
    stmt, err := db.Prepare("SELECT CAST(event_data_id AS CHAR) eventDataId,CAST(event_id AS CHAR) eventId,event_group eventGroup,event_data_name eventDataName,vip_min vipMin,vip_max vipMax,event_data_times eventDataTimes,event_data_delay eventDataDelay,event_data_des eventDataDes,event_condition eventCondition,event_condition_type eventConditionType,condition_value1 conditionValue1,condition_value2 conditionValue2,event_rewards eventRewards,event_rewards_num eventRewardsNum from game_gm_event_data_prototype  WHERE event_id = ?")  
    if err != nil {  
        return "", err  
    }  
    defer stmt.Close()  
    rows, err := stmt.Query(eventId)  
    if err != nil {  
        return "", err  
    }  
    defer rows.Close()  
    columns, err := rows.Columns()  
    if err != nil {  
      return "", err  
    }  
    count := len(columns)  
    tableData := make([]map[string]interface{}, 0)  
    values := make([]interface{}, count)  
    valuePtrs := make([]interface{}, count)  
    for rows.Next() {  
      for i := 0; i < count; i++ {  
          valuePtrs[i] = &values[i]  
      }  
      rows.Scan(valuePtrs...)  
      entry := make(map[string]interface{})  
      for i, col := range columns {  
          var v interface{}  
          val := values[i]  
          b, ok := val.([]byte)  
          if ok {  
              v = string(b)  
          } else {   
              v = val  
          }  
          entry[col] = v
      }  
      tableData = append(tableData, entry)  
    }  
    jsonData, err := json.Marshal(tableData)  
    if err != nil {  
      return "", err  
    }  
    //fmt.Println(string(jsonData))  
    return string(jsonData), nil   
}  