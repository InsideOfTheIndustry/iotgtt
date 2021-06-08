package copyappend

import "testing"

func TestAppendSlice(t *testing.T) {
	AppendSlice()
}

func TestCopyslice(t *testing.T) {
	CopySlice()
}

func TestForiSlice(t *testing.T) {
	ForiSlice()
}

func BenchmarkAppendSlice(t *testing.B) {
	for n := 0; n < t.N; n++ {
		AppendSlice()
	}

}

func BenchmarkCopyslice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CopySlice()
	}
}
func BenchmarkForiSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ForiSlice()
	}
}
