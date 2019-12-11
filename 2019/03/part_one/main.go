package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/maxcleme/adventofcode/2019/03/part_one/utils"
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
	distance, err := utils.ClosestManhattanDistance(wireSegments[0], wireSegments[1])
	if err != nil {
		logrus.WithError(err).Fatal("cannot compute minimal manhattan distance of closest intersection with origin")
	}
	logrus.Info(distance)
}
