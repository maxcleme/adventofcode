// Let’s be honest: I had no idea how to solve this in pure Go without pulling in
// an SMT solver, just like pretty much everyone else for this day/part.
//
// While poking around for a way to do it without external deps, I ran into this
// Reddit post:
// https://www.reddit.com/r/adventofcode/comments/1pk87hl/2025_day_10_part_2_bifurcate_your_way_to_victory/
//
// All the credit goes to the original author. I (kiro) just translated the idea
// to Go and got it working.
//
// Honestly, it’s a fantastic idea. I really wish I’d thought of it first.

package _10

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2025/10/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}

func part2(payload string) int {
	var machines []Machine
	for _, line := range strings.Split(payload, "\n") {
		fields := strings.Fields(line)
		m := Machine{}
		for _, inputs := range fields[1 : len(fields)-1] {
			var buttons []int
			for _, input := range strings.Split(inputs[1:len(inputs)-1], ",") {
				v, _ := strconv.Atoi(input)
				buttons = append(buttons, v)
			}
			m.Buttons = append(m.Buttons, buttons)
		}
		for _, joltage := range fields[len(fields)-1:] {
			for _, j := range strings.Split(joltage[1:len(joltage)-1], ",") {
				n, _ := strconv.Atoi(j)
				m.Joltage = append(m.Joltage, n)
			}
		}
		machines = append(machines, m)
	}

	total := 0
	for i, m := range machines {
		result := solve(m.Joltage, m.Buttons, map[string]int{})
		if result == -1 {
			panic("no solution for machine " + strconv.Itoa(i))
		}
		total += result
	}
	return total
}

func solve(target []int, buttons [][]int, cache map[string]int) int {
	start := !slices.ContainsFunc(target, func(i int) bool {
		return i != 0
	})
	if start {
		return 0
	}

	key := fmt.Sprint(target)
	if val, exists := cache[key]; exists {
		return val
	}

	// Try all 2^n combinations of buttons
	minCost := -1
	n := len(buttons)
	for mask := 0; mask < (1 << n); mask++ {
		// Calculate what we need after pressing this combination
		buttonCount := 0
		remaining := slices.Clone(target)
		for i := range n {
			if (mask & (1 << i)) != 0 {
				buttonCount++
				// Subtract the effect of pressing button i once
				for _, pos := range buttons[i] {
					if pos < len(remaining) {
						remaining[pos]--
					}
				}
			}
		}
		// Stop exploring this path as solution is invalid
		if slices.ContainsFunc(remaining, func(i int) bool { return i < 0 }) {
			continue
		}
		// Or odd numbers remain, can't divide by 2
		if slices.ContainsFunc(remaining, func(i int) bool { return i%2 != 0 }) {
			continue
		}
		// Divide by 2 and recurse
		halved := make([]int, len(remaining))
		for i, r := range remaining {
			halved[i] = r / 2
		}
		subCost := solve(halved, buttons, cache)
		if subCost == -1 {
			continue
		}
		totalCost := 2*subCost + buttonCount
		if minCost == -1 || totalCost < minCost {
			minCost = totalCost
		}
	}
	cache[key] = minCost
	return minCost
}

func solveWithConstraint(target []int, buttons [][]int, requireNonNegative bool) []*big.Int {
	var a [][]*big.Int
	var b []*big.Int
	fmt.Println("Target:", target)
	fmt.Println("Buttons:", buttons)

	// Build constraint matrix: each row represents one position in the target
	for i, t := range target {
		var row []*big.Int
		for range buttons {
			row = append(row, big.NewInt(0))
		}
		// For each button, check if it affects position i
		for j, buttonPositions := range buttons {
			if slices.Contains(buttonPositions, i) {
				row[j] = big.NewInt(1)
			}
		}
		a = append(a, row)
		b = append(b, big.NewInt(int64(t)))
	}

	fmt.Println("Matrix A:")
	for i, row := range a {
		fmt.Printf("Position %d (target=%d): ", i, target[i])
		for _, val := range row {
			fmt.Printf("%d ", val.Int64())
		}
		fmt.Println()
	}

	results, err := gaussJordanFindIntegerSolution(a, b, requireNonNegative)
	if err != nil {
		fmt.Println("No solution found:", err)
		// Return zero solution for unsolvable machines
		zero := make([]*big.Int, len(buttons))
		for i := range zero {
			zero[i] = big.NewInt(0)
		}
		return zero
	}
	fmt.Println("Solution:", results)
	return results
}

func lcmBig(a, b *big.Int) *big.Int {
	if a.Sign() == 0 {
		return new(big.Int).Abs(b)
	}
	if b.Sign() == 0 {
		return new(big.Int).Abs(a)
	}
	g := new(big.Int).GCD(nil, nil, a, b)
	absA := new(big.Int).Abs(a)
	absB := new(big.Int).Abs(b)
	tmp := new(big.Int).Div(new(big.Int).Mul(absA, absB), g)
	return tmp
}

func ratDen(r *big.Rat) *big.Int {
	return new(big.Int).Set(r.Denom())
}

func ratToIntMod(r *big.Rat, mod *big.Int) (*big.Int, error) {
	// Return (r mod mod) represented as integer in [0, mod-1], assuming mod > 0.
	// r = num/den. Need inverse of den mod mod (so gcd(den,mod)=1) for this to work.
	if mod.Sign() <= 0 {
		return nil, errors.New("mod must be positive")
	}
	num := new(big.Int).Mod(new(big.Int).Set(r.Num()), mod)
	den := new(big.Int).Mod(new(big.Int).Set(r.Denom()), mod)

	g := new(big.Int).GCD(nil, nil, den, mod)
	if g.Cmp(big.NewInt(1)) != 0 {
		return nil, errors.New("denominator not invertible modulo mod")
	}

	inv := new(big.Int).ModInverse(den, mod)
	if inv == nil {
		return nil, errors.New("no modular inverse")
	}
	return new(big.Int).Mod(new(big.Int).Mul(num, inv), mod), nil
}

