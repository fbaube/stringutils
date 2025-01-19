package stringutils

import (
	"unsafe"
	"reflect"
)

// StringToBytes uses both reflect and unsafe, so it's a bit of a turd. 
func StringToBytes(s string) []byte {
    const max = 0x7fff0000
    if len(s) > max {
        panic("string too long")
    }
    return (*[max]byte)(unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data))[:len(s):len(s)]
}

