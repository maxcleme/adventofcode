package _01

type lock struct {
	position int
	part1    int
	part2    int
}

func (l *lock) turn(direction string, step int) {
	var crossings int
	if direction == "R" {
		if l.position == 0 {
			crossings = step / 100
		} else if step >= (100 - l.position) {
			crossings = 1 + (step-(100-l.position))/100
		}
		l.position = (l.position + step) % 100
	} else if direction == "L" {
		if l.position == 0 {
			crossings = step / 100
		} else if step >= l.position {
			crossings = 1 + (step-l.position)/100
		}
		l.position = ((l.position-step)%100 + 100) % 100
	}
	if l.position == 0 {
		l.part1++
	}
	l.part2 += crossings
}
