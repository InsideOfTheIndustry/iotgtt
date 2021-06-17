package byte2string

import (
	"reflect"
	"unsafe"
)

func ConverByte2Str(b []byte) string {
	// reflect to get the pointer of byte b
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	str := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}

	return *(*string)(unsafe.Pointer(&str))
}

func ConverStr2Byte(str string) []byte {
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))

	b := reflect.SliceHeader{
		Data: strHeader.Data,
		Len:  strHeader.Len,
		Cap:  strHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&b))
}
