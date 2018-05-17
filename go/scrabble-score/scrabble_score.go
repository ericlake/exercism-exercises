package scrabble

import (
	"strings"
)

// Score receives a word and returns its Scrabble score
func Score(word string) int {

	// There has got to be a more idomatic way of doing the following
	// I think is should be douable as a const as well.
	m := make(map[string]int)
	m["A"] = 1
	m["E"] = 1
	m["I"] = 1
	m["O"] = 1
	m["U"] = 1
	m["L"] = 1
	m["N"] = 1
	m["R"] = 1
	m["S"] = 1
	m["T"] = 1
	m["D"] = 2
	m["G"] = 2
	m["B"] = 3
	m["C"] = 3
	m["M"] = 3
	m["P"] = 3
	m["F"] = 4
	m["H"] = 4
	m["V"] = 4
	m["W"] = 4
	m["Y"] = 4
	m["K"] = 5
	m["J"] = 8
	m["X"] = 8
	m["Q"] = 10
	m["Z"] = 10

	word = strings.ToUpper(word)
	score := 0

	if word == "" {
		return score
	}
	for _, c := range word {
		score += m[string(c)]
	}
	return score
}
