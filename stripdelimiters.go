package stringutils

import (
	S "strings"
	"errors"
)

// StripDelimiters tries to remove corresponding characters 
// from both ends of the input string. From the resulting
// string, surrounding spaces get trimmed. If an error is
// returned, the original string is also returned. 
//
// About the meaning of "corresponding":
//  - If `delimiters` is a single character (like a single or double quote),
//    it tries to remove that character from both ends of the input string.
//  - If `delimiters` is two characters (like parentheses, braces, brackets),
//    they are assumed to form a left/right pair, and so the ends of the
//    input string are treated differently.
// .
func StripDelimiters(in string, delims string) (string, error) {

	in = S.TrimSpace(in)
	strerr := "|"+in+"|"+delims+"|"

	if in == "" || delims == "" { 
		return in, errors.New("stripdelimiters bad input " + strerr)
	}
	var hasBeg, hasEnd bool
	inLen := len(in)

	if len(delims) == 1 {
		hasBeg = S.HasPrefix(in, delims)
		hasEnd = S.HasSuffix(in, delims)
		if hasBeg && hasEnd {
			return S.TrimSpace(in[1 : inLen-1]), nil
		}
		return in, errors.New("stripdelimiters failed " + strerr)
	}
	if len(delims) != 2 {
		return in, errors.New(
		          "stripdelimiters bad delim spec |" + delims + "|")
	}
	sBeg := string(delims[0])
	sEnd := string(delims[1])
	hasBeg = S.HasPrefix(in, sBeg)
	hasEnd = S.HasSuffix(in, sEnd)
	if hasBeg && hasEnd {
		return S.TrimSpace(in[1 : inLen-1]), nil
	}
	return in, errors.New("stripdelimiters failed " + strerr)
}

// StripQuotes tries to strip off matching XML quotes (i.e. 
// either single or double quotes). From the resulting string,
// surrounding spaces get trimmed. If an error is returned,
// the original string is also returned. 
func StripQuotes(in string) (string, error) {
	in = S.TrimSpace(in)
	out1, e1 := StripDelimiters(in, "'")
	if e1 == nil { return out1, nil } 
	out2, e2 := StripDelimiters(in, "\"")
	if e2 == nil { return out2, nil }
	return in, errors.New("stripquotes failed |" + in + "|")
}
