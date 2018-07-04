package stringutils

// Yn returns (as a string) a single character, `Y` or `n`.
//
// They are different case to improve readability, *duh!*
//
// In a table, it can be even more readable to use "-" for false.
// People who fill tables with `Y` and `N` defeat simple visual
// scanning and are idiots.
func Yn(b bool) string {
	if b {
		return "Y"
	}
	return "n"
}
