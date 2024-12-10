// Package day10 solves AoC 2024 day 9.
package day10

import (
	"bufio"
	"fmt"
	"io"
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
	for i, row := range layout.data {
		for j, val := range row {
			if val == 0 {
				// // fmt.Println(i, j)
				total += layout.Score(Location{i, j})
			}
		}
	}

	return total, nil
}

func SolvePart2(in io.Reader) (int, error) {
	layout, err := ParseInput(in)
	if err != nil {
		return 0, fmt.Errorf("error parsing input file: %w", err)
	}

	total := 0
	for i, row := range layout.data {
		for j, val := range row {
			if val == 0 {
				// // fmt.Println(i, j)
				total += layout.Rating(Location{i, j}, val)
			}
		}
	}
	return total, nil
}

type Grid struct {
	data   [][]int
	width  int
	height int
}

type Location struct {
	Row int
	Col int
}

func (g *Grid) Score(loc Location) int {
	// // fmt.Println("scoring", loc)

	// count how many 9s you can get to
	locations := g.SeekNine(loc, 0)

	// // fmt.Println(locations)

	slices.SortFunc(locations, func(a, b Location) int {
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
	})

	return len(slices.Compact(locations))
}

// Rating should count how many ways you can get to a 9
func (g *Grid) Rating(start Location, here int) int {
	if here != g.data[start.Row][start.Col] {
		// TODO: refactor away this possibility
		panic("something is going wrong")
	}
	// fmt.Println("seeking", start, here)
	if here == endOfTheRoad {
		return 1
	}

	total := 0
	next := here + 1
	if start.Row-1 >= 0 && g.data[start.Row-1][start.Col] == next {
		// fmt.Println("look up")
		total += g.Rating(Location{start.Row - 1, start.Col}, next)
	}

	if start.Row+1 < g.height && g.data[start.Row+1][start.Col] == next {
		// fmt.Println("down")
		total += g.Rating(Location{start.Row + 1, start.Col}, next)
	}

	if start.Col-1 >= 0 && g.data[start.Row][start.Col-1] == next {
		// fmt.Println("left")
		total += g.Rating(Location{start.Row, start.Col - 1}, next)
	}

	if start.Col+1 < g.width && g.data[start.Row][start.Col+1] == next {
		// fmt.Println("right")
		total += g.Rating(Location{start.Row, start.Col + 1}, next)
	}
	return total
}

// SeekNine finds walks the grid to find a 9 along an increasing path. It
// returns a slice of those locations.
func (g *Grid) SeekNine(start Location, here int) []Location {
	if here != g.data[start.Row][start.Col] {
		panic("something is going wrong")
	}

	// // fmt.Println("seeking", start, here)
	if here == endOfTheRoad {
		// // fmt.Println("end of the road:", start)
		return []Location{start}
	}

	var nextSteps []Location

	next := here + 1

	if start.Row-1 >= 0 && g.data[start.Row-1][start.Col] == next {
		// // fmt.Println("look up")
		ups := g.SeekNine(Location{start.Row - 1, start.Col}, next)
		nextSteps = append(nextSteps, ups...)
	}

	if start.Row+1 < g.height && g.data[start.Row+1][start.Col] == next {
		// // fmt.Println("down")
		downs := g.SeekNine(Location{start.Row + 1, start.Col}, next)
		nextSteps = append(nextSteps, downs...)
	}

	if start.Col-1 >= 0 && g.data[start.Row][start.Col-1] == next {
		// // fmt.Println("left")
		lefts := g.SeekNine(Location{start.Row, start.Col - 1}, next)
		nextSteps = append(nextSteps, lefts...)
	}

	if start.Col+1 < g.width && g.data[start.Row][start.Col+1] == next {
		// // fmt.Println("right")
		rights := g.SeekNine(Location{start.Row, start.Col + 1}, next)
		nextSteps = append(nextSteps, rights...)
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
