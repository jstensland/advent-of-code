// Package day8 solves AoC 2024 day 8.
package day8

import (
	"bufio"
	"fmt"
	"io"
)

func SolvePart1(in io.Reader) (int, error) {
	layout, err := ParseInput(in)
	if err != nil {
		return 0, fmt.Errorf("error parsing input file: %w", err)
	}
	// fmt.Println(layout)
	layout.CalculateAntinodes(false)
	// fmt.Println(layout)
	return layout.CountAntinodes(), nil
}

func SolvePart2(in io.Reader) (int, error) {
	layout, err := ParseInput(in)
	if err != nil {
		return 0, fmt.Errorf("error parsing input file: %w", err)
	}
	// fmt.Println(layout)
	layout.CalculateAntinodes(true)
	// fmt.Println(layout) // make them visible?
	return layout.CountAntinodes(), nil
}

// CalculateAntinodes updates the map cells with which locations are antinodes based on the
// positions of the antennae. A location is an antinode if it is in line with another antennae
// of the same type and twice as far from one than the other.
func (l *Layout) CalculateAntinodes(part2 bool) {
	// go through the various types of attennae, and calculate their antinodes
	for _, locs := range l.antennas {
		l.AddAntinodes(locs, part2)
	}
	l.CountAntinodes()
}

// AddAntinodes adds the antinodes for antenna in the set of locations given
// it must be aware of the map size. Overwriting is allowed as it will only
// ever overwrite true with true
func (l *Layout) AddAntinodes(locs []Location, part2 bool) {
	// for each pair, calculate the antinodes
	for _, loc1 := range locs {
		for _, loc2 := range locs {
			if loc1 == loc2 {
				continue // antenna don't resinate with themselves
			}

			var locs []Location
			if part2 {
				locs = possibleAntiNodes2(loc1, loc2, l.Width(), l.Height())
			} else {
				locs = possibleAntiNodes(loc1, loc2)
			}
			for _, loc := range locs {
				if !OffMap(loc, l.Width(), l.Height()) {
					l.layout[loc.Row][loc.Col].antinode = true
				}
			}
		}
	}
}

// possibleAntiNodes gives back the two possible antinodes, disregarding
// map size
func possibleAntiNodes(loc1, loc2 Location) []Location {
	return []Location{
		{loc1.Row + (loc1.Row - loc2.Row), loc1.Col + (loc1.Col - loc2.Col)},
		{loc2.Row + (loc2.Row - loc1.Row), loc2.Col + (loc2.Col - loc1.Col)},
	}
}

func possibleAntiNodes2(loc1, loc2 Location, width, height int) []Location {
	// IMPROVEMENT: simplify...
	out := []Location{loc1} // start at one of the nodes and go both ways
	colDiff := loc1.Col - loc2.Col
	rowDiff := loc1.Row - loc2.Row
	// - go in each direction, adding locations while they're on the map
	resonation := 1
	nextAntinode := Location{loc1.Row + rowDiff*resonation, loc1.Col + colDiff*resonation}
	// start with the higher one, and resonate downward
	for !OffMap(nextAntinode, width, height) {
		out = append(out, nextAntinode)
		resonation++
		nextAntinode = Location{loc1.Row + rowDiff*resonation, loc1.Col + colDiff*resonation}
	}
	// back to loc1, but going the other way
	nextAntinode = Location{loc1.Row - rowDiff*resonation, loc1.Col - colDiff*resonation}
	for !OffMap(nextAntinode, width, height) {
		out = append(out, nextAntinode)
		resonation++
		nextAntinode = Location{loc1.Row - rowDiff*resonation, loc1.Col - colDiff*resonation}
	}
	return out
}

func (l *Layout) CountAntinodes() int {
	total := 0
	for _, row := range l.layout {
		for _, cell := range row {
			if cell.antinode {
				total++
			}
		}
	}
	return total
}

type Antenna rune

type CellStatus struct {
	antenna  Antenna
	antinode bool
}

type Layout struct {
	layout   [][]CellStatus
	antennas map[Antenna][]Location
}

type Location struct {
	Row int
	Col int
}

// String implements Stringer and allows printing of current state of the map
func (l *Layout) String() string {
	out := ""
	for _, row := range l.layout {
		for _, cell := range row {
			if cell.antenna != 0 {
				out += string(cell.antenna)
			} else {
				if cell.antinode {
					out += "#" // match prompt on depiction of antinodes
				} else {
					out += "."
				}
			}
		}
		out += "\n"
	}
	return out
}

func OffMap(loc Location, width, height int) bool {
	return loc.Col < 0 || loc.Col >= width || loc.Row < 0 || loc.Row >= height
}

func (l *Layout) Height() int {
	return len(l.layout)
}

func (l *Layout) Width() int {
	return len(l.layout[0])
}

// ParseInput loads the initial layout.
// Collect antenna locations by type as you go
func ParseInput(in io.Reader) (*Layout, error) {
	scanner := bufio.NewScanner(in)
	var grid [][]CellStatus
	antennasLocs := map[Antenna][]Location{}

	for scanner.Scan() {
		line := scanner.Text()
		gridRow := make([]CellStatus, len(line))
		row := []rune(line)
		for idx, charRune := range row {
			if charRune == '.' {
				// empty
				gridRow[idx] = CellStatus{}
			} else {
				// everything other than '.' is an antenna
				ant := Antenna(charRune)
				gridRow[idx] = CellStatus{antenna: ant}
				antennasLocs[ant] = append(antennasLocs[ant], Location{len(grid), idx})
			}
		}
		grid = append(grid, gridRow)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failure scanning input: %w", err)
	}

	return &Layout{grid, antennasLocs}, nil
}
