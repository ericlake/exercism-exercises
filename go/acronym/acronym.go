package acronym

import (
	"strings"
)

// Abbreviate should receive a string and return a valid acronym.
func Abbreviate(s string) string {
	s = strings.ToUpper(strings.TrimSpace(s))
	replacer := strings.NewReplacer(",", "", ".", "", ";", "")
	s = replacer.Replace(s)
	words := strings.Fields(s)
	var acr string
	for _, word := range words {
		if strings.Contains(word, "-") {
			acr += strings.Split(word, "-")[0][0:1]
			acr += strings.Split(word, "-")[1][0:1]
			continue
		}
		acr += word[:1]
	}
	return acr
}
