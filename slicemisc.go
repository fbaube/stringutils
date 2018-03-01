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
