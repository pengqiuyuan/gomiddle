package main
 
import (
    "fmt"
    "time"
)
 
func main() {
    f()
    fmt.Println("end")
}
 
func f() {
    defer func() { 
        fmt.Println("xiaorui.cc start")
        if err := recover(); err != nil {
            fmt.Println(err)
        }
        fmt.Println("xiaorui.cc end")
    }()
    for {
        fmt.Println("1")
        a := []string{"a", "b"}
        fmt.Println(a[3]) 
        //panic("bug")  
	fmt.Println("4")
        time.Sleep(1 * time.Second)
    }
    fmt.Println("lalsalslasas")
}
