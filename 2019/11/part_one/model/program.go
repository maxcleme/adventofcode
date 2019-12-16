package model

import "io"

type Program struct {
	Statements []int
	In         io.Reader
	Out        io.Writer
	Base       int
}
