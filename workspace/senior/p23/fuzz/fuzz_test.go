package fuzz

import (
	"testing"
)

func FuzzReverse(f *testing.F) {
	str_slice := []string{"abc", "吃"}
	for _, v := range str_slice {
		f.Add(v)
	}
	f.Fuzz(func(t *testing.T, str string) {
		rev_str1 := Reverse(str)
		rev_str2 := Reverse(rev_str1)
		if str != rev_str2 {
			t.Errorf("fuzz test failed. str:%s, rev_str1:%s, rev_str2:%s", str, rev_str1, rev_str2)
		}
	})

}
