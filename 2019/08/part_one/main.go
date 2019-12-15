package main

import (
	"os"

	"github.com/maxcleme/adventofcode/2019/08/part_one/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	// read input file
	file, err := os.Open("./input")
	if err != nil {
		logrus.WithError(err).Fatal("cannot read file")
	}
	defer file.Close()

	img, err := utils.Parse(file, 6, 25)
	if err != nil {
		logrus.WithError(err).Fatal("cannot parse input")
	}

	// find layer with fewest digit
	min := img.Layers[0]
	for _, l := range img.Layers[1:] {
		if l.Colors[0] < min.Colors[0] {
			min = l
		}
	}

	// output number of 1 digits multiplied by the number of 2 digits
	logrus.Info(min.Colors[1] * min.Colors[2])
}
