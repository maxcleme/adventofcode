package main

import (
	"os"

	"github.com/maxcleme/adventofcode/2019/12/part_one/utils"
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

	objects = utils.Simulate(objects, 1000)

	total := 0
	for _, o := range objects {
		total += o.Total()
	}
	logrus.Info(total)
}
