package main

import (
	"fmt"
	"os"

	"github.com/maxcleme/adventofcode/2019/08/part_two/utils"
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

	fmt.Println(img)
}
