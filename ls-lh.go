package stringutils

import(
	"fmt"
	"io/fs"
	S "strings"
	HB "github.com/fbaube/humanbytes"
)

var lsTimeFormat = "_2 Jan 2006 15:04"

/* type io/fs.FileInfo interface {
   Name() string // base name of the file
   Size() int64  // length in bytes for regular files; system-dep for others
   Mode() FileMode // file mode bits
   ModTime() time.Time  // mod time
   IsDir() bool  // abbreviation for Mode().IsDir()
   Sys()   any   // underlying data source (can return nil)
   } */

// LS_lh generates a file listing string like for "ls -l -h".
// If optPath is "" or "." then the file base name is taken 
// from the FileInfo argument.
// . 
func LS_lh(fi fs.FileInfo, optPath string) string {
        if fi == nil { return "" }
	// var optSlash string 
	if optPath == "" || optPath == "." {
	   optPath = fi.Name()
	   }
	if fi.IsDir() && !S.HasSuffix(optPath, "/") {
	   optPath += "/" // Slash = "/"
	   }
	return fmt.Sprintf("%s %4s  %s  %s%s",
		fi.Mode(), HB.SizeLS(int(fi.Size())), 
		fi.ModTime().UTC().Format(lsTimeFormat),
		optPath) // ,optSlash)
	}

