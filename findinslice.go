package stringutils

import S "strings"

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
