package main

import "fmt"

func main() {
    hash := map[string]int{"a":1}
    // 方法1，拿到key，再根据key获取value
    for key := range hash{
        fmt.Printf("key=%s, value=%d\n", key, hash[key])
    }
    
    // 方法2，同时拿到key和value
    for key, value := range hash{
        fmt.Printf("key=%s, value=%d\n", key, value)
    }
    
    /* nil map不能存放key-value键值对，比如下面的方式会报错：panic: assignment to entry in nil map
    var hash2 map[string]int 
    hash2["a"] = 1
    */
}