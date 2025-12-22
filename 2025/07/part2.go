package _07

import (
	"fmt"
	"os"

	"github.com/maxcleme/adventofcode/grid"
	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2025/07/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}

type D struct {
	S     string
	Count int
}

func (d D) String() string {
	if d.Count == 0 {
		return d.S
	}
	if d.Count < 10 {
		return fmt.Sprint(d.Count)
	}
	return "(" + fmt.Sprint(d.Count) + ")"
}

func part2(payload string) int {

	d := grid.New[D](payload, func(r rune) D {
		return D{S: string(r)}
	})
	start := d.Find(func(t *grid.Tile[D]) bool {
		return t.Value.S == "S"
	})
	start[0].Value.Count = 1

	var sources = start
	var newBeams []*grid.Tile[D]
	for {
		newBeams = []*grid.Tile[D]{}
		for _, src := range sources {
			b, ok := d.Get(src.X, src.Y+1)
			if !ok {
				continue
			}
			switch b.Value.S {
			case ".":
				b.Value.S = "|"
				b.Value.Count += src.Value.Count
				newBeams = append(newBeams, b)
			case "^":
				left, ok := d.Get(b.X-1, b.Y)
				if ok {
					left.Value.Count += src.Value.Count
					if left.Value.S == "." {
						left.Value.S = "|"
						newBeams = append(newBeams, left)
					}
				}
				right, ok := d.Get(b.X+1, b.Y)
				if ok {
					right.Value.Count += src.Value.Count
					if right.Value.S == "." {
						right.Value.S = "|"
						newBeams = append(newBeams, right)
					}
				}
			case "|":
				b.Value.Count += src.Value.Count
				continue
			}
		}
		if len(newBeams) == 0 {
			break
		}
		sources = newBeams
	}
	total := 0
	for i := 0; i < d.Width(); i++ {
		b, _ := d.Get(i, d.Height()-1)
		total += b.Value.Count
	}
	return total
}
