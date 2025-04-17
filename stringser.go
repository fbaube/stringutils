package stringutils

// Stringser is a handy interface for content:
//  - Echo reproduces the content (not guaranteed 100% tho) 
//  - Infos is a normal, summary level of description
//  - Debug should be verbose
//
// Notes: 
//  - Policy re newlines is TBS
//  - Debug can default to something like [fmt.Printf] ("%v")
//  - [Infos] used to be Info, but that conflicted with
//    [fs.DirEntry.Info] for struct [fileutils.FSItemer]
// .
type Stringser interface {
	Echo()  string
	Infos() string
	Debug() string
}
