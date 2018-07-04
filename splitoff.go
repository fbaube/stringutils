package stringutils

import (
	"fmt"
	S "strings"
)

// SplitOffFirstWord splits the input string around the first whitespace,
// as defined by `strings.Fields(..)`, which uses `unicode.IsSpace`.
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
// (i.e. either single quote or double quote); it splits off the entire
// quoted string, returned without the quotes, and the rest of the
// input string is returned in the other return value.
func SplitOffQuotedToken(in string) (string, string, error) {
	quoteChar := string(in[0])
	if !IsXmlQuote(quoteChar) {
		return in, "", fmt.Errorf("No initial quote char in |%s|", in)
	}
	rest := in[1:]
	i := S.Index(rest, quoteChar)
	if i == -1 {
		return in, "", fmt.Errorf("Unmatched initial quote char in |%s|", in)
	}
	s1 := in[1 : i+1]
	var s2 string
	if len(rest) == i || len(rest) == i+1 {
		s2 = ""
	} else {
		s2 = S.TrimSpace(in[i+3:])
	}
	// fmt.Printf("PUBLIC >>> S1|%s|S2|%s| \n", s1, s2)
	return s1, s2, nil
}
