package main

type I interface{ M() }

type T struct {
	x int
}

func (t T) M() {
	println(t.x)
}

func f() {
	var t = &T{1}
	var i I = t
	defer i.M()
	t.x = 2
	return
}

func g() {
	var t = &T{1}
	var i I = t
	f := i.M
	defer f()
	t.x = 2
	return
}

func main() {
	f() // 1
	g() // 2
}
