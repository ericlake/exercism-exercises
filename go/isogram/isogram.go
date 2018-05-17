package isogram

import (
	"strings"
)

// IsIsogram takes a string and returns true if there are no duplicated
// characters excluding "-" and " ".
func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	seen := 0
	for _, c := range word {

		switch string(c) {
		case "-":
			continue
		case " ":
			continue
		}

		if strings.Count(word, string(c)) > 1 {
			seen++
		}
	}
	if seen > 0 {
		return false
	}
	return true
}
