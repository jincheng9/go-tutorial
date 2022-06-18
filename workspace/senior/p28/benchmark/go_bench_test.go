package main

import (
	"testing"
)

// func TestAdd(t *testing.T) {
// 	result := add(1, 2)
// 	if result == 3 {
// 		t.Log("test success")
// 	} else {
// 		t.Errorf("test failed, expected result: 3, actual result:%d", result)
// 	}
// }

func BenchmarkWrong(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		add(1000000000, 1000000001)
	}
}

// var result int

// func BenchmarkCorrect(b *testing.B) {
// 	var r int
// 	for i := 0; i < b.N; i++ {
// 		r = add(1000000000, 1000000001)
// 	}
// 	result = r
// }
