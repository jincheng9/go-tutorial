package main

import "fmt"

type person interface {
    getName() string
    getAge() int
    getGender() string
}

type Man struct {
    name string
    age int
    gender string
    height float32
}

func(man Man) getName() {
    fmt.Println("name:", man.name)
}
func(man Man) getHeight() {
    fmt.Println("height:", man.height)
}

func main() {
    person := new(Man)
    person.getName()
    person.getHeight()
}