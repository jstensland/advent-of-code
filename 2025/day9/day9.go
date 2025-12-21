// Package day9 solves AoC 2025 day 9
package day9

import (
	"io"
	"slices"
)

func Part1(r io.Reader) (int, error) {
	in, err := ParseIn(r)
	if err != nil {
		return 0, err
	}

	minX, maxX, minY, maxY := FindExtremes(in)
	// Sort the points by distance from the center
	slices.SortFunc(in, FurtherFromCenter(minX, maxX, minY, maxY))
	pool := NewCandidatePool()

	for _, p := range in {
		// if the biggest square from the current point to a corner is SMALLER than our current biggest, then quit
		if biggestToCorner(minX, maxX, minY, maxY, p) < pool.Biggest().Area() {
			break
		}
		pool.Add(p)
	}
	return pool.Biggest().Area(), nil
}

func Part2(r io.Reader) (int, error) {
	_, err := ParseIn(r)
	if err != nil {
		return 0, err
	}
	// TODO: solve part 2
	answer := 0

	return answer, nil
}
