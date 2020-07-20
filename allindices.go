package stringutils

import (
	S "strings"
)

func AllIndices(in string, ch string) []int {
	if len(ch) != 1 {
		panic("Bad arg (len is not 1) to AllIndicesOf: " + ch)
	}
	if in == "" {
		return nil
	}
	// fmt.Printf("AllIndices: in<%s> ch<%s> \n", in, ch)
	var length, idxToStart int
	var idcs []int
	for true {
		length = S.Index(in[idxToStart:], ch)
		// fmt.Printf("idx: %v \n", length)
		if length == -1 {
			return idcs
		}
		idcs = append(idcs, idxToStart+length+1)
		idxToStart += length + 1
		// in = in[idxCume:]
	}
	return idcs
}
