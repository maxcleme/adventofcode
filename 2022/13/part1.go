package _13

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open("./2022/13/input")
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(f)
		var pairs [][2]Packet
		var pair [2]Packet
		for scanner.Scan() {
			l := scanner.Text()
			if len(l) == 0 {
				pairs = append(pairs, pair)
				pair = [2]Packet{}
				continue
			}
			if pair[0] == nil {
				pair[0] = MustParsePacket(l)
			} else {
				pair[1] = MustParsePacket(l)
			}
		}
		if pair[0] != nil && pair[1] != nil {
			pairs = append(pairs, pair)
		}
		var count int
		for i, p := range pairs {
			if IsOrdered(p[0], p[1]) {
				count += i + 1
			}
		}
		fmt.Println(count)
		return nil
	},
}
