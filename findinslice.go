package stringutils

import S "strings"

func InSlice(s string, ss []string) (int, bool) {
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

func InSliceIgnoreCase(s string, ss []string) (int, bool) {
	if s == "" || ss == nil || len(ss) == 0 {
		return -1, false
	}
	for i, tmp := range ss {
		if S.EqualFold(s, tmp) {
			return i, true
		}
	}
	return -1, false
}
