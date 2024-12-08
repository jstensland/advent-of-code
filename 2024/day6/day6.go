// Package day6 solves AoC 2024 day 6.
package day6

import (
	"fmt"
	"io"
	"sync"

	"github.com/jstensland/advent-of-code/2024/runner"
)

func Run(inFile string) error {
	in := runner.Reader(inFile)
	defer in.Close() //nolint:errcheck // no need to check for error
	answer, err := SolvePart1(in)
	if err != nil {
		return err
	}
	fmt.Println("Day 6 part 1:", answer) //nolint:forbidigo // no IO CLI yet

	// TOO SLOW
	fmt.Println("Day 6 part 2:", "skipped") //nolint:forbidigo // no IO CLI yet
	// in2 := runner.Reader(inFile)
	// defer in2.Close()
	// answer, err = SolvePart2(in2)
	// if err != nil {
	//   return err
	// }
	// fmt.Println("Day 6 part 2:", answer)
	return nil
}

func SolvePart1(in io.Reader) (int, error) {
	layout, err := ParseInput(in)
	if err != nil {
		return 0, fmt.Errorf("error parsing input file: %w", err)
	}

	layout.Patrol()

	return layout.Count(), nil
}

func SolvePart2(in io.Reader) (int, error) {
	layout, err := ParseInput(in)
	if err != nil {
		return 0, fmt.Errorf("error parsing input file: %w", err)
	}

	count := 0
	// go through each location on the map, placing a hazard and checking
	// for a loop

	// This is slow still...
	// Optimizaitons:
	// - Is there any info worth caching if hazards have moved?
	wg := sync.WaitGroup{}
	for location := range layout.Locations() {
		if layout.layout[location.Row][location.Col] == Empty {
			// fmt.Println("checking hazard location", location)

			wg.Add(1)
			go func() {
				if layout.LoopCheck(location) {
					count++
				}
				wg.Done()
			}()
		}
	}
	wg.Wait()

	return count, nil
}

// before the parallel checks
// ok      advent2024/day6 145.813s
//
// after
// ok      advent2024/day6 16.450s
