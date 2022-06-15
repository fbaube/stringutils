package stringutils

import (
	"fmt"
	"strconv"
	S "strings"
	"time"
)

// Ito09az converts its int arg (0..35) to
// a string of length one, in the range
// (for int  0..9)  "0".."9",
// (for int 10..35) "a".."z"
func Ito09az(i int) string {
	if i <= 9 {
		return strconv.Itoa(i)
	}
	var bb = make([]byte, 1, 1)
	bb[0] = byte(i - 10 + 'a')
	return string(bb)
}

// NowAsYMDHM maps (reversibly) the current
// time to "YMDhm" (where m is minutes / 2).
func NowAsYMDHM() string {
	var now = time.Now()
	// year = last digit
	var y = fmt.Sprintf("%d", now.Year())[3:]
	var m = Ito09az(int(now.Month()))
	var d = Ito09az(now.Day())
	var h = Ito09az(now.Hour())
	var n = Ito09az(now.Minute() / 2)
	// fmt.Printf("%s-%s-%s-%s-%s", y, m, d, h, n)
	return fmt.Sprintf("%s%s%s%s%s", y, m, d, h, n)
}

// NowAsYM maps (reversibly) the current
// year+month to "YM".
func NowAsYM() string {
	var now = time.Now()
	// year = last digit
	var y = fmt.Sprintf("%d", now.Year())[3:]
	var m = Ito09az(int(now.Month()))
	// fmt.Printf("%s-%s-%s-%s-%s", y, m, d, h, n)
	return fmt.Sprintf("%s%s", y, m)
}

// PrettifyISO converts
// 2022-02-17T15:22:07+02:00 to
// 2022-02-17/15:22:07/+02
func PrettifyISO(in string) string {
	// func ReplaceAll(s, old, new string) string
	out := S.ReplaceAll(in, "T", "/")
	out = S.ReplaceAll(out, "+", "/+")
	out = S.ReplaceAll(out, ":00", "")
	return out
}
