package _09

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

type Space struct {
	FileID int
	Len    int
}

func (s Space) String() string {
	b := strings.Builder{}
	for i := 0; i < s.Len; i++ {
		if s.FileID == -1 {
			b.WriteByte('.')
		} else {
			b.WriteString(strconv.Itoa(s.FileID))
		}
	}
	return b.String()
}

func part2(input string) int {
	fileID := 0
	var disk []Space
	for i, r := range input {
		n, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}

		if i%2 == 0 {
			// file
			disk = append(disk, Space{
				FileID: fileID,
				Len:    n,
			})
			fileID++
		} else {
			// empty block
			disk = append(disk, Space{
				FileID: -1,
				Len:    n,
			})
		}
	}

	// compact
	for {
		changed := false
		for j := len(disk) - 1; ; j-- {
			if j < 0 {
				break
			}
			if disk[j].FileID == -1 {
				continue
			}
			i := 0
			for {
				if i >= j {
					break
				}
				if disk[i].FileID == -1 && disk[i].Len >= disk[j].Len {
					break
				}
				i++
			}
			if i >= j {
				continue
			}
			changed = true
			delta := disk[i].Len - disk[j].Len
			disk[i], disk[j] = disk[j], Space{
				FileID: -1,
				Len:    disk[j].Len,
			}
			if delta > 0 {
				disk = slices.Insert(disk, i+1, Space{FileID: -1, Len: delta})
			}
		}
		if !changed {
			break
		}
	}

	// compute checksum
	checksum := 0
	i := 0
	for _, v := range disk {
		for range v.Len {
			if v.FileID != -1 {
				checksum += i * v.FileID
			}
			i++
		}
	}
	return checksum
}

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/09/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}
