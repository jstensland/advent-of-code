// Package main is a minimal entry point for running each day to help
// us get out of main and allow more testablility
package main

import (
	"log"

	"github.com/jstensland/advent-of-code/2024/day1"
	"github.com/jstensland/advent-of-code/2024/day2"
	"github.com/jstensland/advent-of-code/2024/day3"
	"github.com/jstensland/advent-of-code/2024/day4"
	"github.com/jstensland/advent-of-code/2024/day5"
	"github.com/jstensland/advent-of-code/2024/day6"
	"github.com/jstensland/advent-of-code/2024/day7"
	"github.com/jstensland/advent-of-code/2024/runner"
)

// TODO: create a runner package that's invoked from here
// at most, this package should specify the days to run, similar
// to commandline arguments
func main() {
	for _, day := range []struct {
		name string
		fn   runner.Solver
		in   string
	}{
		{"Day 1 Part 1", day1.SolvePart1, "./day1/input.txt"},
		{"Day 1 Part 2", day1.SolvePart2, "./day1/input.txt"},
		{"Day 2 Part 1", day2.SolvePart1, "./day2/input.txt"},
		{"Day 2 Part 2", day2.SolvePart2, "./day2/input.txt"},
		{"Day 3 Part 1", day3.SolvePart1, "./day3/input.txt"},
		{"Day 3 Part 2", day3.SolvePart2, "./day3/input.txt"},
		{"Day 4 Part 1", day4.SolvePart1, "./day4/input.txt"},
		{"Day 4 Part 2", day4.SolvePart2, "./day4/input.txt"},
		{"Day 5 Part 1", day5.SolvePart1, "./day5/input.txt"},
		{"Day 5 Part 2", day5.SolvePart2, "./day5/input.txt"},
		{"Day 6 Part 1", day6.SolvePart1, "./day6/input.txt"},
		// {"Day 6 Part 2", day6.SolvePart2, "./day6/input.txt"}, // too slow
		{"Day 7 Part 1", day7.SolvePart1, "./day7/input.txt"},
		{"Day 7 Part 2", day7.SolvePart2, "./day7/input.txt"},
	} {
		err := runner.RunIt(day.name, day.fn, day.in)
		if err != nil {
			log.Fatal()
		}
	}
}
