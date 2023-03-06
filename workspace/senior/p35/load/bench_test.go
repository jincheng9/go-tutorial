package main

import (
	"testing"
)

func BenchmarkLoad(b *testing.B) {
	if err := generateLoad(b.N); err != nil {
		b.Errorf("generateLoad got err %v want nil", err)
	}
}
