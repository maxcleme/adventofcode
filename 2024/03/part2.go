package _03

import (
	"fmt"
	"os"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/maxcleme/adventofcode/2024/03/parser"
	"github.com/spf13/cobra"
)

type AOCListener struct {
	*parser.BaseAOCListener
	enabled bool
	total   int
}

func (l *AOCListener) EnterDo(_ *parser.DoContext) {
	l.enabled = true
}

func (l *AOCListener) EnterDon_t(_ *parser.Don_tContext) {
	l.enabled = false
}

func (l *AOCListener) EnterMul(ctx *parser.MulContext) {
	if ctx.NUM(0) == nil || ctx.NUM(1) == nil {
		return
	}

	x, err1 := strconv.Atoi(ctx.NUM(0).GetText())
	y, err2 := strconv.Atoi(ctx.NUM(1).GetText())

	if err1 != nil || err2 != nil {
		return
	}

	// Check if the entire expression is correctly formatted
	expectedText := fmt.Sprintf("mul(%d,%d)", x, y)
	if ctx.GetText() != expectedText {
		return
	}
	fmt.Println(ctx.GetText())
	if l.enabled {
		l.total += x * y
	}
}

func sumPart2(program string) int {
	lexer := parser.NewAOCLexer(antlr.NewInputStream(program))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewAOCParser(stream)
	listener := &AOCListener{
		enabled: true,
	}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Program())
	return listener.total
}

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		input, err := os.ReadFile("./2024/03/input")
		if err != nil {
			return err
		}
		fmt.Printf("Total: %d\n", sumPart2(string(input)))
		return nil
	},
}
