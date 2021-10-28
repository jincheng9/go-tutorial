package main

import "fmt"

func main() {
    // 构造一个map
    str := "aba"
    dict := map[rune]int{}
    for _, value := range str{
        dict[value]++
    }
    fmt.Println(dict) // map[97:2 98:1]
    
    // 访问map里不存在的key，并不会像C++一样自动往map里插入这个新key
    value, ok := dict['z']
    fmt.Println(value, ok) // 0 false
    fmt.Println(dict) // map[97:2 98:1]
    
    // 访问map里已有的key
    value2 := dict['a']
    fmt.Println(value2) // 2
}