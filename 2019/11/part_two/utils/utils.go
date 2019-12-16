package utils

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/maxcleme/adventofcode/2019/11/part_two/model"
	"github.com/maxcleme/adventofcode/2019/11/part_two/modes"
	"github.com/maxcleme/adventofcode/2019/11/part_two/opcode"
	"github.com/sirupsen/logrus"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrInvalidOpCode = Error("invalid opcode")
	ErrUnknownOpCode = Error("unknown opcode")
	ErrUnknownMode   = Error("unknown mode")
	ErrInvalidMode   = Error("invalid mode")

	ErrInvalidOrientation = Error("invalid orientation")
	ErrInvalidRotation    = Error("invalid rotation")

	ErrChannelClose  = Error("channel closed")
	ErrInvalidOutput = Error("invalid output")
)

type Robot struct {
	Statements  []int
	Camera      Camera
	X           int
	Y           int
	Orientation Orientation
}

type ChanReader struct {
	in chan []byte
}

func (r *ChanReader) Read(p []byte) (n int, err error) {
	b, ok := <-r.in
	if !ok {
		return 0, ErrChannelClose
	}

	if len(b) > len(p) {
		r.in <- b[len(p):]
		b = b[:len(p)]
	}

	copy(p, b)
	return len(b), nil
}

type ChanWriter struct {
	out chan []byte
}

func (w *ChanWriter) Write(p []byte) (n int, err error) {
	w.out <- p
	return len(p), nil
}

func (r *Robot) Start() error {
	in := &ChanReader{in: make(chan []byte, 1)}
	out := &ChanWriter{out: make(chan []byte, 1)}

	p := model.Program{
		Statements: r.Statements,
		In:         in,
		Out:        out,
		Base:       0,
	}

	// bootstrap first input
	in.in <- []byte(strconv.Itoa(int(r.Camera.Color(r.X, r.Y))) + "\n")

	exit := make(chan error, 1)

	wg := sync.WaitGroup{}
	// exit at first errors in goroutines
	wg.Add(1)

	go func(exit chan error) {
		defer wg.Done()
		if _, err := Run(p); err != nil {
			exit <- err
		}
	}(exit)

	go func(exit chan error) {
		defer wg.Done()
		for {
			// first input is the color to be paint at robot position
			color, ok := <-out.out
			if !ok {
				// chan closed, assume program is terminated
				return
			}
			var i int
			if _, err := fmt.Fscanf(strings.NewReader(string(color)), "%d\n", &i); err != nil {
				exit <- ErrInvalidOutput
				return
			}

			// paint current panel
			r.Camera.Map[r.Y][r.X].Color = Color(i)
			// set at visited
			r.Camera.Map[r.Y][r.X].Visited = true

			// second input is the next movement of the robot
			movement, ok := <-out.out
			if !ok {
				exit <- ErrChannelClose
				return
			}
			var j int
			if _, err := fmt.Fscanf(strings.NewReader(string(movement)), "%d\n", &j); err != nil {
				exit <- ErrInvalidOutput
				return
			}

			switch Rotation(j) {
			case RightRotation:
				r.Orientation = (r.Orientation + 1) % 4
			case LeftRotation:
				// support negative value
				r.Orientation = ((r.Orientation-1)%4 + 4) % 4
			default:
				exit <- ErrInvalidRotation
				return
			}

			switch r.Orientation {
			case Up:
				r.Y = r.Y - 1
			case Right:
				r.X = r.X + 1
			case Down:
				r.Y = r.Y + 1
			case Left:
				r.X = r.X - 1
			default:
				exit <- ErrInvalidOrientation
				return
			}

			// send current color to input reader
			in.in <- []byte(strconv.Itoa(int(r.Camera.Color(r.X, r.Y))) + "\n")
		}
	}(exit)

	wg.Wait()

	// look for errors
	select {
	case err := <-exit:
		if err != nil {
			return err
		}
	default:
		// everything fine!
	}

	fmt.Println(r.Camera.Map)
	count := 0
	// count panel painted at least once
	for _, row := range r.Camera.Map {
		for _, tile := range row {
			if tile.Visited {
				count++
			}
		}
	}

	logrus.
		WithField("at_least_one_tile_painted", count).
		Info("finished")
	return nil
}

type Camera struct {
	Map Map
}

type Map [][]Tile

func (m Map) String() string {
	s := ""
	for _, row := range m {
		for _, cell := range row {
			s += fmt.Sprintf(ColorMapping[cell.Color], "█")
		}
		s += "\n"
	}
	return s
}

func NewMap(width, height int) Map {
	m := make([][]Tile, height)
	for x := 0; x < height; x++ {
		m[x] = make([]Tile, width)
	}
	return m
}

func (c Camera) Color(x, y int) Color {
	return c.Map[y][x].Color
}

type Tile struct {
	Color   Color
	Visited bool
}

type Color int
type Orientation int
type Rotation int

const (
	Black Color = iota
	White
)

const (
	Up Orientation = iota
	Right
	Down
	Left
)

const (
	LeftRotation Rotation = iota
	RightRotation
)

var ColorMapping = map[Color]string{
	Black: "\033[1;30m%s\033[0m",
	White: "\033[1;37m%s\033[0m",
}

func ParseInstruction(instruction int) (*opcode.Instruction, error) {
	s := strconv.Itoa(instruction)
	if len(s) < 2 {
		s = "0" + s
	}

	runes := []rune(s)
	o, err := strconv.Atoi(string(runes[len(runes)-2:]))
	if err != nil {
		return nil, ErrInvalidOpCode
	}

	code, ok := opcode.Opcodes[o]
	if !ok {
		return nil, ErrUnknownOpCode
	}

	paramModes := []modes.Mode{
		modes.Position{},
		modes.Position{},
		modes.Position{},
	}

	for i, m := range reverse(runes[:len(runes)-2]) {
		mi, err := strconv.Atoi(string(m))
		if err != nil {
			return nil, ErrInvalidMode
		}
		mode, ok := modes.Modes[mi]
		if !ok {
			return nil, ErrUnknownMode
		}
		paramModes[i] = mode
	}

	return &opcode.Instruction{
		Opcode: code,
		Modes:  paramModes,
	}, nil

}

func Run(program model.Program) ([]int, error) {
	// clone model, also add more memory
	input := make([]int, 1024*1024+len(program.Statements)*2)
	for i, statement := range program.Statements {
		input[i] = statement
	}

	p := model.Program{
		Statements: input,
		In:         program.In,
		Out:        program.Out,
		Base:       0,
	}

	offset := 0
	for offset < len(p.Statements) {
		instruction, err := ParseInstruction(p.Statements[offset])
		if err != nil {
			return nil, err
		}

		next, err := instruction.Opcode.Evaluate(&p, offset, instruction.Modes)
		if err != nil {
			return nil, err
		}
		offset = next
	}
	// remove extra memory from output
	return p.Statements[:len(program.Statements)], nil
}

func reverse(runes []rune) []rune {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return runes
}
