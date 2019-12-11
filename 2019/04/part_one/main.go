package main

import (
	"github.com/maxcleme/adventofcode/2019/04/part_one/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	lowerBound := 246540
	upperBound := 787419

	// brute force possible password, count valid one
	valid := 0
	for password := lowerBound; password <= upperBound; password++ {
		if utils.Validate(password) {
			valid += 1
		}
	}

	logrus.Info(valid)
}
