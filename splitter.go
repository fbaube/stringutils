package stringutils

// Splitter is a common pattern, but it has to be a method of a struct.
// So for example to split off a XML prolog, it is necessary to define
// an associated struct, which might as well be a new empty XmlProlog,
// in which case it could also be modified by the method, if it is
// implemented as a pointer receiver.
//
// We could define it wherever the input and outputs are well defined:
// RawAttList => []RawAtt; RawAtt => name,"=",value; etc. 
// 
type Splitter interface {
	// TryToSplit is so named because "Try" 
	// says, that it can fail and return `nil`. 
	TryToSplit(s string) (ss []string)
}
