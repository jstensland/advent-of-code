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
	return (o + 1) % 4 //nolint:mnd // 4 is the number of Orientations
}

type Location struct {
	Row int
	Col int
}

type GuardPosition struct {
	location    Location
	orientation Orientation
}

type Layout struct {
	layout        [][]CellStatus
	guardPosition GuardPosition
}

// LoopCheck returns if this configuration has a loop or not
// it makes a copy of the layout for the check, so it's safe to call in parallel
//
// Optimizations:
// - If I go the worker pool route, reset and reuse layouts
func (l *Layout) LoopCheck(loc Location) bool {
	testLayout := Copy(l)
	// set the hazard
	testLayout.layout[loc.Row][loc.Col] = Hazard
	return testLayout.PatrolTest()
}

func Copy(l *Layout) *Layout {
	newLayout := make([][]CellStatus, len(l.layout))
	for i := range l.layout {
		newLayout[i] = make([]CellStatus, len(l.layout[i]))
		copy(newLayout[i], l.layout[i])
	}
	return &Layout{layout: newLayout, guardPosition: l.guardPosition}
}

// PatrolTest walk the guard until the guard gets back to starting location and orientation
// or walks off the map
// returns false if no loop, but true if loop
func (l *Layout) PatrolTest() bool {
	pastPositions := []GuardPosition{l.guardPosition}
	for {
		currentPosition := l.guardPosition
		if l.OffMap(l.guardPosition.location) {
			return false
		}

		// update position
		if nextCell, ok := l.checkFront(); ok {
			l.layout[currentPosition.location.Row][currentPosition.location.Col] = Visited // mark current cell visited
			l.guardPosition.location = nextCell                                            // update location
		} else {
			l.guardPosition.orientation = Turn(currentPosition.orientation)
		}

		// check if we've been here
		if slices.Contains(pastPositions, l.guardPosition) {
			return true
		}
		pastPositions = append(pastPositions, l.guardPosition) // record updated position
		// fmt.Println(l.guardPosition)
		// fmt.Println(pastPositions)
	}
}

func (l *Layout) Patrol() {
	for {
		currentPosition := l.guardPosition

		if l.OffMap(currentPosition.location) {
			return
		}

		if nextCell, ok := l.checkFront(); ok {
			l.layout[currentPosition.location.Row][currentPosition.location.Col] = Visited // mark current cell visited
			l.guardPosition.location = nextCell                                            // update location
		} else {
			l.guardPosition.orientation = Turn(currentPosition.orientation)
		}
	}
}

// checkFront returns the forward cell, and if it's a Hazard
func (l *Layout) checkFront() (Location, bool) {
	var forwardCell Location
	currentLoc := l.guardPosition.location
	switch l.guardPosition.orientation {
	case Up:
		forwardCell = Location{currentLoc.Row - 1, currentLoc.Col}
	case Right:
		forwardCell = Location{currentLoc.Row, currentLoc.Col + 1}
	case Down:
		forwardCell = Location{currentLoc.Row + 1, currentLoc.Col}
	case Left:
		forwardCell = Location{currentLoc.Row, currentLoc.Col - 1}
	}

	if l.OffMap(forwardCell) {
		return forwardCell, true // let him walk off
	}
	return forwardCell, l.layout[forwardCell.Row][forwardCell.Col] != Hazard
}

// Count visited location
func (l *Layout) Count() int {
	count := 0
	for location := range l.Locations() {
		if l.layout[location.Row][location.Col] == Visited {
			count++
		}
	}
	return count
}

// PatrolledLocations will return locations that were visited by the guard
func (l *Layout) PatrolledLocations() iter.Seq[Location] {
	return func(yield func(l Location) bool) {
		for location := range l.Locations() {
			if l.layout[location.Row][location.Col] == Visited {
				if !yield(location) {
					return
				}
			}
		}
	}
}

func (l *Layout) Locations() iter.Seq[Location] {
	return func(yield func(l Location) bool) {
		for row := range l.Height() {
			for col := range l.Width() {
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

// ParseInput loads the initial layout.
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

	return &Layout{layout: grid, guardPosition: GuardPosition{location: *guardLoc, orientation: Up}}, nil
}
