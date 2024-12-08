// Package day7 solves AoC 2024 day 7.
package day7

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"strconv"
	"strings"
)

func SolvePart1(in io.Reader) (int, error) {
	eqs, err := ParseInput(in)
	if err != nil {
		return 0, fmt.Errorf("error parsing input file: %w", err)
	}

	total := 0
	for _, eq := range eqs {
		if eq.IsPossible([]BinaryOp{Add, Multiple}) {
			total += eq.answer
		}
	}

	return total, nil
}

func SolvePart2(in io.Reader) (int, error) {
	eqs, err := ParseInput(in)
	if err != nil {
		return 0, fmt.Errorf("error parsing input file: %w", err)
	}

	total := 0
	for _, eq := range eqs {
		if eq.IsPossible([]BinaryOp{Add, Multiple, Concat}) {
			total += eq.answer
		}
	}

	return total, nil
}

// Equation first value is the answer, and the rest
type Equation struct {
	answer   int
	operands []int
}

// IsPossible returns true if there is a + and * combo to insert  that equals the answer
func (eq Equation) IsPossible(ops []BinaryOp) bool {
	// for _, opPerm := range Perms(len(eq.operands)-1, ops) {
	for opPerm := range Perms2(len(eq.operands)-1, ops) {
		// fmt.Println("operands:", eq.operands)
		// fmt.Println("ops:", opPerm)
		// fmt.Println("computed", compute(eq.operands, opPerm))
		if compute(eq.operands, opPerm) == eq.answer {
			// fmt.Println("worked!")
			return true
		}
	}
	return false
}

// BinaryOp is a function that takes two integers and returns an integer
// defined for readability
type BinaryOp func(int, int) int

func Add(left, right int) int      { return left + right }
func Multiple(left, right int) int { return left * right }
func Concat(left, right int) int {
	res, err := strconv.Atoi(strconv.Itoa(left) + strconv.Itoa(right))
	if err != nil {
		// impossible since we start with ints
		panic(err)
	}
	return res
}

// String is a hack for debugging. It only works because of my limited number
// of operators where I can pick op results that are not collisions.
// Each must be added here. It is not needed for anything other than printing
// and op function comparison.
//
// IMPROVEMENT: refactor my BinaryOp to a struct type with a Name field for comparisons
func (op BinaryOp) String() string {
	if op(1, 1) == 2 { //nolint:mnd // a little magic
		return "+"
	}
	if op(1, 1) == 1 {
		return "*"
	}
	if op(1, 1) == 11 { //nolint:mnd // a little magic
		return "||"
	}

	return "unknown"
}

// Perms2 is a recursive iterator implementation of all the permutations
//
// This approach is twice as fast as Perms for our input, and less memory intensive
func Perms2(size int, set []BinaryOp) iter.Seq[[]BinaryOp] {
	return func(yield func(op []BinaryOp) bool) {
		// base case
		if size == 1 {
			for _, op := range set {
				if !yield([]BinaryOp{op}) {
					return
				}
			}
			return
		}

		for _, op := range set {
			for perm := range Perms2(size-1, set) { // [*] [+]
				if !yield(append(perm, op)) {
					return
				}
			}
		}
	}
}

// Perms is a recursive implementation of all the permutations
func Perms(size int, set []BinaryOp) [][]BinaryOp {
	if size == 1 {
		opPerSet := [][]BinaryOp{}
		for _, op := range set {
			opPerSet = append(opPerSet, []BinaryOp{op})
		}
		return opPerSet
	}
	perms := [][]BinaryOp{}
	for _, op := range set {
		for _, perm := range Perms(size-1, set) { // [*] [+]
			perm = append(perm, op)     // add your op to the perm...
			perms = append(perms, perm) // add this perm to total set
		}
	}
	return perms
}

func compute(operands []int, ops []BinaryOp) int {
	result := operands[0]
	for i, op := range ops {
		result = op(result, operands[i+1])
	}
	return result
}

func ParseInput(in io.Reader) ([]Equation, error) {
	equations := []Equation{}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		res := strings.Split(line, ": ")
		answer, err := strconv.Atoi(res[0])
		if err != nil {
			return nil, fmt.Errorf("error parsing answer: %w", err)
		}
		operands := make([]int, len(strings.Fields(res[1])))
		rawOps := strings.Fields(res[1])
		for i, op := range rawOps {
			num, err := strconv.Atoi(op)
			if err != nil {
				return nil, fmt.Errorf("bad operand: %w", err)
			}
			operands[i] = num
		}

		equations = append(equations, Equation{answer: answer, operands: operands})
	}

	return equations, nil
}
