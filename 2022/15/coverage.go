package _15

import (
	"math"
	"sort"
)

func Locate(sensors []Sensor, min, max int) Position {
	for i := min; i <= max; i++ {
		r, _ := RowCoverage(sensors, min, max, i)
		if len(r) != 2 {
			continue
		}
		if space := r[1][0] - r[0][1]; space != 2 {
			continue
		}
		return Position{r[1][0] - 1, i}
	}
	return Position{}
}

func RowCoverage(sensors []Sensor, min, max, n int) ([][2]int, int) {
	var covered [][2]int
	for _, sensor := range sensors {
		d := Distance(sensor.Position, sensor.Beacon)
		if !(sensor.Y-d <= n && n <= sensor.Y+d) {
			continue
		}
		covered = append(covered, [2]int{
			int(math.Max(float64(min), float64(sensor.X-(d-int(math.Abs(float64(n-sensor.Y))))))),
			int(math.Min(float64(max), float64(sensor.X+(d-int(math.Abs(float64(n-sensor.Y))))))),
		})
	}
	return coverage(covered)
}

func coverage(covered [][2]int) ([][2]int, int) {
	sort.Slice(covered, func(i, j int) bool {
		if covered[i][0] == covered[j][0] {
			return covered[i][1] < covered[j][1]
		}
		return covered[i][0] < covered[j][0]
	})

	var c int
	var merged [][2]int
	for i := 0; i < len(covered); i++ {
		if len(merged) == 0 {
			merged = append(merged, covered[i])
			continue
		}
		if covered[i][0] <= merged[c][1] {
			merged[c][1] = int(math.Max(float64(merged[c][1]), float64(covered[i][1])))
			continue
		}
		merged = append(merged, covered[i])
		c++
	}

	var sum int
	for _, p := range merged {
		sum += p[1] - p[0]
	}
	return merged, sum
}

func Distance(p1, p2 Position) int {
	return int(math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y)))
}
