package main

import "fmt"


// interface1
type Felines interface {
    feet() 
}

// interface2, 嵌套了interface1
type Mammal interface {
    Felines
    born()
}

// 猫实现Mammal这个interface里的所有方法
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
    /*Mammal有feet和born方法，2个都可以调用*/
    var a Mammal = &cat
    a.feet()
    a.born()
    
    var b Felines = cat
    b.feet()
    // b.born() 调用这个会编译报错，因为Felines没有born方法
}