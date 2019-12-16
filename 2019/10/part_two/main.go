package main

import (
	"os"

	"github.com/maxcleme/adventofcode/2019/10/part_two/utils"
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

	s := utils.Best(*m)
	if s == nil {
		logrus.Fatal("cannot find best station position")
	}

	c := utils.NthVaporized(s, 200)
	if c == nil {
		logrus.Fatal("cannot find nth vaporized")
	}

	logrus.Info(c)
}
