package main

import (
    "fmt"
    "sync"
)

func main() {
    m := sync.Map{}
    m.Delete("a")
    fmt.Println(m)
    
}