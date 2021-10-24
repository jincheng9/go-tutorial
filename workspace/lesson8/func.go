package main


func add(a, b int, c, d string) (int, string) {
	return a+b, c+d
}

func swap(a int, b int) {
	println("[func|swap]a=", a, "b=", b)
	a, b = b, a
	println("[func|swap]a=", a, "b=", b)
}

func swapRef(pa *int, pb *int) {
	println("[func|swapRef]a=", *pa, "b=", *pb)
	var temp = *pa
	*pa = *pb
	*pb = temp
	println("[func|swapRef]a=", *pa, "b=", *pb)
}

func main() {
	a, b := 1, 2
	c, d := "c", "d"
	res1, res2 := add(a, b, c, d)
	println("res1=", res1, "res2=", res2)

	println("[func|main]a=", a, "b=", b)
	swap(a, b)
	println("[func|main]a=", a, "b=", b)

	println("[func|main]a=", a, "b=", b)
	swapRef(&a, &b)
	println("[func|main]a=", a, "b=", b)	
}