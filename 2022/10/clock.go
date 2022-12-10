package _10

type Clock struct {
	callback func(n int)
	next     chan struct{}
}

func (c Clock) Start() <-chan struct{} {
	clock := make(chan struct{})
	go func() {
		cycle := 1
		for {
			select {
			case <-c.next:
				c.callback(cycle)
				clock <- struct{}{}
				cycle++
			}
		}
	}()
	go func() {
		c.next <- struct{}{}
	}()
	return clock
}

func NewClock(callback func(n int)) Clock {
	return Clock{
		callback: callback,
		next:     make(chan struct{}),
	}
}
