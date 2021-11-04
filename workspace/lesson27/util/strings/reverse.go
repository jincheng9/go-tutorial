package strings

func Reverse(str string) string {
	r := []rune(str)
	reverseData := reverseRune(r)
	return string(reverseData)
}
