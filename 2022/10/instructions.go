package _10

import (
	"strings"
)

type Instruction interface {
	Run(clock <-chan struct{}, done chan<- struct{}, registry Registry)
}

func ParseInstruction(l string) Instruction {
	parts := strings.Split(l, " ")
	cmd := parts[0]
	args := parts[1:]
	switch cmd {
	case "noop":
		return NewNoopInstruction(args...)
	case "addx":
		return NewAddXInstruction(args...)
	}
	return nil
}
