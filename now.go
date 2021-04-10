package stringutils

import (
	S "strings"
	"time"
)

func Now() string {
	s := time.Now().UTC().Format(time.RFC3339)
	return S.Replace(s, "T", "_", 1)
}
