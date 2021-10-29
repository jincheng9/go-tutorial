package main

import "fmt"


// interface1，猫科动物的共同行为
type Felines interface {
    feet() 
}

// interface2, 哺乳动物的共同行为
type Mammal interface {
    born()
}

// 猫既是猫科动物也是哺乳动物，2个行为都实现
type Cat struct {
    name string
    age int
}

func(cat Cat) feet() {
    fmt.Println("cat feet")
}

func(cat *Cat) born() {
    fmt.Println("cat born")
}

func main() {
    cat := Cat{"rich", 1}
    var a Felines = cat
    a.feet()
    
    var b Mammal = &cat
    b.born()
}