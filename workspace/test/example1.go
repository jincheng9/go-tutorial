// https://github.com/golang/go/issues/9862
package main

import "fmt"


type any = interface{}

func getValue(m map[string]string, key string) any{
	value, exist := m[key]
	if !exist {
		var a any
		return a
	} else {
		return value
	}
}

func main() {
	m := map[string]string{"a":"1"}
	value := getValue(m, "a")
	fmt.Println(value)
	fmt.Printf("%T\n", value)

	var b interface{} = "a"
	b = b + "1"
	fmt.Println(b)
}