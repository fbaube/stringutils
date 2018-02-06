package stringutils

import (
	"fmt"
	S "strings"
)

// StripDelimiters tries to remove matching characters from both
// ends of the input string; "must"==true makes it mandatory.
// From the resulting string, surrounding spaces get trimmed.
// It returns a slice of the input string, not a new string.
// If "must"==false, an error is returned iff the arguments were bad.
//
// If "delimiters" is a single character (like a single or double quote),
// it tries to remove that character from both ends of the input string.
// If "delimiters" is two characters (like parentheses, braces, brackets),
// the ends of the input string are treated differently.
func StripDelimiters(in string, delims string, must bool) (string, error) {

	in = S.TrimSpace(in)
	if must && (in == "" || delims == "") {
		return in, fmt.Errorf("Bad input to StripDelimiters")
	}
	var hasBeg, hasEnd bool
	inLen := len(in)

	if len(delims) == 1 {
		hasBeg = S.HasPrefix(in, delims)
		hasEnd = S.HasSuffix(in, delims)
		if hasBeg && hasEnd {
			return S.TrimSpace(in[1 : inLen-1]), nil
		}
		if must {
			return in, fmt.Errorf("StripDelimiters failed: |%s|%s|", in, delims)
		}
		return in, nil
	}
	if len(delims) != 2 {
		return in, fmt.Errorf("StripDelimiters bad delim spec: |%s|", delims)
	}
	sBeg := string(delims[0])
	sEnd := string(delims[1])
	hasBeg = S.HasPrefix(in, sBeg)
	hasEnd = S.HasSuffix(in, sEnd)
	if hasBeg && hasEnd {
		return S.TrimSpace(in[1 : inLen-1]), nil
	}
	if must {
		return in, fmt.Errorf("StripDelimiters failed: |%s|%s|", in, delims)
	}
	return in, nil
}
