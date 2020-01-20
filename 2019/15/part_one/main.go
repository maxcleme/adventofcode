package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/maxcleme/adventofcode/2019/15/part_one/model"
	"github.com/maxcleme/adventofcode/2019/15/part_one/utils"
	"github.com/sirupsen/logrus"
)

type Direction int

func (d Direction) Inverse() Direction {
	switch d {
	case Up:
		return Down
	case Right:
		return Left
	case Down:
		return Up
	case Left:
		return Right
	}
	return 0
}

func (d Direction) Apply(x, y int) (int, int) {
	switch d {
	case Up:
		return x, y - 1
	case Right:
		return x + 1, y
	case Down:
		return x, y + 1
	case Left:
		return x - 1, y
	}
	// unknown movement, do nothing
	return x, y
}

var (
	Up    = Direction(1)
	Down  = Direction(2)
	Left  = Direction(3)
	Right = Direction(4)
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

	feedback := make(chan bool)

	p := model.Program{
		Statements: statements,
		In:         reader,
		Out:        writer,
		Base:       0,
	}

	width := 42
	height := 42

	tiles := make([][]string, height)
	for i := 0; i < height; i++ {
		tiles[i] = make([]string, width)
	}

	distances := make([][]int, height)
	for i := 0; i < height; i++ {
		distances[i] = make([]int, width)
		for j := 0; j < width; j++ {
			distances[i][j] = -1
		}
	}

	ox, oy := -1, -1
	sx, sy := width/2, height/2

	state := EngineState{
		x:                sx,
		y:                sy,
		inputReceiver:    cr,
		feedback:         feedback,
		lastDirection:    Up,
		traveledDistance: 1,
		distances:        distances,
	}

	// start user bot
	go func(state *EngineState) {
		exploreAll(state, AllDirections())
		close(state.inputReceiver)
	}(&state)

	// start engine
	go func(state *EngineState) {
		for o := range cw {
			v, err := strconv.Atoi(string(o[:len(o)-1]))
			if err != nil {
				logrus.WithError(err).Fatal("cannot parse output")
			}

			hitWall := false
			// computed expected position
			ex, ey := state.lastDirection.Apply(state.x, state.y)
			switch v {
			case 0:
				// bot has hit a wall
				// mark the wall in the current tiles map
				// do not update bot position since it hasn't move
				tiles[ey][ex] = "#"
				hitWall = true
			case 1, 2:
				// bot has moved to the expected position
				// mark the current position as visited
				// move the bot in the current tiles map
				if v == 1 {
					tiles[state.y][state.x] = "."
				} else {
					tiles[state.y][state.x] = "O"
					ox, oy = state.x, state.y
				}
				state.x, state.y = ex, ey
			}
			// print map
			feedback <- hitWall
		}
	}(&state)

	utils.Run(p)
	display(state.x, state.y, ox, oy, tiles)

	if ox == -1 {
		logrus.Fatal("oxygen capsule not found")
	}
	logrus.
		WithField("oxygen_x", ox).
		WithField("oxygen_y", oy).
		WithField("start_x", sx).
		WithField("start_y", sy).
		WithField("distance", state.distances[oy][ox]).
		Info("")
}

type EngineState struct {
	x, y             int
	inputReceiver    chan []byte
	feedback         chan bool
	lastDirection    Direction
	traveledDistance int
	distances        [][]int
}

func display(dx, dy, ox, oy int, tiles [][]string) {
	// clear term
	for i, row := range tiles {
		for j, v := range row {
			if j == dx && i == dy {
				fmt.Print("D")
			} else if j == ox && i == oy {
				fmt.Print("O")
			} else {
				switch v {
				case "":
					fmt.Print(" ")
				default:
					fmt.Print(v)
				}
			}
		}
		fmt.Println()
	}
}

func try(state *EngineState, direction Direction) bool {
	state.lastDirection = direction

	// forge and send input
	state.inputReceiver <- []byte(strconv.Itoa(int(direction)))
	state.inputReceiver <- []byte("\n")

	// wait for engine feedback
	hitWall, _ := <-state.feedback
	return hitWall
}

func explore(state *EngineState, direction Direction) {
	if hitWall := try(state, direction); hitWall {
		// nothing to do since a willhit result in no position update
	} else {
		// if bot can move, explore the tile reached without current move
		// prevent backward movement
		state.traveledDistance++
		exploreAll(state, DirectionsExcept(direction.Inverse()))

		// come back to previous tile
		try(state, direction.Inverse())
		state.traveledDistance--
	}
}

func exploreAll(state *EngineState, directions []Direction) {
	state.distances[state.y][state.x] = state.traveledDistance
	for _, d := range directions {
		explore(state, d)
	}
}

func AllDirections() []Direction {
	return []Direction{Up, Down, Left, Right}
}

func DirectionsExcept(direction Direction) []Direction {
	directions := AllDirections()
	res := make([]Direction, 0)
	for _, d := range directions {
		if d != direction {
			res = append(res, d)
		}
	}
	return res
}
