package stringutils

func DeleteEmptyStrings(in []string) (out []string) {
	for _, s := range in {
		if s == "" {
			continue
		}
		out = append(out, s)
	}
	return
}

func Enslice(in string) []string {
	out := make([]string, 0)
	out = append(out, in)
	return out
}
