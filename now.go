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

func NowPlus() string {
	return S.ToLower(time.Now().Local().
	       Format("2006-01-02-Mon/15:04:05/-07"))
}

