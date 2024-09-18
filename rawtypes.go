package stringutils

import(
	D "github.com/fbaube/dsmnd"
)

type Raw_type D.SemanticFieldType 

const(
	// Raw_type_BIN is opaque 
	Raw_type_BIN  Raw_type = Raw_type(D.SFT_BLOB_)
	// Raw_type_XML is assumed to be well-formed
	Raw_type_XML  Raw_type = Raw_type(D.SFT_XTEXT)
	// Raw_type_HTML is assumed to be HTML5
	Raw_type_HTML Raw_type = Raw_type(D.SFT_HTEXT)
	// Raw_type_MKDN is assumed to be CommonMark (or GFM?)
	Raw_type_MKDN Raw_type = Raw_type(D.SFT_MTEXT)
	// Raw_type_SQL is, well, hey why not eh
	Raw_type_SQL  Raw_type = Raw_type(D.SFT_QTEXT)
	// Raw_type_DIRLIKE is a hack placeholder for consistent 
        // handling (because of consistent problems in code): 
        // IsDirlike is a more general case of IsDir() - 
        // shorthand for "is not possible to have own content" -
	// but this can be more formally defined as "is/has link(s)
	// to other stuff" - i.e. it is a directory or symbolic link. 
        // Used by [ctoken.TypedRaw].
	Raw_type_DIRLIKE Raw_type = "dirlk"
)

/*
{BDT_BLOB.DT(), SFT_BLOB_.S(), "Blob", "Binary large object (program / data"},
// TEXTS (10)
{BDT_TEXT.DT(), SFT_STRNG.S(), "String", "Generic string, not readable text"},
{BDT_TEXT.DT(), SFT_TOKEN.S(), "Token", "Generic token or datum tag (no spaces or punc.)"},
{BDT_TEXT.DT(), SFT_FTEXT.S(), "Free-text", "Generic free-flowing readable text, format unspecified"},
{BDT_TEXT.DT(), SFT_JTEXT.S(), "JSON", "JSON text"},
{BDT_TEXT.DT(), SFT_XTEXT.S(), "XML-text", "XML text (incl fragments)"},
{BDT_TEXT.DT(), SFT_HTEXT.S(), "HTML5-text", "HTML[5!] text"},
{BDT_TEXT.DT(), SFT_MTEXT.S(), "Markdown", "Markdown/plaintext, incl LwDITA Ext'd-MDITA"},
{BDT_TEXT.DT(), SFT_ATEXT.S(), "Asciidoc", "Asciidoc text"},
{BDT_TEXT.DT(), SFT_MCFMT.S(), "Microformat", "Microformat record"},
{BDT_TEXT.DT(), SFT_QTEXT.S(), "SQL", "SQL (dialect TBS)"},
*/