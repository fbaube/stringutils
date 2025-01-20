package stringutils

import (
	"fmt"
)

// TrimMatchingDelims tried to strip off matching first and last characters.
// It is basically a simplified version of `StripDelimiters(..)`.
func TrimMatchingDelims(txt string, delim string) (string, error) {
	L := len(txt)
	if L < 2 || len(delim) != 1 {
		return txt, nil
	}
	char0 := string(txt[0])
	charN := string(txt[L-1])
	if char0 != delim || charN != delim {
		return txt, fmt.Errorf("Delimiters |%s| not found around |%s|", delim, txt)
	}
	return txt[1 : L-1], nil
}
