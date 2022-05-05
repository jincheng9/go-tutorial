// quiz2.go
package main

import "fmt"

func main() {
	var a *int = new(int)
	var b float32 = 5.0
	*a = b
	fmt.Println(*a)
}
