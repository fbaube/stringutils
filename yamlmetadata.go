package stringutils

import (
	"errors"
	"fmt"
	"strconv"
	S "strings"

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
// The metadata is unmarshalled into a map (i.e. a `PropSet`), so variables
// can be freely added, but there is no checking for required fields.
//
func GetYamlMetadataAsPropSet(instr string) (PropSet, error) {
	propmap, e := ParseYamlMetadata(instr)
	if e != nil {
		return nil, fmt.Errorf("yaml propset: %w", e)
	}
	var ps PropSet
	ps = make(PropSet, 0)
	// Convert all the values to strings
	for kk, v := range propmap {
		k := S.ToLower(kk)
		// println("D=> yaml got:", k)
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
	return ps, nil
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

// ParseYamlMetadata tries to extract a YAML metadata block (YMB) - as a map -
// from the (start of the) input string `instring`.
//
// Only simple fields are supported - no tree structure.
//
func ParseYamlMetadata(instr string) (map[string]interface{}, error) {
	// func Unmarshal(in []byte, out interface{}) (err error) <br/>
	// Unmarshal decodes the first document found within the byte
	// slice and assigns decoded values into the out value. Maps and
	// ptrs (to a struct, string, int, etc) are accepted as out values.
	YMmap := make(map[string]interface{})
	e := yaml.Unmarshal([]byte(instr), YMmap)
	if e != nil {
		return nil, fmt.Errorf("yaml parse: %w", e)
	}
	fmt.Printf("ParseYamlMetadata: unmarshal'd YamlMetaMap: %+v \n", YMmap)
	return YMmap, nil
}

// YamlMetadataHeaderLength assumes "---" AT THE START to open the block,
// and "---" (but not "...") at the start of a new line to end the block.
func YamlMetadataHeaderLength(s string) (int, error) {
	// println("D=> TRY YAML \n", s, "D=> END YAML")
	if !S.HasPrefix(s, "---") {
		return 0, nil
	}
	if !S.Contains(s, "\n") {
		return 0, errors.New("yaml: unterminated opening line")
	}
	// The end of the line of the START marker
	idxBegEOL := S.Index(s, "\n")
	// fmt.Printf("idxBegEOL %d \n", idxBegEOL)
	// The END marker (can be the very next line!)
	idxEnd := S.Index(s[idxBegEOL:], "\n---")
	// If no end marker, reject.
	if idxEnd == -1 {
		return 0, errors.New("yaml: bad or missing end marker")
	}
	idxEnd += idxBegEOL
	// fmt.Printf("idxEnd %d \n", idxEnd)

	// The end of the line of the END marker
	idxEndEOL := S.Index(s[idxEnd+1:], "\n")
	// If no end of line of end-marker, reject.
	if idxEndEOL == -1 {
		return 0, errors.New("yaml: unterminated end marker")
	}
	idxEndEOL += idxEnd + 1
	// fmt.Printf("idxEndEOL %d \n", idxEndEOL)

	// We now have the index of end of the block.
	// But if the next line is empty, include it also.
	if s[idxEndEOL+1] == '\n' {
		idxEndEOL++
	}
	// Let's VERIFY.
	// println("D=> BEG YAML \n", s[:idxEndEOL], "D=> END YAML")

	return idxEndEOL, nil
}

func TrimYamlMetadataDelimiters(s string) string {
	s = S.TrimSpace(s)
	s = S.TrimPrefix(s, "---")
	s = S.TrimSuffix(s, "---")
	return S.TrimSpace(s) + "\n"
}
