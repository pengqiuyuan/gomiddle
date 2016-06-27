package gomiddle

import (
    "fmt"
    "database/sql"
    "log"
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