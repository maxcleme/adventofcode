package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/maxcleme/adventofcode/2019/13/part_two/model"
	"github.com/maxcleme/adventofcode/2019/13/part_two/utils"
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

	cr := make(chan []byte)
	reader := utils.NewChanReader(cr)
	cw := make(chan []byte)
	writer := utils.NewChanWriter(cw)

	p := model.Program{
		Statements: statements,
		In:         reader,
		Out:        writer,
		Base:       0,
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	height := 22
	width := 37
	score := 0
	tiles := make([][]int, height)
	ball := [2]int{}
	paddle := [2]int{}
	delay := time.Millisecond * 33

	// start renderer
	go func() {
		for {
			// clear term
			fmt.Print("\033[H\033[2J")
			// print map
			for _, row := range tiles {
				for _, tileID := range row {
					switch tileID {
					case 0:
						fmt.Print(" ")
					case 1:
						fmt.Print("#")
					case 2:
						fmt.Print("■")
					case 3:
						fmt.Print("‾")
					case 4:
						fmt.Print("O")
					}
				}
				fmt.Println()
			}
			fmt.Printf("Score : %d\n", score)
			time.Sleep(delay)
		}
	}()

	// start user bot
	go func() {
		for {
			time.Sleep(delay)
			if paddle[0] < ball[0] {
				cr <- []byte("1")
			} else if paddle[0] > ball[0] {
				cr <- []byte("-")
				cr <- []byte("1")
			} else {
				cr <- []byte("0")
			}
			cr <- []byte("\n")
		}
	}()

	// start engine
	go func() {
		for i := 0; i < height; i++ {
			tiles[i] = make([]int, width)
		}

		x, y, offset := 0, 0, 0
		for {
			o, ok := <-cw
			if !ok {
				break
			}

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
				if x >= 0 {
					tiles[y][x] = v
					if v == 3 {
						paddle[0], paddle[1] = x, y
					}
					if v == 4 {
						ball[0], ball[1] = x, y
					}
				} else {
					score = v
				}
			}
			offset++
		}
	}()

	// start program
	go func() {
		defer wg.Done()
		// set memory 0 to 2 in order to play for free
		p.Statements[0] = 2
		if _, err := utils.Run(p); err != nil {
			logrus.WithError(err).Fatal("cannot run program")
		}
	}()

	wg.Wait()
	logrus.WithField("final_score", score).Info()
}
