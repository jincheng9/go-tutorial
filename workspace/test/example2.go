// example2.go
package main

func main() {
	a := 10
	func() {
		a = 1
	}()
}