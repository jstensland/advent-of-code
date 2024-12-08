// Package day6 solves AoC 2024 day 6.
package day6

import (
	"fmt"
	"io"
	"sync"
)

// RunPart1 solves part 1.
func RunPart1(in io.Reader) (int, error) {
	layout, err := ParseInput(in)
	if err != nil {
		return 0, fmt.Errorf("error parsing input file: %w", err)
	}

	layout.Patrol()

	return layout.Count(), nil
}

// RunPart2 solves part 2.
func RunPart2(in io.Reader) (int, error) {
	layout, err := ParseInput(in)
	if err != nil {
		return 0, fmt.Errorf("error parsing input file: %w", err)
	}

	count := 0
	// go through each location on the map, placing a hazard and checking
	// for a loop

	// This is too slow...
	// Optimizaitons:
	// - Check each location in parallel, or with a worker pool
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
