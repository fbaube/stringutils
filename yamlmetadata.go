package stringutils

import (
  "fmt"
  S "strings"
  "strconv"
	yaml "gopkg.in/yaml.v2"
)

/*
// OBSOLETE: YamlMeta extracts metadata per LwDITA, see
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
*/

// GetYamlMetadataAsPropSet is a convenience function. It assume that all
// the metadata values are top-level and can be represented as strings.
func GetYamlMetadataAsPropSet(instr string) (PropSet, string, error) {
  m, s, e := GetYamlMetadata(instr)
  // fmt.Printf("Yaml-map 4ps: %+v (err?:%t) \n", m, e != nil)
  if e != nil {
    return nil, instr, e
  }
  var ps PropSet
  ps = make(PropSet, 0)
  // Convert the values
  for kk,v := range m {
    k := S.ToLower(kk)
    // println("GOT:", k)
    switch vtyped := v.(type) {
    case string:
      ps[k] = vtyped
    case bool:
      ps[k] = strconv.FormatBool(vtyped)
    case int:
      i64 := int64(vtyped)
      ps[k] = strconv.FormatInt(i64, 10)
    case float32:
      f64 := float64(vtyped)
      ps[k] = strconv.FormatFloat(f64, 'f', -1, 32)
    case float64:
      ps[k] = strconv.FormatFloat(vtyped, 'f', -1, 64)
    default:
      panic(fmt.Sprintf("GetYamlMetadataAsPropSet: type<%T> (date?) \n", v))
    }
  }
  return ps, s, nil
}

// https://pandoc.org/MANUAL.html#extension-yaml_metadata_block
// A YAML metadata block is delimited by a line of three hyphens (---) at the
// top and a line of three hyphens (---) or three dots (...) at the bottom.
// A YAML metadata block not at the beginning must be preceded by a blank line.
// All string scalars will be interpreted as Markdown. Fields with names ending
// in an underscore will be ignored by pandoc. (They may be given a role by
// external processors.) Field names must not be interpretable as YAML numbers
// or boolean values (e.g. yes, True, and 15 cannot be used as field names).
// A document may contain multiple metadata blocks. If two metadata blocks
// attempt to set the same field, the value from the second block will be taken.

// GetYamlMetadata tries to extract a YAML metadata block (YMB) - as a map -
// from the (start of the) input string `instring`. There are three cases:
//
// If a serious error is encountered in either of the following two
// other cases, the function returns `(nil, instring, the-error)`.
//
// If the input string does NOT begin with a "---" delimiter
// line, then the entire input string is assumed to contain
// the YMB, and the function returns `(aMap, "", nil)`.
//
// If the input begins with an opening "---" delimiter line AND
// also has a matching closing "---" delimiter line - as in LwDITA
// - then the section between them is assumed to contain the YMB,
// and the YMB is trimmed from the input string, and the function
// returns `(aMap, instring-with-YMB-removed, nil)`.
//
// Note that there are basically three ways we could go about doing this:
// 1) Use simple decoding to extract a few supported fields (including of
// course those used by LwDITA MDATA-XP).
// 2) Write something to extract top-level (only) fields (maybe using regex's).
// 3) Use a third-party library to extract a proper tree structure.
//
// #3 is overkill at the nmment, and #1 is used.
//
func GetYamlMetadata(instr string) (map[string]interface{}, string, error) {
  hasDelim := S.HasPrefix(instr, "---")
  // The YAML markers are "---" (at start of line) for both START end END.
  // So, if the content does NOT start with the YAML-metadata START marker...
  if hasDelim && !S.Contains(instr, "\n") {
    return nil, instr, fmt.Errorf("yaml: unterminated opening line")
  }
  var rawYMB, nonYMBretval string
  rawYMB = instr
  nonYMBretval = ""
  // If no delmiters, this entire large block is skipped.
  if hasDelim {
	   // The end of the line of the START marker
	   idx1 := S.Index(instr, "\n")
	   // The END marker
	   idx2tmp := S.Index(instr[idx1+2:], "\n---")
     idx2 := idx2tmp + (idx1+2)
     // The end of the line of the END marker
     idx3tmp := S.Index(instr[idx2+2:], "\n")
     idx3 := idx3tmp + (idx2+2)
	   // If no end marker, or no end of line of end-marker, reject.
	   if idx2tmp == -1 || idx3tmp == -1 {
		    return nil, "", fmt.Errorf("yaml: bad or missing end marker")
	   }
	   rawYMB = instr[idx1+1 : idx2+1]
     nonYMBretval = instr[idx3+1:]
   }
  // raw now contains what we want. Let's VERIFY.
	println("==v YAML? v==\n",
          rawYMB,
          "==^ YAML? ==")
  // func Unmarshal(in []byte, out interface{}) (err error) <br/>
  // Unmarshal decodes the first document found within the byte
  // slice and assigns decoded values into the out value. Maps and
  // ptrs (to a struct, string, int, etc) are accepted as out values.
  YMmap := make(map[string]interface{})
  yaml.Unmarshal([]byte(rawYMB), YMmap)
	// fmt.Printf("GetYamlMetadata: unmarshal'd YamlMetaMap: %+v \n", YMmap)
	return YMmap, nonYMBretval, nil
}
