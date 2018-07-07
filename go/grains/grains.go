package grains

import (
	"errors"
)

// Square returns the number of grains on any given square
func Square(square int) (uint64, error) {
	if square <= 0 || square > 64 {
		return uint64(square), errors.New("squares must be between 1 and 64")
	}
	//return uint64(math.Pow(2, (float64(square) - 1))), nil
	// Use bitwise operations instead of math.Pow()
	return 1 << uint64(square-1), nil
}

// Total returns the number of grains across all squares
func Total() uint64 {
	sum := uint64(0)

	for i := 1; i <= 64; i++ {
		square, _ := Square(i)
		sum += square
	}
	return sum
}
