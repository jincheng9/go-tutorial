package main

func main() {
	str := "abcdefg"
	for index, char := range str {
		println("index=", index, "char=", char)
	}

	numbers := [6]int{1, 2, 3, 5} // size为6的数组，第5个和第6个元素的值是默认值0
	for index, value := range numbers {
		println("index=", index, "value=", value)
	}

	var list []int = []int{1,2}
	for index, value := range list {
		println("index=", index, "value=", value)
	}
	strings := []string{"google", "nb"} // 2个元素的字符串数组
	for index, value := range strings {
		println("index=", index, "value=", value)
	}

	dict := map[string] int{"a":1, "b":2}
	for key, value := range dict {
		println("key=", key, "value=", value)
	}
}