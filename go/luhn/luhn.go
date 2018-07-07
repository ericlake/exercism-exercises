package luhn

// Valid recieves a string of numbers and returns a bool to show if it
// is a valid credit card number combination.
func Valid(input string) bool {
	total := 0
	switch len(input) {
	case 1:
		return false
	case 2:
		total += int(input[0]) * 2
		if total > 9 {
			total -= 9
		}
		total += int(input[1])
		if total%10 == 0 {
			return true
		}
	}

	for i, v := range input {
		if i%2 == 0 {
			poo := int(v) * 2
			if poo > 9 {
				total += poo - 9
			}
			total += poo
		}
		total += int(v)
	}
	if total%10 == 0 {
		return true
	}
	return false
}
