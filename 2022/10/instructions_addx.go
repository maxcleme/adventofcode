package _10

import "strconv"

type AddXInstruction struct {
	v int
}

func NewAddXInstruction(args ...string) *AddXInstruction {
	v, _ := strconv.Atoi(args[0])
	return &AddXInstruction{
		v: v,
	}
}

func (a AddXInstruction) Run(clock <-chan struct{}, done chan<- struct{}, registry Registry) {
	<-clock
	done <- struct{}{}
	<-clock
	registry["X"] += a.v
	done <- struct{}{}
}
