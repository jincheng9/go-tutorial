package main

import (
	"fmt"
	"sync"
)

func test1() {
	var a, b, c []int
	var wg sync.WaitGroup

	a = append(a, 0, 0, 0)
	fmt.Println(len(a), cap(a))
	wg.Add(2)

	go func() {
		b = append(a, 1)
		wg.Done()
	}()

	go func() {
		c = append(a, 2)
		wg.Done()
	}()
	wg.Wait()
}

func test2() {
	s := make([]int, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		s[1] = 1
		wg.Done()
	}()

	go func() {
		s[1] = 1
		wg.Done()
	}()

	wg.Wait()
}

func test3() {
	m := map[int]int{}
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		m[1] = 1
		wg.Done()
	}()

	var mm map[int]int
	go func() {
		mm = m
		// fmt.Println(mm)
		wg.Done()
	}()

	wg.Wait()
}

func test4() {
	var a = []int{0, 0}
	var b, c []int
	var wg sync.WaitGroup

	a = append(a, 0)

	wg.Add(2)

	go func() {
		b = append(a, 1)
		wg.Done()
	}()

	go func() {
		c = append(a, 2)
		wg.Done()
	}()

	wg.Wait()
}

func test5() {
	func() {
		fmt.Println(1)
	}()
}

//#region

func main() {
	test5()
}

//#endregion
