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

