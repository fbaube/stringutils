package stringutils

import (
	"errors"
	S "strings"
)

// SplitOffFirstWord splits the input string around the first whitespace,
// as defined by func [strings.Fields], which uses [unicode.IsSpace].
func SplitOffFirstWord(in string) (string, string) {
	var w = S.Fields(in)
	switch len(w) {
	case 0:
		return "", ""
	case 1:
		return w[0], ""
	default:
		return w[0], S.TrimSpace(in[len(w[0])+1:])
	}
}

// SplitOffQuotedToken expects the string to start with an XML quote
// (i.e. either a single quote or a double quote, but not any wrongly-
// named "smart quotes"); it splits off the entire quoted string,
// returned without the quotes, and the rest of the input string
// is returned in the other return value.
func SplitOffQuotedToken(in string) (string, string, error) {
     	in = S.TrimSpace(in)
	if len(in) < 2 {
	   	return in, "", errors.New("Bad quoting in |"+in+"|")
	}
	quoteChar := string(in[0])
	if !IsXmlQuote(quoteChar) {
		return in, "", errors.New("No initial quote char in |"+in+"|")
	}
	bfr, afr, isMatched := S.Cut(in[1:], quoteChar)
	if !isMatched {
		return in, "", errors.New(
			"Unmatched initial quote char in |"+in+"|")
	}
	// fmt.Printf("PUBLIC >>> S1|%s|S2|%s| \n", s1, s2)
	return bfr, S.TrimSpace(afr), nil
}

