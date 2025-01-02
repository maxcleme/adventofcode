package _01

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2015/04/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}

func part2(input string) int {
	prefix := []byte{'\x00', '\x00', '\x00'}
	result := make(chan int)
	for i := range runtime.NumCPU() {
		go func() {
			for {
				m := md5.New()
				m.Reset()
				_, err := fmt.Fprintf(m, "%s%d", input, i)
				if err != nil {
					panic(err)
				}
				if bytes.HasPrefix(m.Sum(nil), prefix) {
					result <- i
					return
				}
				i += runtime.NumCPU()
			}
		}()
	}
	return <-result
}
