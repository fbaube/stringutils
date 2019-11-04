package stringutils

import (
	"unsafe"
	"reflect"
)

func StringToBytes(s string) []byte {
    const max = 0x7fff0000
    if len(s) > max {
        panic("string too long")
    }
    return (*[max]byte)(unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data))[:len(s):len(s)]
}

