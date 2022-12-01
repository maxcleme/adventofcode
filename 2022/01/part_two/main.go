package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// read input file
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	var elves []int
	sum := 0
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			elves = append(elves, sum)
			sum = 0
			continue
		}
		value, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		sum += value
	}
	elves = append(elves, sum)
	sort.Ints(elves)
	fmt.Println(elves[len(elves)-3] + elves[len(elves)-2] + elves[len(elves)-1])
}
