package main

import (
    "fmt"
    "sync"
)

var m sync.Map

/*
sync.Map里每个key只写一次，属于场景1
*/
func changeMap(key int) {
    m.Store(key, 1)
}

func main() {
    var wg sync.WaitGroup
    size := 2
    wg.Add(size)
    
    for i:=0; i<size; i++ {
        i := i
        go func() {
            defer wg.Done()
            changeMap(i)
        }()
    }
    wg.Wait()
    
    /*使用sync.Map里的Range遍历map*/
    m.Range(func(key, value interface{}) bool{
        fmt.Println(key, value)
        return true
    })
}