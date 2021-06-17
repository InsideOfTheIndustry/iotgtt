package memescape

import (
	"fmt"
	"testing"
)

func BenchmarkNewPersonWithPointer(b *testing.B) {
	var p *Person
	for i := 0; i < b.N; i++ {
		p = NewPersonWithPointer(1, "1", "1")
	}
	_ = fmt.Sprintf("%s", p.Name)
}

func BenchmarkNewPersonWithValue(b *testing.B) {
	var p Person
	for i := 0; i < b.N; i++ {
		NewPersonWithValue(1, "1", "1")
	}
	_ = fmt.Sprintf("%s", p.Name)
}
