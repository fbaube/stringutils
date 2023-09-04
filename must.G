package stringutils

// Must panics if a func returning one string returns an error.
func Must(s string, e error) string {
	if e != nil {
		println("==> SU.Must(s) failed, returning:", s)
		panic(e)
	}
	return s
}

// Must2 panics if a func returning two strings returns an error too.
func Must2(s1, s2 string, e error) (string, string) {
	if e != nil {
		println("==> SU.Must2(s1,s2) failed, returning(1):", s1)
		println("==> SU.Must2(s1,s2) failed, returning(2):", s2)
		panic(e)
	}
	return s1, s2
}
