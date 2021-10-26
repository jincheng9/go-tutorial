package main

import "fmt"

func main() {
    // 定义的是直接初始化赋值
    array1 := [2][3]int {
        {0, 1, 2},
        {3, 4, 5}, // 如果花括号}在下一行，这里必须有逗号。如果花括号在这一行可以不用逗号
    }
   
    fmt.Println("array1=", array1)
    
    // 先给默认值，再对其中某些元素赋值
    array2 := [2][3]int{}
    array2[0][2] = 1
    array2[1][1] = 2
    fmt.Println("array2=", array2)
    
    // 通过append赋值
    twoDimArray := [][]int{}
    row1 := []int{1,2,3}
    row2 := []int{4,5}
    twoDimArray = append(twoDimArray, row1)
    fmt.Println("twoDimArray=", twoDimArray)
    
    twoDimArray = append(twoDimArray, row2)
    fmt.Println("twoDimArray=", twoDimArray)
    
    fmt.Println(twoDimArray[0][2])
    fmt.Println(twoDimArray[1][1])
    
    // 遍历二维数组，按照下标
    for i:=0; i<2; i++ {
        for j:=0; j<3; j++ {
            fmt.Printf("array1[%d][%d]=%d ", i, j, array1[i][j])
        }
        fmt.Println()
    }
    
    // 遍历二维数组，按照range方式迭代
    for index := range twoDimArray {
        fmt.Printf("row %d is ", index) // index的值是0,1，表示二维数组的第1行和第2行
        fmt.Println(twoDimArray[index])
    }
    for row_index, row_value := range twoDimArray {
        for col_index, col_value := range row_value {
            fmt.Printf("twoDimArray[%d][%d]=%d ", row_index, col_index, col_value)
        }
        fmt.Println()
    }
    
    // 可以直接取二维数组的某一行完整数据出来
    oneDimArray := array1[0]
    fmt.Println(oneDimArray) // [0, 1, 2]
    
}