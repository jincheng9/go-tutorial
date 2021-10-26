package main

import "fmt"
import "reflect"

func sum(array [5]int, size int) int{
    sum := 0
    for i:=0; i<size; i++ {
        sum += array[i]
    }
    return sum
}

func sumSlice(array []int, size int) int{
    sum := 0
    for i:=0; i<size; i++ {
        sum += array[i]
    }
    return sum
}

func main() {
    a := [5]int {1, 2, 3, 4, 5} // a := [...]int{1, 2, 3, 4, 5}也可以去调用sum，编译器会自动推导出a的长度5
    fmt.Println("type of a:", reflect.TypeOf(a)) // 输出type of a: []int，是一个slice类型
    ans := sum(a, 5)
    fmt.Println("ans=", ans)
    
    b := []int{1, 2, 3, 4, 5}
    ans2 := sumSlice(b, 5)
    fmt.Println("ans2=", ans2)
    
    array := [...]int {1}
    fmt.Println("type of array:", reflect.TypeOf(array)) // type of array: [1]int，是一个数组类型
    fmt.Println("array=", array)
}