func gaussJordanFindIntegerSolution(A [][]*big.Int, b []*big.Int, requireNonNegative bool) ([]*big.Int, error) {
	m := len(A)
	if m == 0 {
		return nil, errors.New("empty A")
	}
	n := len(A[0])
	if len(b) != m {
		return nil, errors.New("dimension mismatch: len(b) must equal rows(A)")
	}
	for i := range A {
		if len(A[i]) != n {
			return nil, errors.New("ragged A")
		}
	}

	// Build augmented matrix M (m x (n+1)) as *big.Rat
	M := make([][]*big.Rat, m)
	for i := 0; i < m; i++ {
		M[i] = make([]*big.Rat, n+1)
		for j := 0; j < n; j++ {
			M[i][j] = new(big.Rat).SetInt(A[i][j])
		}
		M[i][n] = new(big.Rat).SetInt(b[i])
	}

	// RREF
	pivotColOfRow := make([]int, 0, m)
	row := 0
	for col := 0; col < n && row < m; col++ {
		piv := -1
		for r := row; r < m; r++ {
			if M[r][col].Sign() != 0 {
				piv = r
				break
			}
		}
		if piv == -1 {
			continue
		}
		M[row], M[piv] = M[piv], M[row]

		pivot := new(big.Rat).Set(M[row][col])
		for c := col; c <= n; c++ {
			M[row][c].Quo(M[row][c], pivot)
		}
		for r := 0; r < m; r++ {
			if r == row {
				continue
			}
			f := new(big.Rat).Set(M[r][col])
			if f.Sign() == 0 {
				continue
			}
			for c := col; c <= n; c++ {
				tmp := new(big.Rat).Mul(f, M[row][c])
				M[r][c].Sub(M[r][c], tmp)
			}
		}

		pivotColOfRow = append(pivotColOfRow, col)
		row++
	}

	// Inconsistency check: [0..0 | nonzero]
	for r := 0; r < m; r++ {
		allZero := true
		for c := 0; c < n; c++ {
			if M[r][c].Sign() != 0 {
				allZero = false
				break
			}
		}
		if allZero && M[r][n].Sign() != 0 {
			return nil, errors.New("no solution")
		}
	}

	// Identify free columns
	isPivot := make([]bool, n)
	for _, pc := range pivotColOfRow {
		isPivot[pc] = true
	}
	freeCols := make([]int, 0, n)
	for c := 0; c < n; c++ {
		if !isPivot[c] {
			freeCols = append(freeCols, c)
		}
	}

	// If there are no free vars, just read solution and enforce integrality
	if len(freeCols) == 0 {
		x := make([]*big.Int, n)
		for i := range x {
			x[i] = new(big.Int)
		}
		for r, pc := range pivotColOfRow {
			if !M[r][n].IsInt() {
				return nil, errors.New("unique solution is not integer")
			}
			x[pc].Set(M[r][n].Num())
		}
		return x, nil
	}

	// Build a modulus L = lcm of all denominators that appear in pivot equations.
	L := big.NewInt(1)
	for r := 0; r < len(pivotColOfRow); r++ {
		L = lcmBig(L, ratDen(M[r][n]))
		for _, fc := range freeCols {
			L = lcmBig(L, ratDen(M[r][fc]))
		}
	}

	if L.BitLen() > 30 {
		return nil, errors.New("denominator LCM too large to brute-force")
	}

	Lint := int(L.Int64())
	f := len(freeCols)
	freeVals := make([]int, f)

	var bestSol []*big.Int
	bestSum := big.NewInt(-1)

	next := func() bool {
		for i := 0; i < f; i++ {
			freeVals[i]++
			if freeVals[i] < Lint {
				return true
			}
			freeVals[i] = 0
		}
		return false
	}

	tryAssignment := func() ([]*big.Int, bool) {
		xRat := make([]*big.Rat, n)
		for i := 0; i < n; i++ {
			xRat[i] = new(big.Rat)
		}

		for i, fc := range freeCols {
			xRat[fc].SetInt64(int64(freeVals[i]))
		}

		for r, pc := range pivotColOfRow {
			val := new(big.Rat).Set(M[r][n])
			for _, fc := range freeCols {
				if M[r][fc].Sign() == 0 {
					continue
				}
				tmp := new(big.Rat).Mul(M[r][fc], xRat[fc])
				val.Sub(val, tmp)
			}
			xRat[pc] = val
			if !xRat[pc].IsInt() {
				return nil, false
			}
		}

		// Minimize sum of absolute values
		out := make([]*big.Int, n)
		absSum := big.NewInt(0)
		for i := 0; i < n; i++ {
			if !xRat[i].IsInt() {
				return nil, false
			}
			out[i] = new(big.Int).Set(xRat[i].Num())
			absSum.Add(absSum, new(big.Int).Abs(out[i]))
		}

		if bestSol == nil || absSum.Cmp(bestSum) < 0 {
			bestSum.Set(absSum)
			bestSol = make([]*big.Int, n)
			for i := 0; i < n; i++ {
				bestSol[i] = new(big.Int).Set(out[i])
			}
		}
		return out, true
	}

	// Enumerate all possibilities to find minimum
	for {
		tryAssignment()
		if !next() {
			break
		}
	}

	if bestSol == nil {
		return nil, errors.New("no integer solution found")
	}
	return bestSol, nil
}
