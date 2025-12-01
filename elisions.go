package stringutils

import (
	"os"
	FP "path/filepath"
	S "strings"
	// WU "github.com/fbaube/wasmutils"
)

// PathSep is a token nod to Windoze compatibility.
const PathSep = string(os.PathSeparator)

// These should end in the path separator!
// NOTE See init(), at bottom.
var currentWorkingDir, userHomeDir string

// GetHomeDir is a convenience function, and
// refers to the invoking user's home directory.
func GetHomeDir() string {
	return userHomeDir
}

// ELIDE HOME DIR
// ELIDE CWD

// ElideHomeDir doesn't try to deal with
// a relative path that starts with "."
//
func ElideHomeDir(s string) string {
	if s == "" {
		return ""
	}
	// If it can't be normalised...
	if len(s) <= 2 {
		s = S.TrimSuffix(s, PathSep)
		if s == "." || s == "~" || s == PathSep || s == "" {
			return s
		}
	}
	// If it can't be further normalised...
	if S.HasPrefix(s, "~"+PathSep) {
		return s
	}
	// At this point, if it's not an absolute FP, it's
	// a relative FP, but let it slide and don't check
	// for "." prefix and don't prepend "./".
	if !FP.IsAbs(s) {
		// panic("NiceFP barfs on: " + s)
		return s
	}
	if s == userHomeDir || (s+PathSep) == userHomeDir {
		return "~"
	}
	if S.HasPrefix(s, userHomeDir) {
		return ("~" + PathSep + S.TrimPrefix(s, userHomeDir))
		// bytesToTrim := len(currentUserHomeDir) + 1
		// return "~" + PathSep + s[bytesToTrim:]
	}
	// No luck
	return s

}

// ElideCWD has effects that are session-specific, so 
// it should only be used to make better user messages.
func ElideCWD(s string) string {
	if s == "" || s == currentWorkingDir ||
		(s+PathSep) == currentWorkingDir {
		return "."
	}
	// If it can't be normalised...
	if len(s) <= 2 {
		s = S.TrimSuffix(s, PathSep)
		if s == "." || s == "~" || s == PathSep || s == "" {
			return s
		}
	}
	// If it can't be further normalised...
	if S.HasPrefix(s, "."+PathSep) {
		return s
	}
	// At this point, if it's not an absolute FP, it's a relative FP,
	// but let it slide and don't prepend "./".
	if !FP.IsAbs(s) {
		// panic("NiceFP barfs on: " + s)
		return s
	}
	// println("arg:", s)
	// println("cwd:", currentworkingdir)

	if S.HasPrefix(s, currentWorkingDir) {
		return ("." + PathSep + S.TrimPrefix(s, currentWorkingDir))
		// bytesToTrim := len(currentWorkingDir) + 1
		// return "." + PathSep + s[bytesToTrim:]
	}
	// No luck
	return s
}

// Tildotted shortens a filepath by ediding the current
// user's home directory, or the current working directory,
// by eliding "~" or "." respectively. 
func Tildotted(s string) string {
	// If it's missing, use assumed/default...
	if s == "" {
		return "."
	}
	// If it's CWD...
	if s == currentWorkingDir || (s+PathSep) == currentWorkingDir {
		return "."
	}
	// If it can't be normalised...
	if len(s) <= 2 {
		s = S.TrimSuffix(s, PathSep)
		if s == "." || s == "~" || s == PathSep || s == "" {
			return s
		}
	}
	// If it can't be further normalised...
	if S.HasPrefix(s, "."+PathSep) || S.HasPrefix(s, "~"+PathSep) {
		return s
	}
	// At this point, if it's not an absolute FP, it's a relative FP,
	// but let it slide and don't prepend "./".
	if !FP.IsAbs(s) {
		// panic("NiceFP barfs on: " + s)
		return s
	}
	// println("arg:", s)
	// println("cwd:", currentworkingdir)

	if S.HasPrefix(s, currentWorkingDir) {
		return ("." + PathSep + S.TrimPrefix(s, currentWorkingDir))
		// bytesToTrim := len(currentWorkingDir) + 1
		// return "." + PathSep + s[bytesToTrim:]
	}
	if S.HasPrefix(s, userHomeDir) {
		return ("~" + PathSep + S.TrimPrefix(s, userHomeDir))
		// bytesToTrim := len(currentUserHomeDir) + 1
		// return "~" + PathSep + s[bytesToTrim:]
	}
	// No luck
	return s
}

func init() {
	var e error
	/*
	if WU.IsWasm() {
		userHomeDir = "?"
		currentWorkingDir = "."
		return
	}
	*/
	userHomeDir, e = os.UserHomeDir()
	if e != nil {
		println("Cannot determine current user's home directory (Wasm?)")
		return
	}
	currentWorkingDir, e = os.Getwd()
	if e != nil {
		println("Cannot determine current working directory (Wasm?)")
		return
	}
	if !S.HasSuffix(currentWorkingDir, PathSep) {
		currentWorkingDir = currentWorkingDir + PathSep
	}
	if !S.HasSuffix(userHomeDir, PathSep) {
		userHomeDir = userHomeDir + PathSep
	}
}
