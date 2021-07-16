// Package stringutils does string things.
//
// Some files use Markdown, so `godoc2md` will work on 'em.
//
// This package does not use regular expressions anywhere. Also,
// this package accommodates XML string quoting conventions by
// allowing the use of either single quotes or double quotes.
//
// For functions that return a string or two plus a possibly-nil
// error, they can be wrapped in functions Must(..) or Must2(..),
// respectively, which either return successfully or a panic.
//
package stringutils
