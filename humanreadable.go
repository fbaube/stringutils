package stringutils

import(
	"math"
	"fmt"
)

/*
https://xdg.me/human-readable-bytes/
https://github.com/xdg/zzz-humanbytes
ls -lh has special, irregular semantics I couldn’t
find all implemented together in the same library!
 - Use no decimal pt or unit suffix for values under 1024 bytes: 0 or 897
 - Use one decimal if the integer part is a single digit: 1.0 K or 9.9 M
 - Use  no decimal if the integer part is more than one digit: 10 K or 582 M
 - Aggressively round up: 999 M + 1 byte rounds up to 1 G
*/

type format struct {
	base     float64
	logBase  float64
	suffixes []string
}

var formatLS = format{
	base:     1024,
	logBase:  math.Log(1024),
	suffixes: []string{"", " K", " M", " G", " T", " P", " E", " Z", " Y"},
}

func SizeLS(size int) string {
	return humanSize(float64(size), formatLS)
}

func humanSize(size float64, f format) string {
	if size == 0 {
		return "0"
	}

	mag := math.Floor(math.Log(size) / f.logBase)
	size /= math.Pow(f.base, mag)

	switch {
	case mag == 0:
		// do nothing
	case size < 10:
		size = math.Ceil(size*10) / 10
	default:
		size = math.Ceil(size)
	}

	if size >= f.base {
		size /= f.base
		mag++
	}

	format := "%.1f%s"
	if mag == 0 || size >= 10 {
		format = "%.0f%s"
	}

	return fmt.Sprintf(format, size, f.suffixes[int(mag)])
}

