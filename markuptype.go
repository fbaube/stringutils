package stringutils

// MarkupType specifies the top-level, most general
// type of content in a content entity: XML, HTML,
// Markdown (or BIN: binary, "none of the above").
// Note that XML is in practice a broad classification
// and HTML "obviously" falls within it; however HTML
// is separated out here because HTML5 doesn't have a
// really good definition (like DTD or XSD), and also
// because HTML in general has dedicated parsers that
// are quite a bit more liberal than XML parsers.
//
// NOTE that for source code that uses this enum, files
// can be named "*_xml.go", "*_html.go", "*_mkdn.go".
// .
type MarkupType string

const (
	MU_type_UNK MarkupType = "UNK" // UNKNOWN

	MU_type_XML  = "XML"
	MU_type_HTML = "HTML" // Assumed to be HTML5
	MU_type_MKDN = "MKDN" // Assumed to be CommonMark (or GFM?)
	MU_type_BIN  = "BIN"  // Opaque
)

func (MUT MarkupType) LongForm() string {
	if len(MUT) < 3 || len(MUT) > 4 {
		panic("Bad MU_type: " + MUT)
	}
	switch MUT {
	case MU_type_XML:
		return "XML"
	case MU_type_HTML:
		return "HTML[5]"
	case MU_type_MKDN:
		return "Markdown"
	case MU_type_BIN:
		return "Binary"
	case MU_type_UNK:
		return "Unknown"
	}
	return string(MUT)
}
