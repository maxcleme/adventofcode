package _15

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Action string

const (
	UP    Action = "^"
	DOWN  Action = "v"
	LEFT  Action = "<"
	RIGHT Action = ">"
)

type Entity struct {
	ID            int
	X, Y          int
	Width, Height int
	Kind          EntityKind
}

func (e Entity) Contains(x, y, width, height int) bool {
	return e.X < x+width && x < e.X+e.Width && e.Y < y+height && y < e.Y+e.Height
}

type EntityKind string

const (
	ROBOT EntityKind = "@"
	WALL  EntityKind = "#"
	BOX   EntityKind = "O"
)

func part1(input string) int {
	parts := strings.Split(input, "\n\n")
	var entities []*Entity
	var robot *Entity
	for y, row := range strings.Split(parts[0], "\n") {
		for x, c := range row {
			switch EntityKind(c) {
			case ROBOT:
				robot = newEntity(x, y, 1, 1, ROBOT)
				entities = append(entities, robot)
			default:
				entities = append(entities, newEntity(x, y, 1, 1, EntityKind(c)))
			}
		}
	}
	for _, action := range parseActions(parts[1]) {
		move(entities, robot, action, false)
	}
	return gps(entities)
}

func parseActions(input string) []Action {
	var actions []Action
	for _, r := range input {
		actions = append(actions, Action(r))
	}
	return actions
}

func gps(entities []*Entity) int {
	total := 0
	for _, entity := range entities {
		if entity.Kind != BOX {
			continue
		}
		total += 100*entity.Y + entity.X
	}
	return total
}

var id int

func newEntity(x, y, width, height int, kind EntityKind) *Entity {
	id++
	return &Entity{ID: id, X: x, Y: y, Width: width, Height: height, Kind: kind}
}

func move(entities []*Entity, entity *Entity, action Action, dryRun bool) bool {
	var nextX, nextY int
	switch action {
	case UP:
		nextX, nextY = entity.X, entity.Y-1
	case DOWN:
		nextX, nextY = entity.X, entity.Y+1
	case LEFT:
		nextX, nextY = entity.X-1, entity.Y
	case RIGHT:
		nextX, nextY = entity.X+1, entity.Y
	}
	var pushed []*Entity
	for _, e := range entities {
		if e.ID == entity.ID {
			continue
		}
		if !e.Contains(nextX, nextY, entity.Width, entity.Height) {
			continue
		}
		if e.Kind == WALL {
			return false
		}
		if e.Kind == BOX {
			if !move(entities, e, action, true) {
				return false
			}
			pushed = append(pushed, e)
		}
	}
	if !dryRun {
		for _, e := range pushed {
			move(entities, e, action, false)
		}
		entity.X, entity.Y = nextX, nextY
	}
	return true
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/15/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}
