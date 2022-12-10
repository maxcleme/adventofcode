package _10

type NoopInstruction struct {
}

func NewNoopInstruction(_ ...string) *NoopInstruction {
	return &NoopInstruction{}
}

func (n NoopInstruction) Run(clock <-chan struct{}, done chan<- struct{}, _ Registry) {
	<-clock
	done <- struct{}{}
}
