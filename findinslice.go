package stringutils

import S "strings"

// IsInSlice returns (`i,true`) if the string is found in the slice
// (or an argument is bad). It returns (`-1,false`) if not found.)
func IsInSlice(s string, ss []string) (int, bool) {
	if s == "" || ss == nil || len(ss) == 0 {
		return -1, false
	}
	for i, tmp := range ss {
		if s == tmp {
			return i, true
		}
	}
	return -1, false
}

// IsInSliceIgnoreCase is like IsInSlice but without case matching.
func IsInSliceIgnoreCase(s string, ss []string) bool {
	if s == "" || ss == nil || len(ss) == 0 {
		return false
	}
	for _, tmp := range ss {
		if S.EqualFold(s, tmp) {
			return true
		}
	}
	return false
}
