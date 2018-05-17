package raindrops

import "strconv"

// Convert takes a number and returns funny stuff based on the factors
func Convert(number int) string {
	raindrop := ""

	if number%3 == 0 {
		raindrop += "Pling"
	}
	if number%5 == 0 {
		raindrop += "Plang"
	}
	if number%7 == 0 {
		raindrop += "Plong"
	}
	if raindrop == "" {
		return strconv.Itoa(number)
	}
	return raindrop
}
