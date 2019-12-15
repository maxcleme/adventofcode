package utils

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

const (
	ErrInvalidColor = Error("invalid color")
)

type Error string

func (e Error) Error() string {
	return string(e)
}

type Image struct {
	Pixels [][]*int

	Height int
	Width  int
}

var ColorMapping = map[int]string{
	0: "\033[1;30m%s\033[0m",
	1: "\033[1;37m%s\033[0m",
	2: "\033[1;35m%s\033[0m",
}

func (i *Image) String() string {
	o := ""
	for _, row := range i.Pixels {
		for _, pixel := range row {
			o += fmt.Sprintf(ColorMapping[*pixel], "█")
		}
		o += "\n"
	}
	return o
}

func Parse(reader io.Reader, height, width int) (*Image, error) {
	img := Image{
		Height: height,
		Width:  width,
		Pixels: make([][]*int, height),
	}

	for i := 0; i < height; i++ {
		img.Pixels[i] = make([]*int, width)
	}

	scanner := bufio.NewScanner(reader)

	max := height * width
	i := 0
	for scanner.Scan() {
		for _, b := range scanner.Bytes() {
			color, err := strconv.Atoi(string(b))
			if err != nil {
				return nil, ErrInvalidColor
			}

			x, y := (i%max)%width, (i/width)%height

			if img.Pixels[y][x] == nil || *img.Pixels[y][x] == 2 {
				img.Pixels[y][x] = &color
			}
			i++
		}
	}

	return &img, nil
}
