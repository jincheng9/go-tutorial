package main

func test1(ch chan int) {

}

func test2(ch chan int) {

}

func main() {
	ch1 := make(chan<- int)
	test1(ch1)

	ch2 := make(<-chan int)
	test2(ch2)
}
