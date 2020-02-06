package utils

import (
	"bufio"
	"io"
	"strconv"
)

const (
	ErrInvalidSequence = Error("sequence is invalid")
)

type Error string

func (e Error) Error() string {
	return string(e)
}

type Sequence []int

type Pattern []int

func Repeat(sequence Sequence, times int) Sequence {
	buff := make([]int, len(sequence)*times)
	for i := 0; i < times; i++ {
		for j := 0; j < len(sequence); j++ {
			buff[i*len(sequence)+j] = sequence[j]
		}
	}
	return buff
}

func Offset(sequence Sequence) (int, error) {
	if len(sequence) < 7 {
		return 0, ErrInvalidSequence
	}

	o := sequence[:7]
	offset := 0
	offset += o[6]
	offset += o[5] * 10
	offset += o[4] * 100
	offset += o[3] * 1000
	offset += o[2] * 10000
	offset += o[1] * 100000
	offset += o[0] * 1000000
	return offset, nil
}

func Parse(reader io.Reader) (Sequence, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	raw := scanner.Text()

	if len(raw) == 0 {
		return nil, ErrInvalidSequence
	}

	seq := make([]int, len(raw))
	for i, c := range raw {
		v, err := strconv.Atoi(string(c))
		if err != nil {
			return nil, ErrInvalidSequence
		}
		seq[i] = v
	}
	return seq, nil
}

func Apply(sequence Sequence, phaseCount, repeat, offset int) Sequence {
	// further algorithm works only if requested offset is after the mid point of the repeated sequence
	if offset < (len(sequence)*repeat)/2 {
		return nil
	}

	current := Repeat(sequence, repeat)
	for i := 0; i < phaseCount; i++ {
		apply(current, offset)
	}
	return current[offset : offset+8]
}

func apply(sequence Sequence, offset int) Sequence {
	if sequence == nil {
		return nil
	}

	sum := 0
	for i := len(sequence) - 1; i >= offset; i-- {
		sum += sequence[i]
		sequence[i] = sum%10
	}

	return sequence
}
