package utils

import "strconv"

func Validate(password int) bool {
	if password < 100000 {
		return false
	}
	if password > 999999 {
		return false
	}

	digits := strconv.Itoa(password)
	last := rune(digits[0])
	double := false
	for _, current := range []rune(digits[1:]) {
		if last > current {
			return false
		}
		if last == current {
			double = true
		}
		last = current
	}

	return double
}
