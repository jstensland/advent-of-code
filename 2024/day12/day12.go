// Package day12 solves AoC 2024 day 12
package day12

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"log"
	"slices"
)

// SolvePart1 finds occurrences of XMAS in a wordsearch fashion.
//
// It's a small input, so parse the whole thing into memory
// and then search for XMAS from every X, checking each of the 8
// possible directions.
func SolvePart1(in io.Reader) (int, error) {
	grid, err := ParseGrid(in)
	if err != nil {
		return 0, fmt.Errorf("error loading grid: %w", err)
	}

	return grid.FencePrice(), nil
}

func SolvePart2(in io.Reader) (int, error) {
	grid, err := ParseGrid(in)
	if err != nil {
		return 0, fmt.Errorf("error loading grid: %w", err)
	}

	return grid.BulkFencePrice(), nil
}

type Grid struct {
	data         [][]rune
	lastRegionID int
	regions      map[int]Region // identifier to region mapping.
	width        int
	height       int
}

func (g *Grid) FencePrice() int {
	price := 0
	for _, r := range g.regions {
		price += r.Area() * r.Permimeter()
	}
	return price
}

func (g *Grid) BulkFencePrice() int {
	price := 0
	for _, r := range g.regions {
		price += r.Area() * r.StraightSides()
	}
	return price
}

func (g *Grid) Width() int  { return g.width }
func (g *Grid) Height() int { return g.height }
func (g *Grid) Get(l Position) rune {
	if l.row < 0 || l.row >= g.height || l.col < 0 || l.col >= g.width {
		return rune(0)
	}
	return g.data[l.row][l.col]
}

func (g *Grid) newRegion(label rune, loc Position) Region {
	g.lastRegionID++
	return Region{
		id:    g.lastRegionID,
		label: label,
		area:  []Position{loc},
	}
}

// FindRegions scans the grid and collects contiguous regions
func (g *Grid) FindRegions() {
	// go through each cell in the grid
	for loc := range g.positions() {
		g.AddToRegion(loc)
	}
}

// AddToRegion takes a location, and decides if it should be added to an existing
// region, or create a new one. If it is added to an existing region, it will also
// check if that causes two regions two merge.
func (g *Grid) AddToRegion(loc Position) {
	val := g.Get(loc)
	adjacentRegions := map[int]Region{}
	// check bordering cells for the same values
	for _, side := range g.Contiguous(loc) {
		for id, reg := range g.regions {
			// check if those cells are in regions
			// IMPROVEMENT: this lookup is rough. Keep a map of position to region?
			if val == reg.label && slices.Contains(reg.area, side) {
				adjacentRegions[id] = reg
			}
		}
	}

	switch len(adjacentRegions) {
	case 0:
		// If no bordering regions with the same value, make a new region and store it.
		newReg := g.newRegion(val, loc)
		g.regions[newReg.id] = newReg
	case 1:
		// if only one bordering region, join it
		var region Region
		for _, reg := range adjacentRegions {
			region = reg // only one iteration
		}
		region.area = append(region.area, loc)
		g.regions[region.id] = region
	default:
		// if multiple border regions with this val, merge them
		g.mergeRegions(adjacentRegions, loc)
	}
}

func (g *Grid) mergeRegions(toMerge map[int]Region, loc Position) {
	first := true
	var region Region
	toDelete := make([]int, 0, len(toMerge)-1)
	for _, reg := range toMerge {
		if first { // grab the first region
			region = reg
			first = false
			continue
		}
		toDelete = append(toDelete, reg.id)
		// add the rest to the first
		region.area = append(region.area, reg.area...)
	}
	// in the current loc
	region.area = append(region.area, loc)
	g.regions[region.id] = region // save

	// remove merged regions
	for _, id := range toDelete {
		delete(g.regions, id)
	}
}

// Contiguous returns sides of a location that share its value.
func (g *Grid) Contiguous(loc Position) []Position {
	out := []Position{}

	val := g.Get(loc)
	for _, side := range loc.adjacent() {
		if loc.row < 0 || loc.row >= g.height || loc.col < 0 || loc.col >= g.width {
			continue // skip if off the grid
		}
		if g.Get(side) == val {
			out = append(out, side)
		}
	}
	return out
}

func (g *Grid) positions() iter.Seq[Position] {
	return func(yield func(r Position) bool) {
		for row := range g.height {
			for col := range g.width {
				if !yield(Position{row, col}) {
					return
				}
			}
		}
	}
}

