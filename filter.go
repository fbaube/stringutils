package stringutils

import (
	fp "path/filepath"
)

// FilterStringsBySuffix is TBS.
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
		sfx := fp.Ext(instring)
		if IsInSliceIgnoreCase(sfx, okayExts) {
			OKoutputs = append(OKoutputs, instring)
		}
	}
	return OKoutputs
}
