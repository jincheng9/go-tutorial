// example4.go
package main

import "fmt"

// 测试数据
var indexRuneTests = []struct {
	s    string
	rune rune
	out  int
}{
	{"a A x", 'A', 2},
	{"some_text=some_value", '=', 9},
	{"☺a", 'a', 3},
	{"a☻☺b", '☺', 4},
}

func main() {
	fmt.Println(indexRuneTests)
}
