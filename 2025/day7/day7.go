// Package day7 solves AoC 2025 day 7
package day7

import (
	"io"
)

func Part1(r io.Reader) (int, error) {
	grid, err := ParseIn(r)
	if err != nil {
		return 0, err
	}

	for range grid.height {
		grid.Progress()
	}

	return grid.SplitCount(), nil
}

func Part2(r io.Reader) (int, error) {
	grid, err := ParseIn(r)
	if err != nil {
		return 0, err
	}
	startRow, startCol := grid.Start()
	answer := grid.ProgressTimeline(startRow+1, startCol)

	return answer, nil
}
