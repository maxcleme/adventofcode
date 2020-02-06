package main

import (
	"os"

	"github.com/maxcleme/adventofcode/2019/16/part_two/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	// read input file
	file, err := os.Open("./input")
	if err != nil {
		logrus.WithError(err).Fatal("cannot read file")
	}
	defer file.Close()

	sequence, err := utils.Parse(file)
	if err != nil {
		logrus.WithError(err).Fatal("cannot parse sequence")
	}

	offset, err := utils.Offset(sequence)
	if err != nil {
		logrus.WithError(err).Fatal("cannot find offset")
	}
	output := utils.Apply(sequence, 100, 10000, offset)
	logrus.Info(output)
}
