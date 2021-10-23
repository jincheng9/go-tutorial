package main

import "fmt"


/*
常量定义的时候必须赋值，定义后值不能被修改，否则编译报错
*/
const a int = 10

const b, c int = 20, 30

const d = "str"

const d1, d2 = "str1", "str2"

/*
常量可以定义枚举
*/
const (
	unknown = 0
	male = 1
	female = 2
	super
)

const (
	flower = 1
	wind = "ab"
	sun   // the value of sun is the same as wind, both are "ab"
)

const (
	f1 = iota // f1 = 0
	f2 // f2 = 1
	f3 // f3 = 2
)
const (
	v1 = iota  // the value of v1 is 0
	v2 = 1 << iota   // current iota is 1, the value of v2 is 1<<1 = 2
	v3   // current iota is 2, v3 = (1<<iota) = 1<<2 = 4
)

const h = iota // h = 0

const (
	class1 = 0
	class2 // class2 = 0
	class3 = iota  // class3 = 2
	class4 // class4 = 3
	class5 = "abc" 
	class6 // class6 = "abc"
	class7 = iota // class7 is 6
)

func main() {
	fmt.Println(a, b, c)
	println(d)
	println("d1=", d1, "d2=", d2)
	println(unknown, male, female, super)
	println(flower, wind, sun)
	println("f1=", f1, "f2=", f2, "f3=", f3)
	println(v1, v2, v3)
	println("h=", h)
	println("class value:", class1, class2, class3, class4, class5, class6, class7)

	const g1 = iota // g=0
	const g2 = iota 
	println("g1=", g1, "g2=", g2)

	const (
		h1 = 0
		h2 = iota // h2 is 1
		h3 // h3 is 2
	)
	println("h1=", h1, "h2=", h2, "h3=", h3)
}