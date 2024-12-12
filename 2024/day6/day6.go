// Package day6 solves AoC 2024 day 6.
package day6

import (
	"fmt"
	"io"
	"sync"
)

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
	initialPatrol := Copy(layout)
	initialPatrol.Patrol()

	count := 0
	// go through each location on the map where he patrolled and place a hazard
	// test for the loop

	// Optimizaitons:
	// - Is there any info worth caching if hazards have moved?
	// - copy less.
	wg := sync.WaitGroup{}
	for location := range initialPatrol.PatrolledLocations() {
		// fmt.Println("checking hazard location", location)

		wg.Add(1)
		go func() {
			if layout.LoopCheck(location) {
				count++
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return count, nil
}

// original
// ok      advent2024/day6 145.813s
//
// parallel checking
// ok      advent2024/day6 16.450s
//
// after only checking visited locations
//
// place barrels on original path instead of everywhere
// ok      github.com/jstensland/advent-of-code/2024/day6  2.648s
