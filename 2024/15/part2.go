package _15

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func part2(input string) int {
	parts := strings.Split(input, "\n\n")
	var entities []*Entity
	var robot *Entity
	for y, row := range strings.Split(parts[0], "\n") {
		x := 0
		for _, c := range row {
			switch EntityKind(c) {
			case ROBOT:
				robot = newEntity(x, y, 1, 1, ROBOT)
				entities = append(entities, robot)
			case WALL:
				entities = append(entities, newEntity(x, y, 1, 1, WALL))
				entities = append(entities, newEntity(x+1, y, 1, 1, WALL))
			case BOX:
				entities = append(entities, newEntity(x, y, 2, 1, BOX))
			}
			x += 2
		}
	}
	for _, action := range parseActions(parts[1]) {
		move(entities, robot, action, false)
	}
	return gps(entities)
}

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/15/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}
