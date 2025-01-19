package stringutils

import(
	"strings"
	"slices"
)

// DeleteEmptyStrings returns the slice with any
// empty or all-whitespace strings deleted.
// It also compacts the slice, so element indices change. 
func DeleteEmptyStrings(in []string) (out []string) {
/*
	for _, s := range in {
		if s == "" {
			continue
		}
		out = append(out, s)
	}
*/
	// func DeleteFunc[S ~[]E, E any](s S, del func(E) bool) S
	// It removes any elements from s for which del returns true,
	// returning the modified slice. It zeroes the elements btwn
	// the new length and the original length.
	return slices.DeleteFunc(in, func (s string) bool {
	       return s == "" || strings.TrimSpace(s) == ""
	})
}

// Enslice turns the string into a string slice of length 1.
func Enslice(in string) []string {
	out := make([]string, 0)
	out = append(out, in)
	return out
}

// TruncateTo truncates & ends the string with triple dots if it's too long.
func TruncateTo(in string, outmaxlen int) string {
	if len(in) <= outmaxlen {
		return in
	}
	return in[:outmaxlen-3] + "..."
}
