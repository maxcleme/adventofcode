package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// read input file
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	max := 0
	sum := 0
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			sum = 0
			continue
		}
		value, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		sum += value
		if sum > max {
			max = sum
		}
	}
	fmt.Println(max)
}
