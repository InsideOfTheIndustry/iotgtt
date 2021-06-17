package byte2string

import (
	"testing"
)

func TestConverByte2Str(t *testing.T) {
	testCases := []struct {
		desc          []byte
		expectedValue string
	}{
		{
			desc:          []byte(""),
			expectedValue: "",
		},
		{
			desc:          []byte("string convert to []byte"),
			expectedValue: "string convert to []byte",
		},
		{
			desc:          []byte("111222"),
			expectedValue: "111222",
		},
	}
	for _, tC := range testCases {
		if actualValue := ConverByte2Str(tC.desc); tC.expectedValue != actualValue {
			t.Errorf("Got %v expected %v.", actualValue, tC.expectedValue)
		}
	}
}

func TestConverStr2Byte(t *testing.T) {
	testCases := []struct {
		expectedValue []byte
		desc          string
	}{
		{
			expectedValue: []byte(""),
			desc:          "",
		},
		{
			expectedValue: []byte("string convert to []byte"),
			desc:          "string convert to []byte",
		},
		{
			expectedValue: []byte("111222"),
			desc:          "111222",
		},
	}

	for _, tC := range testCases {
		actualValue := ConverStr2Byte(tC.desc)
		for i := 0; i < len(actualValue); i++ {
			if actualValue[i] != tC.expectedValue[i] {
				t.Errorf("Got %v expected %v.", actualValue, tC.expectedValue)
				goto failed
			}
		}
	}
failed:
}

func BenchmarkOriginByte2Str(b *testing.B) {
	bt := []byte("Convert me to string!")
	for i := 0; i < b.N; i++ {
		_ = string(bt)
	}
}

func BenchmarkConverByte2Str(b *testing.B) {
	bt := []byte("Convert me to string!")
	for i := 0; i < b.N; i++ {
		_ = ConverByte2Str(bt)
	}
}

func BenchmarkOriginStr2Byte(b *testing.B) {
	str := "Convert me to string!"
	for i := 0; i < b.N; i++ {
		_ = []byte(str)
	}
}
func BenchmarkConverStr2Byte(b *testing.B) {
	str := "Convert me to string!"
	for i := 0; i < b.N; i++ {
		_ = ConverStr2Byte(str)
	}
}
