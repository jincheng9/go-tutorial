package main

import "fmt"

type Person interface {
    getName() string
    getAge() int
}

type Man struct {
    name string
    age int
    height float32
}

func(man Man) getName() string{
    fmt.Println("name:", man.name)
    return man.name
}

func(man Man) getAge() int{
    fmt.Println("age:", man.age)
    return man.age
}

/*
func(man Man) getHeight() {
    fmt.Println("height:", man.height)
}*/

func main() {
    var boy Person = Man{"test2", 10, 1.80}
    
    boy.getName()
}