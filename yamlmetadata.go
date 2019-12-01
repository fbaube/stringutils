package stringutils

import (
  "fmt"
  S "strings"
  "time"
	yaml "gopkg.in/yaml.v2"
)

// YamlMeta extracts metadata per LwDITA, see
// https://github.com/jelovirt/org.lwdita/wiki/Syntax-reference ,
// plus a few other fields expected to be useful.
type YamlMeta struct {
  // LwDITA...
	Author      string
  Source      string
  Publisher   string
  Permissions string
  Audience    string
  Category    string
  Keyword     string
  Resourceid  string
  // Bloggenator...
  // type Meta struct {
  Title      string
  Short      string
  Tags       []string
  Date       string
  ParsedDate time.Time
  // Others ...
	ID    string
}

// GetYamlMeta tries to extract a YAML metadata block (YMB) - as a
// `struct YamlMeta` - from the (start of the) input string `instring`.
// There are three cases:
//
// If a serious error is encountered in either of the following
// two other cases, the function returns `(nil, "", the-error)`.
//
// If the input string does NOT begin with a "---" delimiter
// line, then the entire input string is assumed to contain
// the YMB, and the function returns `(*YamlMeta, "", nil)`.
//
// If the input begins with an opening "---" delimiter line AND
// also has a matching closing "---" delimiter line - as in LwDITA
// - then the section between them is assumed to contain the YMB,
// and the YMB is trimmed from the input string, and the function
// returns `(*YamlMeta, instring-with-YMB-removed, nil)`.
//
// Note that there are basically three ways we could go about doing this:
// 1) Use simple decodingt oextract a few supported fields (including of course
// those used by LwDITA MDATA-XP).
// 2) Write something to extract all top-level (only) fields (maybe using regex's).
// 3) Use a third-party library to extract a proper tree structure.
//
// #3 is overkill at the nmment, and #1 is used.
//
func GetYamlMetadata(instr string) (map[string]interface{}, string, error) {
  hasDelim := S.HasPrefix(instr, "---")
  // The YAML markers are "---" (at start of line) for both START end END.
  // So, if the content does NOT start with the YAML-metadata START marker...
  if hasDelim && !S.Contains(instr, "\n") {
    return nil, "", fmt.Errorf("yaml: unterminated opening line")
  }
  var rawYMB, nonYMBretval string
  rawYMB = instr
  nonYMBretval = ""
  // If no delmiters, this entire large block is skipped.
  if hasDelim {
	   // The end of the line of the START marker
	   idx1 := S.Index(instr, "\n")
	   // The END marker
	   idx2 := S.Index(instr[idx1+2:], "\n---")
     // The end of the line of the END marker
     idx3 := S.Index(instr[idx2+2:], "\n")
	   // If no end marker, or no end of line of end-marker, reject.
	   if idx2 == -1 || idx3 == -1 {
		    return nil, "", fmt.Errorf("yaml: bad or missing end marker")
	   }
	   rawYMB = instr[idx1+1 : idx2+1]
     nonYMBretval = instr[idx3+1:]
   }
  // raw now contains what we want. Let's VERIFY.
	println("=== YAML? ===\n" +
          "vvv       vvv\n",
          rawYMB,
          "\n^^^       ^^^\n" +
          "=== YAML? ===")
	// YM := new(YamlMeta)
  // func Unmarshal(in []byte, out interface{}) (err error) <br/>
  // Unmarshal decodes the first document found within the in
  // byte slice and assigns decoded values into the out value.
  // Maps and ptrs (to a struct, string, int, etc) are accepted as out
  // values. If an internal ptr within a struct is not initialized, the
  // yaml package will initialize it if necessary for unmarshalling the
  // provided data. The out parameter must not be nil.
	// yaml.Unmarshal([]byte(rawYMB), YM)
	// fmt.Printf("YamlMeta: %+v \n", *YM)
  YMmap := make(map[string]interface{})
  // YMmap := make(map[interface{}]interface{})
  yaml.Unmarshal([]byte(rawYMB), YMmap)
	fmt.Printf("YamlMetaMap: %+v \n", YMmap)
	return YMmap, nonYMBretval, nil
}
