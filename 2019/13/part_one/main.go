package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/maxcleme/adventofcode/2019/13/part_one/model"
	"github.com/maxcleme/adventofcode/2019/13/part_one/utils"
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

	c := make(chan []byte)
	writer := utils.NewChanWriter(c)

	p := model.Program{
		Statements: statements,
		In:         os.Stdin,
		Out:        writer,
		Base:       0,
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		size := 100
		tiles := make([][]int, size)
		for i := 0; i < size; i++ {
			tiles[i] = make([]int, size)
		}

		x, y, offset := 0, 0, 0
		for {
			o, ok := <-c
			if !ok {
				break
			}
			logrus.Info(o)

			v, err := strconv.Atoi(string(o[:len(o)-1]))
			if err != nil {
				logrus.WithError(err).Fatal("cannot parse output")
			}

			switch offset % 3 {
			case 0:
				x = v
			case 1:
				y = v
			case 2:
				tiles[y][x] = v
			}
			offset++
			logrus.Info(offset)

		}

		block := 0
		for _, row := range tiles {
			for _, tileID := range row {
				if tileID == 2 {
					block++
				}
				fmt.Print(tileID)
			}
			fmt.Println()
		}
		logrus.WithField("block_count", block).Info()
	}()

	go func() {
		defer wg.Done()
		if _, err := utils.Run(p); err != nil {
			logrus.WithError(err).Fatal("cannot run program")
		}
		close(c)
	}()

	wg.Wait()
}
