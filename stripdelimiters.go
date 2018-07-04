package stringutils

import (
	"fmt"
	S "strings"
)

// StripDelimiters tries to remove corresponding characters from
// both ends of the input string; `must`==`true` makes it mandatory.
// From the resulting string, surrounding spaces get trimmed.
//
// If `must`==`true`, it returns an error if the delimiters are not found.
// If `must`==`false`, an error is returned *iff* an argument was bad.
//
// About the meaning of "corresponding":
// If `delimiters` is a single character (like a single or double quote),
// it tries to remove that character from both ends of the input string.
// If `delimiters` is two characters (like parentheses, braces, brackets),
// they are assumed to form a left/right pair, and so the ends of the
// input string are treated differently.
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

// StripQuotes tries to strip off matching XML quotes (i.e. either single
// or double quotes). `success` indicates whether matching quotes were
// found, and `must` makes the function fail if they are not found.
func StripQuotes(in string, must bool) (out string, success bool) {
	in = S.TrimSpace(in)
	out, _ = StripDelimiters(in, "'", false)
	out, _ = StripDelimiters(out, "\"", false)
	if len(in)-2 == len(out) {
		return out, true
	}
	if must {
		return out, false
	}
	return out, true
}
