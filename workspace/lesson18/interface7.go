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
    /*
    v, ok := x.(string)
    if ok {
        fmt.Println("assert true, value:", v)
    } else{
        fmt.Println("assert false")
    }*/
    
    
    x = Cat{"hugo", 3}
    checkType(x)
}