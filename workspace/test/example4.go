package main

import "fmt"

var battle = make(chan string)

func warrior(name string, done chan struct{}) {
	select {
	case opponent := <-battle:
		fmt.Printf("%s beat %s\n", name, opponent)
	case battle <- name:
		// I lost :-(
	}
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})
	langs := []string{"Go", "C", "C++", "Java", "Perl", "Python"}
	for _, l := range langs { go warrior(l, done) }
	for _ = range langs { fmt.Println(0);<-done }
}