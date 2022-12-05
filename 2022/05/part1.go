package _05

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open("./2022/05/input")
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(f)
		cratesRgx := regexp.MustCompile(`\s?(\s{3}|\[[A-Z]])`)
		moveRgx := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
		var crates []crate
		for scanner.Scan() {
			l := scanner.Text()
			if strings.TrimSpace(l) == "" {
				continue
			}
			if group := cratesRgx.FindAllString(l, -1); len(group) != 0 {
				if len(crates) < len(group) {
					for i := len(crates); i < len(group); i++ {
						crates = append(crates, MakeCrate())
					}
				}
				for i, c := range group {
					crates[i].Add(c)
				}
			}
			if group := moveRgx.FindAllStringSubmatch(l, -1); len(group) != 0 {
				count, _ := strconv.Atoi(group[0][1])
				from, _ := strconv.Atoi(group[0][2])
				to, _ := strconv.Atoi(group[0][3])
				crates = Move9000(crates, count, from, to)
			}
		}
		var s string
		for _, c := range crates {
			s += string(c.Pop())
		}
		fmt.Println(s)
		return nil
	},
}
