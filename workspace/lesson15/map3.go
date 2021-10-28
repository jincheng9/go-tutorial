package main

import "fmt"

func main() {
    dict :=  map[string]int{"a":1, "b":2}
    fmt.Println(dict)
    
    // 删除"a"这个key
    delete(dict, "a")
    fmt.Println(dict)
    
    // 删除"c"这个不在的key，对map结果无影响
    delete(dict, "c")
    fmt.Println(dict)
}