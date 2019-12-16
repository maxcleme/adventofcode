package main

import (
	"os"

	"github.com/maxcleme/adventofcode/2019/10/part_one/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	// read input file
	file, err := os.Open("./input")
	if err != nil {
		logrus.WithError(err).Fatal("cannot read file")
	}
	defer file.Close()

	m, err := utils.Parse(file)
	if err != nil {
		logrus.WithError(err).Fatal("cannot parse map")
	}

	origin, count := utils.Best(*m)
	if origin == nil {
		logrus.Fatal("there is no best place")
	}

	logrus.
		WithField("origin", origin).
		WithField("count", count).
		Info("found")

}
