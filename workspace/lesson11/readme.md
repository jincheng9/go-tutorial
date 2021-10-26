# 指针

* 声明和初始化

  * 语法

    ```go
    var var_name *vartype
    ```

  * 示例

    ```go
    var intPtr *int
    ```

  * 初始化

    ```go
    package main
    
    import "fmt"
    import "reflect"
    
    func main() {
        i := 10
        // 方式1
        var intPtr *int = &i
        fmt.Println("pointer value:", intPtr, " point to: ", *intPtr)
        fmt.Println("type of pointer:", reflect.TypeOf(intPtr))
        
        // 方式2
        intPtr2 := &i
        fmt.Println(*intPtr2)
        fmt.Println("type of pointer:", reflect.TypeOf(intPtr2))
        
        // 方式3
        var intPtr3 = &i;
        fmt.Println(*intPtr3)
        fmt.Println("type of pointer:", reflect.TypeOf(intPtr3))
        
        // 方式4
        var intPtr4 *int
        intPtr4 = &i
        fmt.Println(*intPtr4)
        fmt.Println("type of pointer:", reflect.TypeOf(intPtr4))
    }
    ```

    

* 默认值

  * 不赋值的时候，默认值是nil

    ```go
    var intPtr5 *int    
    fmt.Println("intPtr5==nil:", intPtr5==nil) // intPtr5==nil: true
    ```

* 指针数组

  * 定义

    ```go
    var ptr [SIZE]*int // 指向int的指针数组，数组里有多个指针，每个都指向一个int
    ```

  * 使用

* 指向指针的指针

* 向函数传递指针参数

