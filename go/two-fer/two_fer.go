package twofer

import "fmt"

// ShareWith takes an optional name and returns a
// string
func ShareWith(name string) string {
	if len(name) > 0 {
		return fmt.Sprintf("One for %s, one for me.", name)
	}
	return "One for you, one for me."
}
