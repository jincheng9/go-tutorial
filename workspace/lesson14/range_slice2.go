package main

import "fmt"
import "reflect"

func main() {
    slice := [][]int{{1,2}, {3, 4, 5}}
    fmt.Println(len(slice))
    // 方法1，拿到行索引
    for index := range slice{
        fmt.Printf("index=%d, type:%v, value=%v\n", index, reflect.TypeOf(slice[index]), slice[index])
    }
    
    // 方法2，拿到行索引和该行的值，每行都是一维切片
    for row_index, row_value := range slice{
        fmt.Printf("index=%d, type:%v, value=%v\n", row_index, reflect.TypeOf(row_value), row_value)
    }
    
    // 方法3，双重遍历，获取每个元素的值
    for row_index, row_value := range slice {
        for col_index, col_value := range row_value {
            fmt.Printf("slice[%d][%d]=%d ", row_index, col_index, col_value)
        }
        fmt.Println()
    }
}