package _17

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/maxcleme/adventofcode/2024/17/parser"
	"github.com/spf13/cobra"
)

type instruction func(int, map[string]int, io.Writer) int

type ProgramListener struct {
	*parser.BaseAOCListener

	program      string
	registries   map[string]int
	instructions []instruction
}

func (s *ProgramListener) EnterInstructions(ctx *parser.InstructionsContext) {
	var instructions []string
	for _, i := range ctx.AllInstruction() {
		instructions = append(instructions, i.GetText())
	}
	s.program = strings.Join(instructions, ",")
}

func (s *ProgramListener) EnterRegister(ctx *parser.RegisterContext) {
	fmt.Println("register", ctx.ID(), ctx.INTEGER())
	s.registries[ctx.ID().GetText()] = 0
	if ctx.INTEGER() == nil {
		return
	}
	n, err := strconv.Atoi(ctx.INTEGER().GetText())
	if err != nil {
		panic(err)
	}
	s.registries[ctx.ID().GetText()] = n
}

func (s *ProgramListener) EnterAdv(ctx *parser.AdvContext) {
	s.instructions = append(s.instructions, func(ip int, registries map[string]int, stdout io.Writer) int {
		registries["A"] = int(float64(registries["A"]) / math.Pow(2, float64(s.evaluate(ctx.Operand())(registries))))
		return ip + 1
	})
}

func (s *ProgramListener) EnterBxl(ctx *parser.BxlContext) {
	s.instructions = append(s.instructions, func(ip int, registries map[string]int, stdout io.Writer) int {
		registries["B"] = registries["B"] ^ s.evaluate(ctx.Operand())(registries)
		return ip + 1
	})
}

func (s *ProgramListener) EnterBst(ctx *parser.BstContext) {
	s.instructions = append(s.instructions, func(ip int, registries map[string]int, stdout io.Writer) int {
		registries["B"] = s.evaluate(ctx.Operand())(registries) % 8
		return ip + 1
	})
}

func (s *ProgramListener) EnterJnz(ctx *parser.JnzContext) {
	s.instructions = append(s.instructions, func(ip int, registries map[string]int, stdout io.Writer) int {
		if registries["A"] != 0 {
			return s.evaluate(ctx.Operand())(registries)
		}
		return ip + 1
	})
}

func (s *ProgramListener) EnterBxc(_ *parser.BxcContext) {
	s.instructions = append(s.instructions, func(ip int, registries map[string]int, stdout io.Writer) int {
		registries["B"] = registries["B"] ^ registries["C"]
		return ip + 1
	})
}

func (s *ProgramListener) EnterOut(ctx *parser.OutContext) {
	s.instructions = append(s.instructions, func(ip int, registries map[string]int, stdout io.Writer) int {
		_, err := fmt.Fprint(stdout, s.evaluate(ctx.Operand())(registries)%8)
		if err != nil {
			panic(err)
		}
		return ip + 1
	})
}

func (s *ProgramListener) EnterBdv(ctx *parser.BdvContext) {
	s.instructions = append(s.instructions, func(ip int, registries map[string]int, stdout io.Writer) int {
		registries["B"] = int(float64(registries["A"]) / math.Pow(2, float64(s.evaluate(ctx.Operand())(registries))))
		return ip + 1
	})
}

func (s *ProgramListener) EnterCdv(ctx *parser.CdvContext) {
	s.instructions = append(s.instructions, func(ip int, registries map[string]int, stdout io.Writer) int {
		registries["C"] = int(float64(registries["A"]) / math.Pow(2, float64(s.evaluate(ctx.Operand())(registries))))
		return ip + 1
	})
}

func (s *ProgramListener) evaluate(operand parser.IOperandContext) func(map[string]int) int {
	l := &OperandListener{}
	antlr.ParseTreeWalkerDefault.Walk(l, operand)
	return l.read
}

type OperandListener struct {
	*parser.BaseAOCListener
	read func(map[string]int) int
}

func (s *OperandListener) EnterLiteral_value(ctx *parser.Literal_valueContext) {
	n, _ := strconv.Atoi(ctx.GetText())
	s.read = func(map[string]int) int {
		return n
	}
}

func (s *OperandListener) EnterReg_a(_ *parser.Reg_aContext) {
	s.read = func(registries map[string]int) int {
		return registries["A"]
	}
}
func (s *OperandListener) EnterReg_b(_ *parser.Reg_bContext) {
	s.read = func(registries map[string]int) int {
		return registries["B"]
	}
}
func (s *OperandListener) EnterReg_c(_ *parser.Reg_cContext) {
	s.read = func(registries map[string]int) int {
		return registries["C"]
	}
}

func part1(input string) string {
	lexer := parser.NewAOCLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewAOCParser(stream)

	listener := &ProgramListener{
		registries: map[string]int{},
	}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Program())
	return strings.Join(strings.Split(run(listener.registries, listener.instructions), ""), ",")
}

func run(registries map[string]int, instructions []instruction) string {
	var buf bytes.Buffer
	ip := 0
	for ip < len(instructions) {
		ip = instructions[ip](ip, registries, &buf)
	}
	return buf.String()
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/17/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}
