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

func changeArray(array [3]int) {
    array[0] = 10
}

func changeArray2(array *[3]int) {
    array[0] = 10
}

func changeArray3(array []int) {
    array[0] = 10
}

func main() {
    a := [5]int {1, 2, 3, 4, 5} // a := [...]int{1, 2, 3, 4, 5}也可以去调用sum，编译器会自动推导出a的长度5
    fmt.Println("type of a:", reflect.TypeOf(a)) // 输出type of a: []int，是一个slice类型
    ans := sum(a, 5)
    fmt.Println("a length=", len(a), " ans=", ans)
    
    b := []int{1, 2, 3, 4, 5}
    ans2 := sumSlice(b, 5)
    fmt.Println("b length=", len(b), " ans2=", ans2)
    
    c := [2][3]int{{1,2,3}, {3,4,5}}
    fmt.Println("c length=", len(c))
    
    param := [3]int{1,2,3}
    changeArray(param)
    fmt.Println("param=", param) // param= [1 2 3]
    changeArray2(&param)
    fmt.Println("param=", param) // param= [10 2 3]
    
    sliceArray := []int{1,2,3}
    changeArray3(sliceArray)
    fmt.Println("sliceArray=", sliceArray) // sliceArray= [10 2 3]
}