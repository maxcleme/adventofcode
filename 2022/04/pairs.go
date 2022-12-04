package _04

import (
	"strconv"
	"strings"
)

type Pair struct {
	fromInclusive int
	toInclusive   int
}

func (p Pair) Includes(pair Pair) bool {
	if p.fromInclusive > pair.fromInclusive {
		return false
	}
	if p.toInclusive < pair.toInclusive {
		return false
	}
	return true
}

func (p Pair) Overlaps(pair Pair) bool {
	return p.fromInclusive <= pair.toInclusive && pair.fromInclusive <= p.toInclusive
}

func ParsePair(s string) (Pair, error) {
	bounds := strings.Split(s, "-")
	from, err := strconv.Atoi(bounds[0])
	if err != nil {
		return Pair{}, err
	}
	to, err := strconv.Atoi(bounds[1])
	if err != nil {
		return Pair{}, err
	}
	return Pair{
		fromInclusive: from,
		toInclusive:   to,
	}, nil
}

func ParsePairs(s string) (Pair, Pair, error) {
	ps := strings.Split(s, ",")
	p1, err := ParsePair(ps[0])
	if err != nil {
		return Pair{}, Pair{}, err
	}
	p2, err := ParsePair(ps[1])
	if err != nil {
		return Pair{}, Pair{}, err
	}
	return p1, p2, nil
}
