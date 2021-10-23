package main

const num int = 10

func main() {
	var sum = 0
	for i:=1; i<=num; i++ {
		sum +=i
	}

	var sum2 = 0
	var j int = 0
	for j<=10 {
		sum2 += j
		j++
	}
	println("sum2=", sum2)

	for {
		println("test infinite loop")
		break
	}

	for true {
		println("test infinite loop2")
		break
	}
}