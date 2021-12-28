package main

func f() bool {
	return false
}

func main() {
	switch f(); 
	{
	case true:
		println(1)
	case false:
		println(0)
	default:
		println(-1)
	}
	a := 1
	fmt.println(a++)
}
