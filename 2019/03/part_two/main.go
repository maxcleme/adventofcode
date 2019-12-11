package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/maxcleme/adventofcode/2019/03/part_two/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	// read input file
	file, err := os.Open("./input")
	if err != nil {
		logrus.WithError(err).Fatal("cannot read file")
	}
	defer file.Close()

	wireSegments := make([][]string, 0)

	// scan line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		segments := strings.Split(scanner.Text(), ",")
		wireSegments = append(wireSegments, segments)
	}

	// output result
	delay, err := utils.MinimalDelay(wireSegments[0], wireSegments[1])
	if err != nil {
		logrus.WithError(err).Fatal("cannot compute minimal delay of intersection with origin")
	}
	logrus.Info(delay)
}
