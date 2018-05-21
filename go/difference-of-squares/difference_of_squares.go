package diffsquares

import (
	"math"
)

// SquareOfSums receives a number and takes a range from 1 to
// number, adds each number in the range then returns the square
// of the sum.
func SquareOfSums(number int) int {
	sum := 0
	for i := 1; i <= number; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), float64(2)))
}

// SumOfSquares receives a number and takes a range from 1 to
// number, adds the square of each number in the range then returns
// the sum.
func SumOfSquares(number int) int {
	var sum float64
	sum = 0
	for i := 1; i <= number; i++ {
		sum += math.Pow(float64(i), float64(2))
	}
	return int(sum)
}

// Difference receives a number and returns the difference
// between the square of their sums and the sum of their squares.
func Difference(number int) int {
	return SquareOfSums(number) - SumOfSquares(number)
}
