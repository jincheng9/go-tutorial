// quiz3.go
package main

import "fmt"

func main() {
	var a *int = new(int)
	*a = 5.1
	fmt.Println(*a)
}
