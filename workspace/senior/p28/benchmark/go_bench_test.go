package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := add(1, 2)
	if result == 3 {
		t.Log("test success")
	} else {
		t.Errorf("test failed, expected result: 3, actual result:%d", result)
	}
}

// var result int

func BenchmarkAdd(b *testing.B) {
	// var r int
	for i := 0; i < b.N; i++ {
		add(1, 2)
	}
	// result = r
}
