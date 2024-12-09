package _09

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func part1(input string) int {
	fileID := 0
	var disk []int
	for i, r := range input {
		n, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}

		for range n {
			if i%2 == 0 {
				// file
				disk = append(disk, fileID)
			} else {
				// empty block
				disk = append(disk, -1)
			}
		}
		if i%2 == 0 {
			fileID++
		}
	}

	// compact
	for i := 0; ; i++ {
		if i >= len(disk) {
			break
		}
		if disk[i] != -1 {
			continue
		}
		j := len(disk) - 1
		for {
			if disk[j] != -1 {
				break
			}
			j--
		}
		disk[i] = disk[j]
		disk = disk[:j]
	}

	// compute checksum
	checksum := 0
	for i, v := range disk {
		if v == -1 {
			continue
		}
		checksum += i * v
	}
	return checksum
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/09/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}
