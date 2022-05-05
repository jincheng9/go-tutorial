// quiz1.go
package main

import "fmt"

func main() {
	var a *int = new(int)
	*a = 5.0
	fmt.Println(*a)
}
