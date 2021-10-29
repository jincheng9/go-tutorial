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
    // 传map实参给空接口
    dict := map[string]int{"a":1}
    print(dict) // type:map[string]int, value:map[a:1]
    
    // 传struct实参给空接口
    cat := Cat{"nimo", 2}
    print(cat) // type:main.Cat, value:{nimo 2}
}