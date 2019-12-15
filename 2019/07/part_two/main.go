package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/maxcleme/adventofcode/2019/07/part_two/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	// read input file
	file, err := os.Open("./input")
	if err != nil {
		logrus.WithError(err).Fatal("cannot read file")
	}
	defer file.Close()

	// read input
	scanner := bufio.NewScanner(file)
	// Define a split function that separates on commas.
	// According to https://golang.org/pkg/bufio/#Scanner examples
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)

	// scan values separated by coma
	statements := make([]int, 0)
	for scanner.Scan() {
		statement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			logrus.WithError(err).Fatal("cannot parse statement")
		}
		statements = append(statements, statement)
	}

	phases := permutations([]int{5, 6, 7, 8, 9})

	pwg := sync.WaitGroup{}
	pwg.Add(len(phases))

	for _, permutation := range phases {
		go func(permutation []int) {
			defer pwg.Done()

			phase, init := permutation[0], 0
			feedBackReader, feedbackWriter := io.Pipe()

			seqInput := io.MultiReader(bufio.NewReader(strings.NewReader(fmt.Sprintf("%d\n%d\n", phase, init))), feedBackReader)
			seqOutput := io.MultiWriter(os.Stdout, feedbackWriter)
			seqLength := 5

			programs := make([]utils.Program, seqLength)
			var pReader *os.File
			var pWriter *os.File
			for i := 0; i < len(programs); i++ {
				p := utils.Program{
					Statements: statements,
				}
				if i == 0 {
					p.In = seqInput
				} else {
					p.In = pReader
				}

				if i == len(programs)-1 {
					p.Out = seqOutput
				} else {
					pReader, pWriter, err = os.Pipe()
					if err != nil {
						logrus.
							WithField("program_id", i).
							WithError(err).
							Fatal("cannot create pipe")
					}
					p.Out = pWriter
					p.Out.Write([]byte(fmt.Sprintf("%d\n", permutation[i+1])))
				}
				programs[i] = p
			}

			wg := sync.WaitGroup{}
			wg.Add(seqLength)

			for i := 0; i < len(programs); i++ {
				go func(i int) {
					defer wg.Done()
					// find program in sequence
					p := programs[i]

					// start program
					if _, err = utils.Run(p); err != nil {
						logrus.
							WithField("program_id", i).
							WithError(err).
							Fatal("cannot run program")
					}

					if i == 0 {
						// discard feedback once first program exit
						io.Copy(ioutil.Discard, feedBackReader)
					}

					if i == len(programs)-1 {
						// close feedback loop once last program exit
						feedBackReader.Close()
						feedbackWriter.Close()
					}

				}(i)
			}
			wg.Wait()
		}(permutation)
	}

	pwg.Wait()
}

func permutations(values []int) [][]int {
	return _permutations(values, make([][]int, 0), 0, len(values))
}

func _permutations(values []int, buff [][]int, l, r int) [][]int {
	if l == r {
		buff = append(buff, values)
		return buff
	}
	for i := 0; i < r; i++ {
		buff = _permutations(swap(values, l, i), buff, l+1, r)
	}
	return buff
}

func swap(values []int, i, j int) []int {
	clone := make([]int, len(values))
	for i, v := range values {
		clone[i] = v
	}
	clone[i], clone[j] = clone[j], clone[i]
	return clone
}
