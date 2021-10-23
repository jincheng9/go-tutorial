package main


func add(a, b int, c, d string) (int, string) {
	return a+b, c+d
}

func main() {
	a, b := 1, 2
	c, d := "c", "d"
	res1, res2 := add(a, b, c, d)
	println("res1=", res1, "res2=", res2)
}