package fuzz

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	str_slice := []string{"abc"}
	for _, v := range str_slice {
		f.Add(v)
	}
	f.Fuzz(func(t *testing.T, str string) {
		rev_str1 := Reverse(str)
		rev_str2 := Reverse(rev_str1)
		if str != rev_str2 {
			t.Errorf("fuzz test failed. str:%s, rev_str1:%s, rev_str2:%s", str, rev_str1, rev_str2)
		}
		if utf8.ValidString(str) && !utf8.ValidString(rev_str1) {
			t.FailNow()
			//t.Errorf("reverse result is not utf8. str:%s, len: %d, rev_str1:%s", str, len(str), rev_str1)
		}
	})
}
