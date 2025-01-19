package stringutils

import (
	S "strings"
	"time"
)

func Now() string {
/*
	s := time.Now().UTC().Format(time.RFC3339)
	return S.Replace(s, "T", "_", 1)
	*/
	return NowPlus()
}

// NowPlus returns the current local date+time in a sensible format,
// i.e. "2006-01-02-mon/15:04:05/-07".
// 
// BTW "T" in ISO date-times is horrible for readability; use "_" instead.
func NowPlus() string {
	return S.ToLower(time.Now().Local().
	       Format("2006-01-02-Mon/15:04:05/-07"))
}

