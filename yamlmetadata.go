package stringutils

import (
  "fmt"
  S "strings"
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
  // Others ...
	ID    string
  Title string
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
func GetYamlMetadata(instr string) (*YamlMeta, string, error) {
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
	println("=== v YAML? v ===\n", rawYMB, "\n=== ^ YAML? ^ ===")
	YM := new(YamlMeta)
	yaml.Unmarshal([]byte(rawYMB), YM)
	fmt.Printf("YamlMeta: %+v \n", *YM)
	return YM, nonYMBretval, nil
}

func (p *YamlMeta) AsMap() map[string]string {
  m := make(map[string]string)
  // Grunt work
  if p.Author != "" {
    m["author"] = p.Author
  }
  if p.Source != "" {
    m["source"] = p.Source
  }
  if p.Publisher != "" {
    m["publisher"] = p.Publisher
  }
  if p.Permissions != "" {
    m["permissions"] = p.Permissions
  }
  if p.Audience != "" {
    m["audience"] = p.Audience
  }
  if p.Category != "" {
    m["category"] = p.Category
  }
  if p.Keyword != "" {
    m["keyword"] = p.Keyword
  }
  if p.Resourceid != "" {
    m["resourceid"] = p.Resourceid
  }
  if p.ID != "" {
    m["id"] = p.ID
  }
  if p.Title != "" {
    m["title"] = p.Title
  }
  return m
}
