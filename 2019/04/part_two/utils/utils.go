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
	sequenceSize := 1
	double := false
	for _, current := range []rune(digits[1:]) {
		if last > current {
			return false
		}
		if last == current {
			sequenceSize++
		} else {
			if sequenceSize == 2 {
				double = true
			}
			sequenceSize = 1
		}
		last = current
	}

	return double || sequenceSize == 2
}
