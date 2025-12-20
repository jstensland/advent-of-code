// Package day8 solves AoC 2025 day 8
package day8

import (
	"io"
)

func Part1(r io.Reader) (int, error) {
	const part1Iterations = 1000
	return Part1N(r, part1Iterations)
}

func Part1N(r io.Reader, rounds int) (int, error) {
	points, err := ParseIn(r)
	if err != nil {
		return 0, err
	}

	f := NewField(points)
	f.ConnectN(rounds)
	const top3 = 3
	return f.MultiplyLargest(top3), nil
}

func Part2(r io.Reader) (int, error) {
	points, err := ParseIn(r)
	if err != nil {
		return 0, err
	}

	f := NewField(points)
	d := f.ConnectAll()

	ends := d.Ends()
	return ends[0].X * ends[1].X, nil
}
