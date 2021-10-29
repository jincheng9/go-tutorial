package main

import "fmt"

// all animals can speak
type Animal interface {
    speak()
}

// cat
type Cat struct {
    name string
    age int
}

func(cat Cat) speak() {
    fmt.Println("cat miaomiaomiao")
}

// dog
type Dog struct {
    name string
    age int
}

func(dog *Dog) speak() {
    fmt.Println("dog wangwangwang")
}


func main() {
    var animal Animal = Cat{"gaffe", 1}
    animal.speak() // cat miaomiaomiao
    
    /*
    因为Dog的speak方法用的是指针接受者，因此给interface赋值的时候，要赋指针
    */
    animal = &Dog{"caiquan", 2}
    animal.speak() // dog wangwangwang
}