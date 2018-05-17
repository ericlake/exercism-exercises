package hamming

import "errors"

// Distance takes 2 strings and compares them for differences
// returning the number of differences between them
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Strings are not the same length")
	}
	diffCount := 0
	for i := range a {
		if string([]rune(a)[i]) != string([]rune(b)[i]) {
			diffCount++
		}
	}
	return diffCount, nil
}
