package stringutils

// MarkupType specifies the top-level, most general
// type of content in a content entity: XML, HTML,
// Markdown, or BIN (binary, "none of the above").
//
// There is also "Dirlike", for consistent handling
// of this field in a tree environment. 
// 
// Note that altho HTML is "obviously" XML, HTML is
// separated out here because HTML5 doesn't have a
// really good definition (like DTD or XSD), and also
// because HTML has a dedicated parser in the Go std
// lib that is quite a bit more liberal than the XML
// parser.
//
// MarkupType implements interface [stringutils.Stringser].
//
// NOTE that for source code that uses this enum, files
// can be named "*_xml.go", "*_html.go", "*_mkdn.go".
// .
type MarkupType string

const (
	MU_type_UNK MarkupType = "UNK" // UNKNOWN
	// MU_type_XML is assumed to be well-formed 
	MU_type_XML  = "XML"
	// MU_type_HTML is assumed to be HTML5
	MU_type_HTML = "HTML"
	// MU_type_MKDN is assumed to be CommonMark (or GFM?)
	MU_type_MKDN = "MKDN"
	// MU_type_BIN is opaque 
	MU_type_BIN  = "BIN"
	// MU_type_SQL is, well, hey why not eh 
	MU_type_SQL  = "SQL"
	// MU_type_DIRLIKE is a placeholder for consistent handling
	// (because of consistent problems in code): IsDirlike is
	// a more general case of IsDir() - shorthand for "is not 
	// possible to have own content" - but this can be more 
	// formally defined as "ia/has link(s) to other stuff" 
	// - i.e. it is a directory or symbolic link. Used by
	// [ctoken.TypedRaw].
	MU_type_DIRLIKE = "DIRLIKE" 
)

func (mt MarkupType) Echo() string {
	return string(mt)
}

func (mt MarkupType) Info() string {
	if len(mt) < 3 || len(mt) > 7 {
		panic("Bad MU_type: " + mt)
	}
	switch mt {
	case MU_type_XML:
		return "XML"
	case MU_type_HTML:
		return "HTML[5]"
	case MU_type_MKDN:
		return "Markdown"
	case MU_type_BIN:
		return "Binary"
	case MU_type_SQL:
		return "SQL"
	case MU_type_DIRLIKE:
		return "Dirlike"
	case MU_type_UNK:
		return "Unknown"
	}
	return string(mt)
}

func (mt MarkupType) Debug() string {
	return mt.Info()
}