type (
	Border struct {
		loc  Position
		side side
	}
	side string
)

const (
	top    side = "top"
	bottom side = "bottom"
	left   side = "left"
	right  side = "right"
)

func SortBorders(a, b Border) int {
	if a.side == top || a.side == bottom {
		return a.loc.col - b.loc.col
	}
	// left/right
	return a.loc.row - b.loc.row
}

type Region struct {
	id    int
	label rune
	area  []Position
}

func (r Region) Area() int { return len(r.area) }
func (r Region) Permimeter() int {
	perimeter := 0
	// go through each location in the area
	for _, loc := range r.area {
		perimeter += r.countBorders(loc)
	}
	// I could track border locations as I form the regions if I needed more efficiency,
	// but I doubt it
	return perimeter
}

// countBorders deterins how many sides of a given area in a region are part of the border
// returns a number between 0 and 4
func (r Region) countBorders(loc Position) int {
	borders := 0
	for _, side := range loc.adjacent() {
		if !slices.Contains(r.area, side) {
			borders++
		}
	}
	return borders
}

// StraightSides returns the number of straight sides for a given region
//
//nolint:gocognit,cyclop // IMPROVEMENT, remove this. for now choosing to ignore complexity
func (r Region) StraightSides() int {
	borders := []Border{}

	// go through each block of the region to identify the borders
	for _, p := range r.area {
		upper := Position{p.row - 1, p.col} // up
		if !slices.Contains(r.area, upper) {
			borders = append(borders, Border{p, top})
		}
		lower := Position{p.row + 1, p.col} // down
		if !slices.Contains(r.area, lower) {
			borders = append(borders, Border{p, bottom})
		}
		righter := Position{p.row, p.col + 1} // right
		if !slices.Contains(r.area, righter) {
			borders = append(borders, Border{p, right})
		}
		lefter := Position{p.row, p.col - 1} // left
		if !slices.Contains(r.area, lefter) {
			borders = append(borders, Border{p, left})
		}
	}
	// fmt.Println("region:", r)
	// fmt.Println("borders:", borders)

	// go through each border and collect them into sides
	buckets := map[string][]Border{}
	for _, b := range borders {
		// bucket them into row/column+side
		var key string
		if b.side == top || b.side == bottom {
			key = fmt.Sprintf("%d%s", b.loc.row, b.side)
		} else {
			key = fmt.Sprintf("%d%s", b.loc.col, b.side)
		}
		buckets[key] = append(buckets[key], b)
	}

	var sides int
	// group them into consequetive sides
	for _, segments := range buckets {
		slices.SortFunc(segments, SortBorders)

		// IMPROVEMENT: remove the extra "sameSide" control structure.
		var sameSide bool // indicates if the next segment looks like part of the same side
		for i, segment := range segments {
			if !sameSide {
				sides++ // new side
			}
			if i+1 == len(segments) {
				sameSide = false // doesn't matter?...
				continue         // no next
			}
			if segment.side == top || segment.side == bottom {
				if segment.loc.col+1 == segments[i+1].loc.col {
					// part of the same side
					sameSide = true
					continue
				}
				// new side
				sameSide = false
			} else {
				if segment.loc.row+1 == segments[i+1].loc.row {
					// part of the same side
					sameSide = true
					continue
				}
				// new side
				sameSide = false
			}
		}
	}
	return sides
}

// ParseGrid reads in the full grid into memory before manipulating it.
//
// Input is 140x140, so opted to manipulate it after for flexibility
// instead of trying to process it as a stream
func ParseGrid(in io.Reader) (Grid, error) {
	scanner := bufio.NewScanner(in)
	grid := Grid{
		data:    [][]rune{},
		regions: map[int]Region{},
		width:   0,
		height:  0,
	}
	for scanner.Scan() {
		row := []rune(scanner.Text())
		grid.data = append(grid.data, row)
		grid.height++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	grid.width = len(grid.data[0])
	grid.FindRegions() // populate regions. IMPROVEMENT: make private, use export_test.go if needed
	return grid, nil
}

type Position struct {
	// row starts at zero and goes top to bottom
	row int
	// col starts at zero and goes left to right
	col int
}

func (p Position) adjacent() []Position {
	return []Position{
		{p.row - 1, p.col}, // up
		{p.row + 1, p.col}, // down
		{p.row, p.col + 1}, // right
		{p.row, p.col - 1}, // left
	}
}
