package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    a := 1
    b := 2
    wg := sync.WaitGroup{}
    wg.Add(1)
    go func() {
        fmt.Println(a, b)
        wg.Done()
    }()
    time.Sleep(time.Second)
    a = 99
    wg.Wait()
}
