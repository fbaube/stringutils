package stringutils

// Stringser is a handy interface for content. Echo reproduces
// the content, Info is a normal level of debuggery, and Debug
// should be verbose. 
type Stringser interface {
	Echo() string
	Info() string
	Debug() string
}
