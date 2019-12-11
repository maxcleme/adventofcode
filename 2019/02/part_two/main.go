package main

import (
	"bufio"
	"os"
	"strconv"

	"github.com/maxcleme/adventofcode/2019/02/part_two/utils"
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
	instructions := make([]int, 0)
	for scanner.Scan() {
		instruction, err := strconv.Atoi(scanner.Text())
		if err != nil {
			logrus.WithError(err).Fatal("cannot parse instruction")
		}
		instructions = append(instructions, instruction)
	}

	// try all combinations
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			instructions[1] = noun
			instructions[2] = verb

			results, err := utils.Run(instructions)
			if err != nil {
				logrus.WithError(err).Fatal("cannot run program")
			}

			if results[0] == 19690720 {
				logrus.
					WithField("noun", noun).
					WithField("verb", verb).
					Info("combination found")
				break
			}
		}
	}

}
