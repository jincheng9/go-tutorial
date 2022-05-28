package main

func main() {
	n := 0

	f := func() func(int) {
		n = 1
		return func(int) {}
	}
	g := func() int {
		println(n)
		return 0
	}

	f()(g())
}
