package main

import (
	"bufio"
	"os"
	"strconv"

	"github.com/maxcleme/adventofcode/2019/01/part_one/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	// read input file
	file, err := os.Open("./input")
	if err != nil {
		logrus.WithError(err).Fatal("cannot read file")
	}
	defer file.Close()

	// scan line by line
	scanner := bufio.NewScanner(file)
	sum := 0.0
	for scanner.Scan() {
		// compute mass
		mass, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			logrus.WithError(err).Fatal("cannot parse mass")
		}
		// sum require fuel
		sum += utils.RequiredFuel(mass)
	}

	// output result
	logrus.Info(sum)
}
