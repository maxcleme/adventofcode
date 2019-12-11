package main

import (
	"bufio"
	"os"
	"strconv"

	"github.com/maxcleme/adventofcode/2019/02/part_one/utils"
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

	// according to rules we need to do the following actions before running the input program
	// - replace position 1 with the value 12
	// - replace position 2 with the value 2
	instructions[1] = 12
	instructions[2] = 2

	results, err := utils.Run(instructions)
	if err != nil {
		logrus.WithError(err).Fatal("cannot run program")
	}

	logrus.Info(results[0])
}
