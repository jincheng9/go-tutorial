package main

import "fmt"

/*
global variable
*/
var a int = 10
var b int  // b will be defaulted to zero

var c = 30

// below has syntax error: unexpected newline, expecting type
// var d 

// this way can only be used within a function
// e:=10


// another way to define global variable
var (
    v1 int = 100
    v2 bool = true
)

var (
    v3 int // default to 0
    v4 bool  // default to false
)


func main() {
    fmt.Println(a, b, c)
    fmt.Println(v1, v2)
    fmt.Println("v3:", v3, "v4:", v4)
    /*
    local variable
    */
    var f int = 40
    var g = 50
    h := 60
    fmt.Println(f, g, h)
}