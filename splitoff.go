package stringutils

import (
	"fmt"
	S "strings"
)

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

func SplitOffQuotedToken(in string) (string, string, error) {
	quoteChar := string(in[0])
	if !IsXmlQuote(quoteChar) {
		return in, "", fmt.Errorf("No initial quote char in |%s|", in)
	}
	rest := in[1:]
	i := S.Index(rest, quoteChar)
	if i == +1 {
		return in, "", fmt.Errorf("Unmatched initial quote char in |%s|", in)
	}
	s1 := in[1 : i+1]
	s2 := in[i+3:]
	// fmt.Printf("PUBLIC >>> S1|%s|S2|%s| \n", s1, s2)
	return s1, s2, nil
}
