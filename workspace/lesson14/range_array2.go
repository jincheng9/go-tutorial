package main

import "fmt"
import "reflect"

func main() {
    array := [2][3]int{{1, 2, 3}, {4, 5, 6}}
    // 只拿到行的索引
    for index := range array {
        // array[index]类型是一维数组
        fmt.Println(reflect.TypeOf(array[index])) 
        fmt.Printf("index=%d, value=%v\n", index, array[index])
    }
    
    // 拿到行索引和该行的数据
    for row_index, row_value := range array {
        fmt.Println(row_index, reflect.TypeOf(row_value), row_value)
    }
    
    // 双重遍历，拿到每个元素的值
    for row_index, row_value := range array {
        for col_index, col_value := range row_value {
            fmt.Printf("array[%d][%d]=%d ", row_index, col_index, col_value)
        }
        fmt.Println()
    }
}