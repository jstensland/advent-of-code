// Package day2 solves AoC 2025 day 2
package day2

import (
	"io"
)

// InInfo is a go representation of the input.
// TODO: Update name, type, attributes etc. to match the day.
type InInfo struct{}

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

func ParseIn(r io.Reader) (InInfo, error) {
	// TODO: Update to parse input
	_ = r
	return InInfo{}, nil
}
