package scrabble

import (
	"strings"
)

// Score receives a word and returns its Scrabble score
func Score(word string) int {

	word = strings.ToUpper(word)
	score := 0

	for _, c := range word {
		switch string(c) {
		case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
			score++
		case "D", "G":
			score += 2
		case "B", "C", "M", "P":
			score += 3
		case "F", "H", "V", "W", "Y":
			score += 4
		case "K":
			score += 5
		case "J", "X":
			score += 8
		case "Q", "Z":
			score += 10
		}
	}
	return score
}
