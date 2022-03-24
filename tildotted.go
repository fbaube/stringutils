package stringutils

import (
	"os"
	FP "path/filepath"
	S "strings"

	WU "github.com/fbaube/wasmutils"
)

// A token nod to Windoze compatibility.
const PathSep = string(os.PathSeparator)

// These should end in the path separator!
// NOTE See init(), at bottom.
var currentWorkingDir, currentUserHomeDir string

// GetHomeDir is a convenience function, and
// refers to the invoking user's home directory.
func GetHomeDir() string {
	return currentUserHomeDir
}

// Tildotted shortens a filepath by substituting "~" or ".".
func Tildotted(s string) string {
	// If it's missing, use assumed/default...
	if s == "" {
		return "."
	}
	// If it's CWD...
	if s == currentWorkingDir {
		return "."
	}
	// If it can't be normalised...
	if s == "." || s == "~" || s == PathSep {
		return s
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
	if S.HasPrefix(s, currentUserHomeDir) {
		return ("~" + PathSep + S.TrimPrefix(s, currentUserHomeDir))
		// bytesToTrim := len(currentUserHomeDir) + 1
		// return "~" + PathSep + s[bytesToTrim:]
	}
	// No luck
	return s
}

func init() {
	var e error
	if WU.IsWasm() {
		currentUserHomeDir = "?"
		currentWorkingDir = "."
		return
	}
	currentUserHomeDir, e = os.UserHomeDir()
	if e != nil {
		println("==> ERROR: Cannot determine current user's home directory")
		return
	}
	currentWorkingDir, e = os.Getwd()
	if e != nil {
		println("==> ERROR: Cannot determine current working directory")
		return
	}
	if !S.HasSuffix(currentWorkingDir, PathSep) {
		currentWorkingDir = currentWorkingDir + PathSep
	}
	if !S.HasSuffix(currentUserHomeDir, PathSep) {
		currentUserHomeDir = currentUserHomeDir + PathSep
	}
}
