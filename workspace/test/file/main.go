package main

import (
	"fmt"
	"os"
)

func main() {
	if file_info, err := os.Stat("main.go"); os.IsNotExist(err) {
		fmt.Println(file_info, err)
	} else {
		fmt.Println(err)
	}
}
