package main

import (
    "fmt"
    "sync"
)

func main() {
    /*统计字符串里每个字符出现的次数*/
    m := sync.Map{}
    str := "abcabcd"
    for _, value := range str {
        temp, ok := m.Load(value)
        //fmt.Println(temp, ok)
        if !ok {
            m.Store(value, 1)
        } else {
            /*temp是个interface变量，要转int才能和1做加法*/
            m.Store(value, temp.(int)+1)
        }
    }
    
    /*使用sync.Map里的Range遍历map*/
    m.Range(func(key, value interface{}) bool{
        fmt.Println(key, value)
        return true
    })
}