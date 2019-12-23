package main

import (
	"os"

	"github.com/maxcleme/adventofcode/2019/12/part_two/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	// read input file
	file, err := os.Open("./input")
	if err != nil {
		logrus.WithError(err).Fatal("cannot read file")
	}
	defer file.Close()

	objects, err := utils.Parse(file)
	if err != nil {
		logrus.WithError(err).Fatal("cannot parse objects")
	}

	logrus.Info(utils.NextIdentity(objects))
}
