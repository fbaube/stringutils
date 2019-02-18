package stringutils

// func TryOpenRO(path AbsFilePath) (*os.File, error) {

func Must(s string, e error) string {
	if e != nil {
		// if s is non-nil, we could write a big ERR to Stdout, and continue
		panic(e)
	}
	return s
}
