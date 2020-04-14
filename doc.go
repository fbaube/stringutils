// Package stringutils does string things.
//
// Some files in this directory use Markdown, so `godoc2md` will work on 'em.
//
// This package does not use regular expressions anywhere, and more
// importantly, accommodates how XML permits string quoting using
// either single quotes or double quotes.
//
// For functions that return a string or two plus a possibly-nil error, the
// functions Must(..) and Must2(..) can return either success or a panic.
//
package stringutils
