package day6

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// input represents the input data as a grid of integers.
type input struct {
	Grid [][]int
	ops  []Operator
}

func ParseIn(r io.Reader) (*Worksheet, error) {
	scanner := bufio.NewScanner(r)
	var grid [][]int
	var ops []Operator

	for scanner.Scan() {
		line := scanner.Text()
		// Skip empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Check if this is the operator row (contains * or +)
		if strings.Contains(line, "*") || strings.Contains(line, "+") {
			stringToOperator := map[string]Operator{"*": Multiple, "+": Add}
			for field := range strings.FieldsSeq(line) {
				ops = append(ops, stringToOperator[field])
			}
			break
		}

		// Parse the numbers from this row
		row, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		if len(row) > 0 {
			grid = append(grid, row)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed reading input file: %w", err)
	}

	return inputToProblems(input{Grid: grid, ops: ops}), nil
}

func parseLine(line string) ([]int, error) {
	fields := strings.Fields(line)
	row := make([]int, 0, len(fields))
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			return nil, fmt.Errorf("failed processing input file numbers: %w", err)
		}
		row = append(row, num)
	}
	return row, nil
}

func inputToProblems(in input) *Worksheet {
	problems := []Problem{}
	for colIdx := range len(in.Grid[0]) {
		// go down the column and collect the operands
		operands := []int{}
		for rowIdx := range len(in.Grid) {
			operands = append(operands, in.Grid[rowIdx][colIdx])
		}

		problems = append(problems, Problem{
			operands: operands,
			operator: in.ops[colIdx],
		})
	}
	return &Worksheet{problems: problems}
}

func ParseInPart2(r io.Reader) (*Worksheet, error) {
	scanner := bufio.NewScanner(r)
	var grid [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		if len(row) > 0 {
			grid = append(grid, row)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed reading input file: %w", err)
	}
	return runeToWorkshet(grid), nil
}

func runeToWorkshet(grid [][]rune) *Worksheet {
	width := len(grid[0])
	height := len(grid)
	problems := []Problem{}

	runeToOperator := map[rune]Operator{'*': Multiple, '+': Add}

	operands := []int{}
	for colIdx := width - 1; colIdx >= 0; colIdx-- {
		columnBlanks := 0
		columnDigits := []string{}
		// collect digits from the top rows
		for rowIdx := range height - 1 {
			value := grid[rowIdx][colIdx]
			if value == ' ' {
				columnBlanks++
			} else {
				columnDigits = append(columnDigits, string(value))
			}
		}
		// combine digits into an integer
		if len(columnDigits) > 0 {
			operand, err := strconv.Atoi(strings.Join(columnDigits, ""))
			if err != nil {
				panic(err)
			}
			operands = append(operands, operand)
		}

		// If you hit blank column record the problem
		if columnBlanks == height-1 {
			problems = append(problems, Problem{
				operands: operands,
				operator: runeToOperator[grid[height-1][colIdx+1]],
			})
			operands = []int{}
		} else if colIdx == 0 {
			// If you hit the end record the problem
			problems = append(problems, Problem{
				operands: operands,
				operator: runeToOperator[grid[height-1][colIdx]],
			})
			operands = []int{}
		}
	}

	return &Worksheet{problems}
}
