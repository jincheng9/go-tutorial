package main

import "fmt"

func main() {
    var dict map[string]int = map[string]int{}
    dict["a"] = 1
    fmt.Println(dict)
    
    var dict2 = map[string]int{}
    dict2["b"] = 2
    fmt.Println(dict2)
    
    dict3 := map[string]int{"test":0}
    dict3["c"] = 3
    fmt.Println(dict2)
    
    dict4 := make(map[string]int)
    dict4["d"] = 4
    fmt.Println(dict4)
}