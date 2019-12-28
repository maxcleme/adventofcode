package main

import (
	"os"

	"github.com/maxcleme/adventofcode/2019/14/part_one/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	// read input file
	file, err := os.Open("./input")
	if err != nil {
		logrus.WithError(err).Fatal("cannot read file")
	}
	defer file.Close()

	recipes, err := utils.Parse(file)
	if err != nil {
		logrus.WithError(err).Fatal("cannot parse recipes")
	}

	c, err := utils.FuelToOre(recipes)
	if err != nil {
		logrus.WithError(err).Fatal("cannot compute minimum ore required for one fuel unit")
	}
	logrus.Info(c)
}
