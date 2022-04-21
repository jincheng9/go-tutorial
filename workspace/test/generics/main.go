package main

import (
	"fmt"
	"sort"
)

func sortMethod1() {
	f := []float64{1.0, -2.0, -1.0, 0}
	sort.Float64s(f)
	fmt.Println(f)
}

func sortMethod2() {
	f := sort.Float64Slice{1.0, -2.0, -1.0, 0}
	sort.Sort(f)
	fmt.Printf("%T, %v\n", f, f)
}

func main() {
	sortMethod2()
}
