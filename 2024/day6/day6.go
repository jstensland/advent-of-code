package day6

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"iter"
	"slices"
)

var ErrNoGuard = errors.New("no guard found in layout")

type CellStatus int

const (
	Empty CellStatus = iota
	Visited
	Hazard
	Guard
)

type Orientation int

const (
	Up Orientation = iota
	Right
	Down
	Left
)

// Turn returns the orientation to the right of the given
// Orientation
func Turn(o Orientation) Orientation {
	return Orientation((o + 1) % 4)
}

func RunPart1(in io.Reader) (int, error) {
	layout, err := ParseInput(in)
	if err != nil {
		return 0, fmt.Errorf("error parsing input file: %w", err)
	}

	layout.Patrol()

	return layout.Count(), nil
}

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
	for location := range layout.Locations() {
		if layout.layout[location.Row][location.Col] == Empty {
			fmt.Println("checking hazard location", location)
			if layout.LoopCheck(location) {
				count++
			}
		}
	}

	return count, nil
}

// To check if a location makes a loop, record the starting position and orientation
// Run Patrol (or a version of it) until that same position is reached (loop), OR the guard
// walks off the map (no loop)

type Location struct {
	Row int
	Col int
}

type GuardPosition struct {
	location    Location
	orientation Orientation
}

type Layout struct {
	layout [][]CellStatus
	// TODO: refactor these to use GuardPosition
	guardLoc         Location
	guardOrientation Orientation
}

// LoopCheck returns if this configuration has a loop or not
// Optimizations:
// - If I go the worker pool route, reset and reuse layouts
func (l *Layout) LoopCheck(loc Location) bool {
	// copy the layout
	newLayout := make([][]CellStatus, len(l.layout))
	for i := range l.layout {
		newLayout[i] = make([]CellStatus, len(l.layout[i]))
		copy(newLayout[i], l.layout[i])
	}
	// set the hazard
	newLayout[loc.Row][loc.Col] = Hazard

	// create a new layout with the hazard placed
	testLayout := Layout{layout: newLayout, guardLoc: l.guardLoc, guardOrientation: l.guardOrientation}
	return testLayout.PatrolTest()
}

// PatrolTest will until the guard gets back to starting location and orientation, or leaves
// returns false if no loop, but true if loop
func (l *Layout) PatrolTest() bool {
	pastPositions := []GuardPosition{{location: l.guardLoc, orientation: l.guardOrientation}}
	for {
		if l.OffMap(l.guardLoc) {
			return false
		}

		if nextCell, ok := l.checkFront(); ok {
			l.layout[l.guardLoc.Row][l.guardLoc.Col] = Visited // mark current cell visited
			l.guardLoc = nextCell                              // update location
		} else {
			l.guardOrientation = Turn(l.guardOrientation)
		}

		if slices.Contains(pastPositions, GuardPosition{l.guardLoc, l.guardOrientation}) {
			// he's been here before!
			return true
		}
		pastPositions = append(pastPositions, GuardPosition{l.guardLoc, l.guardOrientation})
	}
}

func (l *Layout) Patrol() {
	for {
		if l.OffMap(l.guardLoc) {
			return
		}

		if nextCell, ok := l.checkFront(); ok {
			l.layout[l.guardLoc.Row][l.guardLoc.Col] = Visited // mark current cell visited
			l.guardLoc = nextCell                              // update location
		} else {
			l.guardOrientation = Turn(l.guardOrientation)
		}
	}
}

// checkFront returns the forward cell, and if it's a Hazard
func (l *Layout) checkFront() (Location, bool) {
	var forwardCell Location
	switch l.guardOrientation {
	case Up:
		forwardCell = Location{l.guardLoc.Row - 1, l.guardLoc.Col}
	case Right:
		forwardCell = Location{l.guardLoc.Row, l.guardLoc.Col + 1}
	case Down:
		forwardCell = Location{l.guardLoc.Row + 1, l.guardLoc.Col}
	case Left:
		forwardCell = Location{l.guardLoc.Row, l.guardLoc.Col - 1}
	}

	if l.OffMap(forwardCell) {
		return forwardCell, true // let him walk off
	}
	return forwardCell, l.layout[forwardCell.Row][forwardCell.Col] != Hazard
}

func (l *Layout) Count() int {
	count := 0
	for location := range l.Locations() {
		if l.layout[location.Row][location.Col] == Visited {
			count++
		}
	}
	return count
}

func (l *Layout) Locations() iter.Seq[Location] {
	return func(yield func(l Location) bool) {
		for row := 0; row < l.Height(); row++ {
			for col := 0; col < l.Width(); col++ {
				if !yield(Location{row, col}) {
					return
				}
			}
		}
	}
}

func (l *Layout) OffMap(loc Location) bool {
	return loc.Col < 0 || loc.Col >= l.Width() || loc.Row < 0 || loc.Row >= l.Height()
}

func (l *Layout) Height() int {
	return len(l.layout)
}

func (l *Layout) Width() int {
	return len(l.layout[0])
}

// ParseInput
func ParseInput(in io.Reader) (*Layout, error) {
	scanner := bufio.NewScanner(in)
	var grid [][]CellStatus
	var guardLoc *Location

	charToStatus := map[rune]CellStatus{
		'.': Empty,
		'#': Hazard,
		'^': Guard,
		'X': Visited, // since none start visted, this is for testing only
	}

	for scanner.Scan() {
		line := scanner.Text()
		gridRow := make([]CellStatus, len(line))
		row := []rune(line)
		for idx, cell := range row {
			gridRow[idx] = charToStatus[cell]
			if gridRow[idx] == Guard {
				guardLoc = &Location{len(grid), idx}
			}
		}
		grid = append(grid, gridRow)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failure scanning input: %w", err)
	}

	if guardLoc == nil {
		return nil, ErrNoGuard
	}

	return &Layout{layout: grid, guardLoc: *guardLoc, guardOrientation: Up}, nil
}
