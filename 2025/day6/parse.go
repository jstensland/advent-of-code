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

	return InputToProblems(input{Grid: grid, ops: ops}), nil
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

func InputToProblems(in input) *Worksheet {
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
