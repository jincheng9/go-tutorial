package strings

import "fmt"

func init() {
	fmt.Println("reverse init")
}

func Reverse(str string) string {
	r := []rune(str)
	reverseData := reverseRune(r)
	return string(reverseData)
}
