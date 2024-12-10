// Package day10 solves AoC 2024 day 10.
package day10

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"log"
	"slices"
	"strconv"
)

const endOfTheRoad = 9

func SolvePart1(in io.Reader) (int, error) {
	layout, err := ParseInput(in)
	if err != nil {
		return 0, fmt.Errorf("error parsing input file: %w", err)
	}
	total := 0
	for th := range layout.trailheads() {
		total += layout.Score(th)
	}
	return total, nil
}

func SolvePart2(in io.Reader) (int, error) {
	layout, err := ParseInput(in)
	if err != nil {
		return 0, fmt.Errorf("error parsing input file: %w", err)
	}
	total := 0
	for th := range layout.trailheads() {
		total += layout.Rating(th)
	}
	return total, nil
}

type Grid struct {
	data   [][]int
	width  int
	height int
}

// Value returns the value at the given location.
func (g *Grid) Value(loc Location) int {
	return g.data[loc.Row][loc.Col]
}

// IsOnGrid returns true if the location is within the grid.
func (g *Grid) IsOnGrid(loc Location) bool {
	return loc.Row >= 0 && loc.Row < g.height && loc.Col >= 0 && loc.Col < g.width
}

// Score is the method for part 1, counting how many 9s you can get to.
func (g *Grid) Score(loc Location) int {
	// count how many 9s you can get to
	locations := g.seekNine(loc, g.data[loc.Row][loc.Col])
	slices.SortFunc(locations, LocationSort)
	return len(slices.Compact(locations))
}

// Rating is for part 2, counting how many ways you can get to a 9.
func (g *Grid) Rating(start Location) int {
	return g.rating(start, g.data[start.Row][start.Col])
}

func (g *Grid) rating(start Location, elevation int) int {
	// fmt.Println("seeking", start, here)
	if elevation == endOfTheRoad {
		return 1
	}

	total := 0
	next := elevation + 1
	for _, dir := range []Location{start.Up(), start.Down(), start.Left(), start.Right()} {
		if g.IsOnGrid(dir) && g.Value(dir) == next {
			total += g.rating(dir, next)
		}
	}
	return total
}

// seekNine finds walks the grid to find a 9 along an increasing path. It
// returns a slice of those locations.
func (g *Grid) seekNine(start Location, here int) []Location {
	if here == endOfTheRoad {
		return []Location{start}
	}
	var nextSteps []Location
	next := here + 1
	for _, dir := range []Location{start.Up(), start.Down(), start.Left(), start.Right()} {
		if g.IsOnGrid(dir) && g.Value(dir) == next {
			nextSteps = append(nextSteps, g.seekNine(dir, next)...)
		}
	}
	return nextSteps
}

func ParseInput(in io.Reader) (*Grid, error) {
	scanner := bufio.NewScanner(in)
	grid := Grid{[][]int{}, 0, 0}
	for scanner.Scan() {
		inRow := scanner.Text()
		row := make([]int, len(inRow))
		for idx, cell := range inRow {
			if cell == '.' {
				// for testing
				row[idx] = -1
			} else {
				val, err := strconv.Atoi(string(cell))
				if err != nil {
					return nil, fmt.Errorf("error parsing input file: %w", err)
				}
				row[idx] = val
			}
		}
		grid.data = append(grid.data, row)
		grid.height++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	grid.width = len(grid.data[0])
	return &grid, nil
}

// IMPROVEMENT: could record these while parsing rather than dynamically discovered
func (g *Grid) trailheads() iter.Seq[Location] {
	return func(yield func(l Location) bool) {
		for i, row := range g.data {
			for j, val := range row {
				if val == 0 {
					if !yield(Location{i, j}) {
						return
					}
				}
			}
		}
	}
}

type Location struct {
	Row int
	Col int
}

func (l Location) Up() Location    { return Location{l.Row - 1, l.Col} }
func (l Location) Down() Location  { return Location{l.Row + 1, l.Col} }
func (l Location) Left() Location  { return Location{l.Row, l.Col - 1} }
func (l Location) Right() Location { return Location{l.Row, l.Col + 1} }

func LocationSort(a, b Location) int {
	// treat row as more significant
	if a.Row < b.Row {
		return -1
	}
	if a.Row > b.Row {
		return 1
	}

	// if rows are equal, check columns
	if a.Col < b.Col {
		return -1
	}
	if a.Col > b.Col {
		return 1
	}
	// all equal
	return 0
}
