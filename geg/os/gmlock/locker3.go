package main

import (
    "sync"
    "github.com/gogf/gf/g/os/glog"
    "time"
    "github.com/gogf/gf/g/os/gmlock"
)

// 内存锁 - TryLock
func main() {
    key := "lock"
    wg  := sync.WaitGroup{}
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            if gmlock.TryLock(key) {
                glog.Println(i)
                time.Sleep(time.Second)
                gmlock.Unlock(key)
            } else {
                glog.Println(false)
            }
            wg.Done()
        }(i)
    }
    wg.Wait()
}
