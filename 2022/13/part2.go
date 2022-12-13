package _13

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open("./2022/13/input")
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(f)
		packets := []Packet{
			PacketDistressStart,
			PacketDistressEnd,
		}
		for scanner.Scan() {
			l := scanner.Text()
			if len(l) == 0 {
				continue
			}
			packets = append(packets, MustParsePacket(l))
		}
		sort.Slice(packets, func(i, j int) bool {
			return IsOrdered(packets[i], packets[j])
		})
		var start, end int
		for i, p := range packets {
			if reflect.DeepEqual(p, PacketDistressStart) {
				start = i + 1
			}
			if reflect.DeepEqual(p, PacketDistressEnd) {
				end = i + 1
			}
		}
		fmt.Println(start * end)
		return nil
	},
}
