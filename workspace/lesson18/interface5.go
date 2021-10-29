package main

import "fmt"


type Cat struct {
    name string
    age int
}

// 打印空interface的类型和具体的值
func print(x interface{}) {
    fmt.Printf("type:%T, value:%v\n", x, x)
}

func main() {
    // 定义空接口x
    var x interface{}
    // 将map变量赋值给空接口x
    x = map[string]int{"a":1}
    print(x) // type:map[string]int, value:map[a:1]
    
    // 传struct变量估值给空接口x
    cat := Cat{"nimo", 2}
    x = cat
    print(x) // type:main.Cat, value:{nimo 2}
}