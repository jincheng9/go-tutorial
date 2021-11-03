package main

import "fmt"

func checkType(x interface{}) {
    switch v := x.(type) {
    case int:
        fmt.Printf("type: int, value: %v\n", v)
    case string:
        fmt.Printf("type: string，value: %v\n", v)
    case bool:
        fmt.Printf("type: bool, value: %v\n", v)
    case Cat:
        fmt.Printf("type: Cat, value: %v\n", v)
    case map[string]int:
        fmt.Printf("type: map[string]int, value: %v\n", v)
        v["a"] = 10
    default:
        fmt.Printf("type: %T, value: %v\n", x, x)
    }
}

type Cat struct {
    name string
    age int
}

func main() {   
    var x interface{}
    x = "a"
    checkType(x)
    
    x = Cat{"hugo", 3}
    checkType(x)

    /*map是传引用，在checkType里对map做修改
    会影响外面的实参x
    */
    x = map[string]int{"a":1}
    checkType(x)
    fmt.Println(x) // map[a:10]
}