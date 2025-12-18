// Package day6 solves AoC 2025 day 6
package day6

import (
	"io"
)

func Part1(r io.Reader) (int, error) {
	worksheet, err := ParseIn(r)
	if err != nil {
		return 0, err
	}

	total := 0
	for _, answer := range worksheet.Solve() {
		total += answer
	}

	return total, nil
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
