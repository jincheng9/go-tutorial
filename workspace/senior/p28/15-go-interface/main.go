package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

func copySourceToDest(src io.Reader, dest io.Writer) error {
	p := make([]byte, 10)
	byte_read_len, err := src.Read(p)
	if err != nil {
		return errors.New("read error")
	}
	fmt.Println(byte_read_len, err, p)
	for i := 0; i < 10; i++ {
		fmt.Println(p[i])
	}
	byte_write_len, err := dest.Write(p)
	if err != nil {
		return errors.New("write error")
	}
	fmt.Println(byte_write_len, err, p)
	return nil
}

func main() {
	const input = "foo"
	source := strings.NewReader(input)
	dest := bytes.NewBuffer(make([]byte, 0))

	err := copySourceToDest(source, dest)
	if err != nil {
		fmt.Println("error")
	}

	got := dest.String()
	if got != input {
		fmt.Printf("expected: %s, got: %s\n", input, got)
	}
}
