package copyappend

import (
	"testing"
)

func TestAppendSlice(t *testing.T) {
	if actualValue := AppendSlice(NewTestSlice()); actualValue != expectedValue {
		t.Errorf("got %v unexpected", actualValue)
	}
}

func TestCopyslice(t *testing.T) {
	if actualValue := CopySlice(NewTestSlice()); actualValue != expectedValue {
		t.Errorf("got %v unexpected", actualValue)
	}
}

func TestForiSlice(t *testing.T) {
	if actualValue := ForiSlice(NewTestSlice()); actualValue != expectedValue {
		t.Errorf("got %v unexpected", actualValue)
	}
}

func BenchmarkAppendSlice(t *testing.B) {
	for n := 0; n < t.N; n++ {
		AppendSlice(NewTestSlice())
	}

}

func BenchmarkCopyslice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CopySlice(NewTestSlice())
	}
}

func BenchmarkForiSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ForiSlice(NewTestSlice())
	}
}

var expectedValue = "2345671234567123456712345671234567123456712345671234567123456712345671234567123456712345671234567123456712345671234567123456712345671234567123456712345671234567123456712345671234567123456712345671234567123456712345671234567"
