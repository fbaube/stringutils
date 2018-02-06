package stringutils

import S "strings"

// var theSuffix = "_gxml"

// DirNameFromFileName is TBS. NO period on suffix !
func DirNameFromFileName(path string, suffix string) (dirName string, usable bool) {
	// Even bother ?
	if S.HasPrefix(path, ".") || S.HasSuffix(path, "~") || suffix == "" {
		return path, false
	}
	// Is it already suffixed ?
	if S.HasSuffix(path, suffix) {
		return path, false
	}
	// Usual case
	return path + suffix, true
}
