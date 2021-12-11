// example1.go
package main

import "fmt"

func changeMap(data map[string]interface{}) {
	data["c"] = 3
}

func main() {
	counter := map[string]interface{}{"a": 1, "b": 2}
	fmt.Println("begin:", counter)
	changeMap(counter)
	fmt.Println("after:", counter)
}
