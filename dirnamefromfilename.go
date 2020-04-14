package stringutils

import S "strings"

// var theSuffix = "_gxml"

// DirNameFromFileName takes a file name and appends a specified
// suffix-to-add (like, say, "_xml") in order to create a directory that
// will hold files somehow related to the input file. NO period on suffix!
func DirNameFromFileName(path string, suffixToAdd string) (dirName string, usable bool) {
	// Even bother ?
	if S.HasPrefix(path, ".") || S.HasSuffix(path, "~") || suffixToAdd == "" {
		return path, false
	}
	// Is it already suffixed ?
	if S.HasSuffix(path, suffixToAdd) {
		return path, false
	}
	// Usual case
	return path + suffixToAdd, true
}
