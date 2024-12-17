package _17

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/maxcleme/adventofcode/2024/17/parser"
	"github.com/spf13/cobra"
)

func part2(input string) int {
	lexer := parser.NewAOCLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewAOCParser(stream)
	listener := &ProgramListener{
		registries: map[string]int{},
	}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Program())
	program := strings.ReplaceAll(listener.program, ",", "")
	checkpoint := 0
	for n := len(program) - 1; n >= 0; n-- {
		target := program[n-1:]
		candidate := checkpoint * 8
		for {
			listener.registries["A"], listener.registries["B"], listener.registries["C"] = candidate, 0, 0
			result := run(listener.registries, listener.instructions)
			if result == program {
				return candidate
			}
			if result == target {
				checkpoint = candidate
				break
			}
			candidate++
		}
	}
	return -1
}

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/17/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}
