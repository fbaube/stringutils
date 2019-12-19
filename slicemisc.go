package stringutils

// DeleteEmptyStrings returns the slice with any empty strings omitted.
func DeleteEmptyStrings(in []string) (out []string) {
	for _, s := range in {
		if s == "" {
			continue
		}
		out = append(out, s)
	}
	return
}

// Enslice turns the string into a string slice of length 1.
func Enslice(in string) []string {
	out := make([]string, 0)
	out = append(out, in)
	return out
}

func TruncateTo(in string, outmaxlen int) string {
	if len(in) <= outmaxlen {
		return in
	}
	return in[:outmaxlen-3] + "..."
}
