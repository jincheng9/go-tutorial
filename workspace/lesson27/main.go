package main

import (
    "fmt"
    "sync"
    "project/util"
)

func main() {
    var result int
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        defer wg.Done()
        result = util2.Sub(1, 2)
    }()
    wg.Wait()
    fmt.Println(result)
    fmt.Println(util2.Sub(1,2))
}