package diffsquares

// SquareOfSums receives a number and takes a range from 1 to
// number, adds each number in the range then returns the square
// of the sum.
func SquareOfSums(number int) int {
	sum := 0
	for i := 1; i <= number; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares receives a number and takes a range from 1 to
// number, adds the square of each number in the range then returns
// the sum.
func SumOfSquares(number int) int {
	sum := 0
	for i := 1; i <= number; i++ {
		sum += i * i
	}
	return int(sum)
}

// Difference receives a number and returns the difference
// between the square of their sums and the sum of their squares.
func Difference(number int) int {
	return SquareOfSums(number) - SumOfSquares(number)
}
