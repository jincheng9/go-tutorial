package fuzz

import (
	"testing"
)

func TestReverse(t *testing.T) {
	str_map := map[string]string{"abc": "cba", "b": "b", "吃": "吃"}
	for k, v := range str_map {
		result := Reverse(k)
		if v != result {
			t.Errorf("unit test failed, k: %s, result:%s", k, result)
		}
	}
}
