// Package day4 solves AoC 2025 day 4
package day4

import (
	"bufio"
	"fmt"
	"io"
)

func Part1(r io.Reader) (int, error) {
	grid, err := ParseIn(r)
	if err != nil {
		return 0, err
	}

	// Option 1
	// - create a grid of counts to match the input grid. or just add an attribut of adjacent counts
	// - iterate over the input grid, incrementing the adjacent cell's counts for each paper found
	// - look through the "counts grid" for how many paper locations have few than 4 adjacent cells

	// Option 2
	// - for each cell, make a func that counts papers in adjacent cells
	movable := 0
	for pos := range grid.positions() {
		movable += grid.CanMove(pos)
	}
	return movable, nil
}

func Part2(r io.Reader) (int, error) {
	grid, err := ParseIn(r)
	if err != nil {
		return 0, err
	}

	// Option 1: Iterate multiple waves of removal

	// Option 2: Iterate once, but each time a cell is removed, check how it affected all adjacent cells
	removed := 0
	for pos := range grid.positions() {
		removed += grid.TryRemoval(pos)
	}

	return removed, nil
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
		width:  width,
		height: height,
	}, nil
}
