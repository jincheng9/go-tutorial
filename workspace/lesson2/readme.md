# 数据类型
* 数字
    * 整数：int, uint8, uint16, uint32, uint64, int8, int16, int32, int64
    * 浮点数：float32, float64
    * 复数：complex64, complex128
    
* 字符串：string

    * len(str)函数可以获取字符串长度

        ```go
        package main
        
        import "fmt"
        
        func main() {
            str := "abcdgfg"
            fmt.Println(len(str)) // 7
        }
        ```

        

* bool

* 其它数字类型
    * byte：类似uint8，数据范围0-255，定义的时候超过这个范围会编译报错
    * rune：类似int32，数据范围-2147483648-2147483647
    * uint：32位或64位
    * uintptr: 无符号整数，是内存地址的整数表示形式，应用代码一般用不到（https://stackoverflow.com/questions/59042646/whats-the-difference-between-uint-and-uintptr-in-golang）
