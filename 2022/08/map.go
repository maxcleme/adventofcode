package _08

import (
	"strconv"
)

type Map [][]int

func ParseRow(s string) []int {
	var row []int
	for j := 0; j < len(s); j++ {
		height, _ := strconv.Atoi(string(s[j]))
		row = append(row, height)
	}
	return row
}

func (m Map) CountVisible() int {
	bounds := 4*len(m) - 4
	for i := 1; i < len(m)-1; i++ {
		for j := 1; j < len(m)-1; j++ {
			if m.Visible(i, j) {
				bounds++
			}
		}
	}
	return bounds
}

func (m Map) Visible(x, y int) bool {
	left := m.VisibleLeft(x, y)
	right := m.VisibleRight(x, y)
	up := m.VisibleUp(x, y)
	down := m.VisibleDown(x, y)
	return left || right || up || down
}

func (m Map) VisibleLeft(x, y int) bool {
	height := m[x][y]
	for i := y - 1; i >= 0; i-- {
		if m[x][i] >= height {
			return false
		}
	}
	return true
}

func (m Map) VisibleRight(x, y int) bool {
	height := m[x][y]
	for i := y + 1; i < len(m); i++ {
		if m[x][i] >= height {
			return false
		}
	}
	return true
}

func (m Map) VisibleUp(x, y int) bool {
	height := m[x][y]
	for i := x - 1; i >= 0; i-- {
		if m[i][y] >= height {
			return false
		}
	}
	return true
}

func (m Map) VisibleDown(x, y int) bool {
	height := m[x][y]
	for i := x + 1; i < len(m); i++ {
		if m[i][y] >= height {
			return false
		}
	}
	return true
}

func (m Map) BestScore() int {
	var max int
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			if score := m.Score(i, j); score > max {
				max = score
			}
		}
	}
	return max
}

func (m Map) Score(x, y int) int {
	left := m.ScoreLeft(x, y)
	right := m.ScoreRight(x, y)
	up := m.ScoreUp(x, y)
	down := m.ScoreDown(x, y)
	return left * right * up * down
}

func (m Map) ScoreLeft(x, y int) int {
	height := m[x][y]
	var score int
	for i := y - 1; i >= 0; i-- {
		if m[x][i] >= height {
			score++
			break
		}
		score++
	}
	return score
}

func (m Map) ScoreRight(x, y int) int {
	height := m[x][y]
	var score int
	for i := y + 1; i < len(m); i++ {
		if m[x][i] >= height {
			score++
			break
		}
		score++
	}
	return score
}

func (m Map) ScoreUp(x, y int) int {
	height := m[x][y]
	var score int
	for i := x - 1; i >= 0; i-- {
		if m[i][y] >= height {
			score++
			break
		}
		score++
	}
	return score
}

func (m Map) ScoreDown(x, y int) int {
	height := m[x][y]
	var score int
	for i := x + 1; i < len(m); i++ {
		if m[i][y] >= height {
			score++
			break
		}
		score++
	}
	return score
}
