package bob

import (
	"regexp"
	"strings"
)

// Hey returns takes a string remark and returns a string response
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if IsEmpty(remark) {
		return "Fine. Be that way!"
	}
	if IsYelling(remark) && IsQuestion(remark) {
		return "Calm down, I know what I'm doing!"
	}
	if IsQuestion(remark) {
		return "Sure."
	}
	if IsYelling(remark) {
		return "Whoa, chill out!"
	}

	return "Whatever."
}

// IsQuestion tests a string to see if it is a question
func IsQuestion(sentence string) bool {
	return strings.HasSuffix(sentence, "?")
}

// IsExclamation tests a string to see if it is a exclamation
func IsExclamation(sentence string) bool {
	return strings.HasSuffix(sentence, "!")
}

// IsEmpty tests for an empty string
func IsEmpty(remark string) bool {
	if len(remark) == 0 {
		return true
	} else {
		return false
	}
}

// IsYelling tests a string to see if it contains only uppercase letters
func IsYelling(remark string) bool {
	containsLetters, _ := regexp.MatchString("[A-Z]", remark)
	if remark == strings.ToUpper(remark) && containsLetters {
		return true
	} else {
		return false
	}
}
