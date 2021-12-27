package main

func main() {
	c := make(chan int, 1)
	for done := false; !done; {
		select {
		default:
			print(1)
			done = true
		case <-c:
			print(2)
			c = nil
		case c <- 1:
			print(3)
		}
	}
}
