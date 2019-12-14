package main

import (
	"bufio"
	"os"

	"github.com/maxcleme/adventofcode/2019/06/part_two/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	// read input file
	file, err := os.Open("./input")
	if err != nil {
		logrus.WithError(err).Fatal("cannot read file")
	}
	defer file.Close()

	input := make([]string, 0)

	// scan line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	count, err := utils.Between(input, "YOU", "SAN")
	if err != nil {
		logrus.WithError(err).Fatal("cannot compute total orbits number")
	}

	logrus.Info(count)
}
