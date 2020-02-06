package utils

import (
	"bufio"
	"io"
	"strconv"
	"unicode"
)

const (
	ErrInvalidSequence = Error("sequence is invalid")
)

type Error string

func (e Error) Error() string {
	return string(e)
}

type Sequence string

type Pattern []int

func Parse(reader io.Reader) (*Sequence, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	raw := scanner.Text()

	if len(raw) == 0 {
		return nil, ErrInvalidSequence
	}

	for _, r := range []rune(raw) {
		if !unicode.IsDigit(r) {
			return nil, ErrInvalidSequence
		}
	}
	s := Sequence(raw)
	return &s, nil
}

func Apply(sequence Sequence, pattern Pattern, phaseCount int) *Sequence {
	current := sequence
	for i := 0; i < phaseCount; i++ {
		r := apply(current, pattern)
		current = *r
	}
	return &current
}

func apply(sequence Sequence, pattern Pattern) *Sequence {
	result := make([]rune, len(sequence))

	for times := 0; times < len(sequence); times++ {
		coeffs := coeff(pattern, times+1, len(sequence))
		sum := 0
		for i, r := range []rune(sequence) {
			v, _ := strconv.Atoi(string(r))
			sum += v * coeffs[i%len(coeffs)]
		}
		itoa := strconv.Itoa(sum)
		result[times] = rune(itoa[len(itoa)-1])
	}

	s := Sequence(string(result))
	return &s
}

func coeff(pattern Pattern, times, length int) []int {
	coeffs := make([]int, 0)
	for {
		for _, c := range pattern {
			for i := 0; i < times; i++ {
				coeffs = append(coeffs, c)
				if len(coeffs) == length+1 {
					return coeffs[1:]
				}
			}
		}
	}
}
