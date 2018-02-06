package stringutils

// Yn returns a single Y/n character.
// They are different case to improve readability!
func Yn(b bool) string {
	if b {
		return "Y"
	}
	return "n"
}
