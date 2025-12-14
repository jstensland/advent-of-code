// Package day4 solves AoC 2025 day 4
package day4

import (
	"bufio"
	"fmt"
	"io"
)

// CellState represents the state of a location in the grid.
type CellState int

const (
	Empty CellState = iota
	PaperRoll
)

// Grid represents the parsed input grid where each location can either be a paper roll or empty.
type Grid struct {
	Cells  [][]CellState
	Width  int
	Height int
}

func Part1(r io.Reader) (int, error) {
	answer := 0
	_, err := ParseIn(r)
	if err != nil {
		return 0, err
	}
	// TODO: solve part 1

	return answer, nil
}

func Part2(r io.Reader) (int, error) {
	answer := 0
	_, err := ParseIn(r)
	if err != nil {
		return 0, err
	}
	// TODO: solve part 2

	return answer, nil
}

func ParseIn(r io.Reader) (Grid, error) {
	scanner := bufio.NewScanner(r)
	var cells [][]CellState

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		row := make([]CellState, len(line))
		for i, char := range line {
			switch char {
			case '@':
				row[i] = PaperRoll
			case '.':
				row[i] = Empty
			}
		}
		cells = append(cells, row)
	}

	if err := scanner.Err(); err != nil {
		return Grid{}, fmt.Errorf("scanner error: %w", err)
	}

	height := len(cells)
	width := 0
	if height > 0 {
		width = len(cells[0])
	}

	return Grid{
		Cells:  cells,
		Width:  width,
		Height: height,
	}, nil
}
