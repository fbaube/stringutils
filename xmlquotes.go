package stringutils

import (
	"fmt"
	S "strings"
)

// IsXmlQuote checks whether the string is either 
// one single quote or one double quote.
func IsXmlQuote(s string) bool {
	return s == "\"" || s == "'"
}

// IsXmlQuoted checks whether the string is surrounded 
// by either single quotes or double quotes.
func IsXmlQuoted(txt string) bool {
	L := len(txt)
	if L < 2 {
		return false
	}
	char0 := string(txt[0])
	if char0 != "'" && char0 != "\"" {
		return false
	}
	return (char0 == string(txt[L-1]))
}

// MustXmlUnquote removes either paired single quotes or
// paired double quotes. It panics if neither is found.
func MustXmlUnquote(txt string) string {
	if !IsXmlQuoted(txt) {
		panic(fmt.Sprintf("stringutils.MustXmlUnquote<%s>", txt))
	}
	return txt[1 : len(txt)-1]
}

// NormalizeWhitespace replaces weird whitespace junk
// (including newlines) with spaces.
func NormalizeWhitespace(s string) string {
	return S.Join(S.Fields(s), " ")
}
