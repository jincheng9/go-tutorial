package main

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"path"
	"reflect"
	"time"
)

const c1 = iota
const c2 = iota

func scope() func() int{
	outer_var := 2
	foo := func() int { return outer_var}
	return foo
}

// Closures
func outer() (func() int, int) {
	outer_var := 2
	inner := func() int {
		outer_var += 99 // outer_var from outer scope is mutated.
		return outer_var
	}
	fmt.Println(outer_var)
	inner()
	return inner, outer_var // return inner func and mutated outer_var 101
}

func sumTotal(args ...int) int {
	total := 0
	fmt.Println(reflect.TypeOf(args))
	for _, num := range args {
		total += num
	}
	return total
}

func sliceTest() {
	var s []int
	s = append(s, []int{1,2}...)
	fmt.Println(sumTotal(s...))
}

func str() {
	//s := "abc"
	//c := s[0]
	var a rune = 'a'
	fmt.Println(reflect.TypeOf(a), a)
}

func control() {
here:
	for i := 0; i < 3; i++ {
		fmt.Println("i=", i)
		for j := i + 1; j < 3; j++ {
			if i == 0 {
				continue here
			}
			fmt.Println(j)
			if j == 2 {
				break
			}
		}
	}

there:
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 3; j++ {
			if j == 1 {
				continue
			}
			fmt.Println(j)
			if j == 2 {
				break there
			}
		}
	}
}

func sliceTest2() {
	a := []int{1,2}
	b :=a[:]
	fmt.Println(b)

	c := make([]int, 5)
	fmt.Println(cap(c))

	x := [3]string{"a", "b", "c"}
	s := x[:]
	s[0] = "z"
	fmt.Println(x, s)

	for range s {
		fmt.Println(s)
	}
	for range time.Tick(time.Second) {
		fmt.Println("tick")
	}
}


func sqrt(x float64) (float64, error){
	if x<0 {
		return 0, errors.New("negative value")
	}
	return math.Sqrt(x), nil
}

func printTest() {
	fmt.Println("Hello, 你好, नमस्ते, Привет, ᎣᏏᏲ") //基本的打印，会自动换行
	p := struct { X, Y int }{ 17, 2 }
	fmt.Println( "My point:", p, "x coord=", p.X ) // 打印结构体和字段值
	s := fmt.Sprintln( "My point:", p, "x coord=", p.X ) // print to string variable
	fmt.Println(s)
}

// 定义http响应的类型
type Hello struct{}

// 结构体Hello实现接口类型http.Handler里的方法ServeHTTP
// 这样结构体Hello的实例就可以作为http的Handler来接收http请求，返回http响应结果
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func strTest() {
	s := "a"
	fmt.Println(len(s))
	s2 := s[:1]
	//s2 = "b"
	s = "b"
	fmt.Println(s, s2)
}

func lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}

func joinPaths(absolutePath, relativePath string) string {
	fmt.Println(absolutePath, relativePath)
	if relativePath == "" {
		return absolutePath
	}
	finalPath := path.Join(absolutePath, relativePath)
	fmt.Println(finalPath)
	if lastChar(relativePath) == '/' && lastChar(finalPath) != '/' {
		return finalPath + "/"
	}
	fmt.Println(finalPath)
	return finalPath
}

func structTest() {
	type A struct {
		a int
		b int
	}

	type B struct {
		b float32
		c string
		d string
	}

	type C struct {
		A
		B
		a string
		c string
	}

	var c C
	fmt.Println(c.a, c.A.a, c.A.b, c.B.b)
}

type i interface{
	open(int) string
	close(int) string
}

type S struct{
	a,b int
	pending func (int, int) int
}
func(s S) open(a int) string {
	s.a = 10
	return string(a)
}

func(s *S) close(a int) string {
	s.a = 10
	return string(a)
}

func interfaceTest() {
	a := S{a:1, b:2, pending: nil}
	a.open(1)
	fmt.Println(a)
	a.close(1)
	fmt.Println(a)
}

type handleFunc func(int)int
type st struct{
	handleFunc
}

func structTest2() {
	a := st{func(i2 int) int {
		return i2
	}}
	result := a.handleFunc(1)
	fmt.Println(result)
}

func main() {
	//var h Hello
	//http.ListenAndServe("localhost:4000", h)
	structTest2()
	var a *int
	fmt.Println(*a)
}
