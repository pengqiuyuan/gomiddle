
package main

import (
    "fmt" 
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
    "log"
	"sync"
	fb "./gomiddle/fb"
	gomiddle "./gomiddle"
)

func main() {
	var wg sync.WaitGroup
	
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

	wg.Add(2)
	go fb.PlacardHandler() 
	go gomiddle.TcpCon(db)
   
	wg.Wait()
	fmt.Println("exit")
}




