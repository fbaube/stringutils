package stringutils

import (
	"fmt"
	S "strings"
)

// IsXmlQuote checks for both single quote and double quote,
// cos XML uses them equivalently.
func IsXmlQuote(s string) bool {
	return s == "\"" || s == "'"
}

// IsXmlQuoted handles both single quotes and double quotes,
// cos XML uses them equivalently.
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

// MustXmlUnquote removed either single quotes or double quotes,
// cos XML uses them equivalently.
func MustXmlUnquote(txt string) string {
	if !IsXmlQuoted(txt) {
		panic(fmt.Sprintf("stringutils.MustXmlUnquote<%s>", txt))
	}
	return txt[1 : len(txt)-1]
}

// NormalizeWhitespace replaces weird
// whitespace crap (incl. newlines) with spaces.
func NormalizeWhitespace(s string) string {
	return S.Join(S.Fields(s), " ")
}
