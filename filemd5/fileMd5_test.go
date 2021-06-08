package filemd5

import (
	"testing"
)

func TestByReadall(t *testing.T) {
	actualValue := ByReadall()
	if actualValue != "e6b19a79bed47a8951bc6b478d1ceeed" {
		t.Errorf("Got %v, expected %v", actualValue, "e6b19a79bed47a8951bc6b478d1ceeed")
	}
}

func TestByPieceRead(t *testing.T) {
	actualValue := ByPieceRead()
	if actualValue != "e6b19a79bed47a8951bc6b478d1ceeed" {
		t.Errorf("Got %v, expected %v", actualValue, "e6b19a79bed47a8951bc6b478d1ceeed")
	}
}

func BenchmarkByReadall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ByReadall()
	}
}

func BenchmarkByPieceRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ByPieceRead()
	}
}
