package main

import (
	"os"

	"github.com/maxcleme/adventofcode/2019/16/part_one/utils"
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

	output := utils.Apply(*sequence, utils.Pattern{0, 1, 0, -1}, 100)
	logrus.Info((*output)[:8])
}
