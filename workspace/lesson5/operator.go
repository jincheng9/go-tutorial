package main

func main() {
	var a int = 10
	var b *int = &a
	println("a=",a, "address=", &a)
	println("b=", b, "*b=", *b)
}