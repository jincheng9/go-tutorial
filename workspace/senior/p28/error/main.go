package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	cause := errors.New("whoops")
	err1 := errors.Wrapf(cause, "oh noes")
	err2 := errors.Wrapf(err1, "hi")
	fmt.Println(cause)
	fmt.Println(err1)
	fmt.Println(err2)
}
