package main

import (
    "github.com/gogf/gf/g/util/gconv"
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    t  := gconv.Time(now.UnixNano()/100)
    fmt.Println(now.UnixNano())
    fmt.Println(t.Nanosecond())
    fmt.Println(now.Nanosecond())
}