package main

import "fmt"

func main() {
    str := "abcdgfg"
    // 方法1：可以通过range拿到字符串的下标索引
    for index := range(str) {
        fmt.Printf("index:%d, value:%d\n", index, str[index])
    }
    fmt.Println()
    
    // 方法2：可以通过range拿到字符串的下标索引和对应的值
    for index, value := range(str) {
        fmt.Println("index=", index, ", value=", value)
    }
    fmt.Println()
    
    // 也可以直接通过len获取字符串长度进行遍历
    for index:=0; index<len(str); index++ {
        fmt.Printf("index:%d, value:%d\n", index, str[index])
    }
}