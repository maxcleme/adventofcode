package _22

import (
	"fmt"
	"maps"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type order struct {
	A, B, C, D int
}

func part2(input string) int {
	best := 0
	var secrets []int
	for _, row := range strings.Split(input, "\n") {
		secret, err := strconv.Atoi(row)
		if err != nil {
			panic(err)
		}
		secrets = append(secrets, secret)
	}
	changes := map[order]struct{}{}
	history := map[int]map[order]int{}
	for _, secret := range secrets {
		rules := all(secret)
		history[secret] = rules
		for k := range rules {
			changes[k] = struct{}{}
		}
	}
	for rule := range maps.Keys(changes) {
		total := 0
		for _, secret := range secrets {
			total += history[secret][rule]
		}
		best = max(best, total)
	}
	return best
}

func all(secret int) map[order]int {
	result := map[order]int{}
	var changes []int
	for range 2000 {
		currentPrice := price(secret)
		secret = next(secret)
		nextPrice := price(secret)
		diff := nextPrice - currentPrice
		changes = append(changes, diff)
		if len(changes) > 4 {
			changes = changes[1:]
		}
		if len(changes) != 4 {
			continue
		}
		r := order{changes[0], changes[1], changes[2], changes[3]}
		if _, ok := result[r]; ok {
			continue
		}
		result[r] = nextPrice
	}
	return result
}

func price(next int) int {
	s := strconv.Itoa(next)
	n, err := strconv.Atoi(s[len(s)-1:])
	if err != nil {
		panic(err)
	}
	return n
}

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/22/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}
