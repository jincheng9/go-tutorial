package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}
	fmt.Println(m)
}