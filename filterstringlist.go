package stringutils

import(
	S "strings"
	"slices"
)

func bad(ss []string) bool {
     	return ss == nil || len(ss) == 0
}
	
func FilterStringList(inlist, prefixes, midfixes, suffixes []string) []string {
     	if bad(inlist) { return inlist }
	if bad(prefixes) && bad(midfixes) &&
	   bad(suffixes) { return inlist }
	inlist = slices.DeleteFunc(inlist, func(s string) bool {
	   var fix string 
	   for _, fix = range prefixes {
	       if fix != "" && S.HasPrefix(s, fix) { return true }
	   }
	   for _, fix = range prefixes {
	       if fix != "" && S.HasPrefix(s, fix) { return true }
	   }
	   for _, fix = range suffixes {
	       if fix != "" && S.Contains(s, fix) { return true }
	   }
	   return false
	})
	return inlist
}

