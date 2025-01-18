package stringutils

import (
	FP "path/filepath"
)

// FilterStringsBySuffix takes a list of filenames and filters
// out those whose file extensions are not in the list that is
// passed in. NOTE:
//  - No periods on the okay file extensions.
//  - The comparison is case-INsensitive. 
//  - This func is not currently used anywhere (2025.01), 
//    but it's useful so leave it in anyways. 
func FilterStringsBySuffix(inputs []string,
	okayExts []string) (OKoutputs []string) {
	if okayExts == nil || len(okayExts) == 0 {
		return inputs
	}
	if inputs == nil || len(inputs) == 0 {
		return inputs
	}
	OKoutputs = make([]string, 0, len(inputs))
	for _, instring := range inputs {
		sfx := FP.Ext(instring)
		if IsInSliceIgnoreCase(sfx, okayExts) {
			OKoutputs = append(OKoutputs, instring)
		}
	}
	return OKoutputs
}
