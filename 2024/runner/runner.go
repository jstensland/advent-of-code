// Package runner has utilities for handling input files and day runs.
package runner

import (
	"fmt"
	"io"

	"github.com/jstensland/advent-of-code/2024/day1"
	"github.com/jstensland/advent-of-code/2024/day10"
	"github.com/jstensland/advent-of-code/2024/day11"
	"github.com/jstensland/advent-of-code/2024/day12"
	"github.com/jstensland/advent-of-code/2024/day13"
	"github.com/jstensland/advent-of-code/2024/day14"
	"github.com/jstensland/advent-of-code/2024/day15"
	"github.com/jstensland/advent-of-code/2024/day16"
	"github.com/jstensland/advent-of-code/2024/day2"
	"github.com/jstensland/advent-of-code/2024/day3"
	"github.com/jstensland/advent-of-code/2024/day4"
	"github.com/jstensland/advent-of-code/2024/day5"
	"github.com/jstensland/advent-of-code/2024/day6"
	"github.com/jstensland/advent-of-code/2024/day7"
	"github.com/jstensland/advent-of-code/2024/day8"
	"github.com/jstensland/advent-of-code/2024/day9"
	"github.com/jstensland/advent-of-code/2024/input"
)

type Solver func(io.Reader) (int, error)

func Run() error {
	for _, day := range []struct {
		name string
		fn   Solver
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
		{"Day 6 Part 2", day6.SolvePart2, "./day6/input.txt"},
		{"Day 7 Part 1", day7.SolvePart1, "./day7/input.txt"},
		{"Day 7 Part 2", day7.SolvePart2, "./day7/input.txt"},
		{"Day 8 Part 1", day8.SolvePart1, "./day8/input.txt"},
		{"Day 8 Part 2", day8.SolvePart2, "./day8/input.txt"},
		{"Day 9 Part 1", day9.SolvePart1, "./day9/input.txt"},
		{"Day 9 Part 2", day9.SolvePart2, "./day9/input.txt"},
		{"Day 10 Part 1", day10.SolvePart1, "./day10/input.txt"},
		{"Day 10 Part 2", day10.SolvePart2, "./day10/input.txt"},
		{"Day 11 Part 1", day11.SolvePart1, "./day11/input.txt"},
		{"Day 11 Part 2", day11.SolvePart2, "./day11/input.txt"},
		{"Day 12 Part 1", day12.SolvePart1, "./day12/input.txt"},
		{"Day 12 Part 2", day12.SolvePart2, "./day12/input.txt"},
		{"Day 13 Part 1", day13.SolvePart1, "./day13/input.txt"},
		{"Day 13 Part 2", day13.SolvePart2, "./day13/input.txt"},
		{
			"Day 14 Part 1",
			//nolint:mnd // magic numbers are dimensions asked for
			func(in io.Reader) (int, error) { return day14.SolvePart1(in, 103, 101) },
			"./day14/input.txt",
		},
		{
			"Day 14 Part 2",
			//nolint:mnd // magic numbers are dimensions asked for
			func(in io.Reader) (int, error) { return day14.SolvePart2(in, 103, 101) },
			"./day14/input.txt",
		},

		{"Day 15 Part 1", day15.SolvePart1, "./day15/input.txt"},
		{"Day 15 Part 2", day15.SolvePart2, "./day15/input.txt"},
		{"Day 16 Part 1", day16.SolvePart1, "./day16/input.txt"},
		{"Day 16 Part 2", day16.SolvePart2, "./day16/input.txt"},
	} {
		err := RunIt(day.name, day.fn, day.in)
		if err != nil {
			return err
		}
	}
	return nil
}

func RunIt(name string, fn Solver, inFile string) error {
	in := input.Reader(inFile)
	defer in.Close() //nolint:errcheck // no need to check for error

	answer, err := fn(in)
	if err != nil {
		return err
	}
	fmt.Printf("%s: %v\n", name, answer) //nolint:forbidigo // no IO CLI yet
	return nil
}
