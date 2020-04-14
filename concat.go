package stringutils

import "bytes"

// ConcatAll concatenates all its (vargs) arguments into a new string.
// It uses a bytes.Buffer for efficiency.
func ConcatAll(values ...string) string {
	var bb bytes.Buffer
	for _, s := range values {
		bb.WriteString(s)
	}
	return bb.String()
}

// ConcatAllSpaced concatenates all its (vargs) arguments
// into a new string, with spaces inserted in-between 'em.
// It uses a bytes.Buffer for efficiency. 
func ConcatAllSpaced(values ...string) string {
	var bb bytes.Buffer
	for _, s := range values {
		bb.WriteString(s)
		bb.WriteString(" ")
	}
	s := bb.String()
	return s[:len(s)-1] // sweeeet
}
