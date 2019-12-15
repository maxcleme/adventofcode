package utils

import (
	"bufio"
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

type Layer struct {
	// Count of color occurrence
	Colors map[int]int
}

type Image struct {
	Height int
	Width  int
	Layers []Layer
}

func Parse(reader io.Reader, height, width int) (*Image, error) {
	img := Image{
		Height: height,
		Width:  width,
		Layers: make([]Layer, 0),
	}

	scanner := bufio.NewScanner(reader)
	max := height * width
	curr := Layer{
		Colors: make(map[int]int),
	}

	tmp := 0
	for scanner.Scan() {
		for _, b := range scanner.Bytes() {

			if tmp != 0 && tmp%max == 0 {
				img.Layers = append(img.Layers, curr)
				curr = Layer{
					Colors: make(map[int]int),
				}
			}

			color, err := strconv.Atoi(string(b))
			if err != nil {
				return nil, ErrInvalidColor
			}

			if _, ok := curr.Colors[color]; !ok {
				curr.Colors[color] = 0
			}

			curr.Colors[color]++
			tmp++
		}
	}

	if len(curr.Colors) != 0 {
		img.Layers = append(img.Layers, curr)
	}

	return &img, nil
}
