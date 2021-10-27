package main

import "fmt"

func main() {
    slice := []int{1, 2, 3}
    // 往原切片里加一个元素
    test := append(slice, 4)
    // append不会改变slice的值，除非把append的结果重新赋值给slice
    fmt.Println(slice) // [1 2 3]
    fmt.Println(test) // [1 2 3 4]
    
    // append添加切片
    temp := []int{1,2}
    test = append(test, temp...) // 注意，第2个参数有...结尾
    fmt.Println(test) // [1 2 3 4 1 2]
    
    /*下面对array数组做append就会报错:  first argument to append must be slice; have [3]int
    array := [3]int{1, 2, 3}
    array2 := append(array, 1)
    fmt.Println(array2)
    */
}