package main

import (
	"bufio"
	"os"
	"strconv"

	"github.com/maxcleme/adventofcode/2019/11/part_two/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	// read input file
	file, err := os.Open("./input")
	if err != nil {
		logrus.WithError(err).Fatal("cannot read file")
	}
	defer file.Close()

	// read input
	scanner := bufio.NewScanner(file)
	// Define a split function that separates on commas.
	// According to https://golang.org/pkg/bufio/#Scanner examples
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)

	// scan values separated by coma
	statements := make([]int, 0)
	for scanner.Scan() {
		statement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			logrus.WithError(err).Fatal("cannot parse statement")
		}
		statements = append(statements, statement)
	}

	width := 100
	height := 15

	c := utils.Camera{
		Map: utils.NewMap(width, height),
	}

	r := utils.Robot{
		Statements:  statements,
		Camera:      c,
		X:           width / 2,
		Y:           height / 2,
		Orientation: utils.Up,
	}


	// set starting tile to white color
	r.Camera.Map[r.Y][r.X].Color = utils.White

	// start model
	if err := r.Start(); err != nil {
		logrus.
			WithError(err).
			Fatal("cannot run robot")
	}

}
