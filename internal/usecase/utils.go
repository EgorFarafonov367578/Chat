package usecase

import "unicode"

func stringIsMeaningless(s string) bool {
	for _, c := range s {
		if !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}
