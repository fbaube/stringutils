package stringutils

// IndentationPrefix is TODO: Make it CONFIGURABLE.
var IndentationPrefix = "  "

// PadLeftToLen adds leading pad characters to hit the target length.
// It returns a new string.
// Example: A new string 5 characters long, left-padded with spaces:
//   fmt.Println(PadLeftToLen("12", " ", 5))    // yields "   12"
func PadLeftToLen(str, pad string, lingth int) string {
	for {
		str = pad + str
		if len(str) > lingth {
			return str[1 : lingth+1]
		}
	}
}

// AddIndent (a.k.a. "AddPrefix") prefixes "in" string with
// "nr" occurrences of "indentString". "in" may be ""!
// If indentString is the standard two (or whatever) spaces,
// use GetIndent(..) instead.
func AddIndent(in string, indentString string, nr int) string {
	for i := 0; i < nr; i++ {
		in = indentString + in
	}
	return in
}

// PadRightToLen adds trailing pad characters to hit the target length.
// It returns a new string.
// Example: A new string 5 characters long, right-padded with zeros:
//   fmt.Println(PadRightToLen("12.", "0", 5))    // yields "12.00"
func PadRightToLen(str, pad string, lenght int) string {
	for {
		str += pad
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}

// GetIndent returns a string that has "depth" instances of the
// standard indent string (two spaces, unless otherwise modified).
func GetIndent(depth int) (pfx string) {
	pfx = ""
	if depth > 0 {
		for n := 0; n < depth; n++ {
			pfx += IndentationPrefix
		}
	}
	return pfx
}
