package main

import "fmt"


func main() {
    // 定义一个map类型的变量，key是string类型，value是空接口类型
    dict := make(map[string]interface{})
    // value可以是int类型
    dict["a"] = 1 
    // value可以是字符串类型
    dict["b"] = "b"
    // value可以是bool类型
    dict["c"] = true
    fmt.Println(dict) // map[a:1 b:b c:true]
    fmt.Printf("type:%T, value:%v\n", dict["b"], dict["b"]) // type:string, value:b
}