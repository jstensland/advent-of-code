package day15

import (
	"fmt"
	"io"
)

func SolvePart2(in io.Reader) (int, error) {
	widthFactor := 2
	offset := widthFactor - 1
	grid, err := ParseIn(in, widthFactor, offset)
	if err != nil {
		return 0, fmt.Errorf("error loading input: %w", err)
	}
	grid.RunRobotsV2()

	return grid.TotalGPSV2(), nil
}

type Location struct {
	Row int
	Col int
}

// Add will add a vector to a location and returns the resulting location.
func (loc Location) Add(vec MoveVector) Location {
	return Location{Row: loc.Row + vec.deltaRow, Col: loc.Col + vec.deltaCol}
}

func (g *Grid) TotalGPSV2() int {
	total := 0
	for row := range g.Height {
		for col := range g.Width {
			// find each box by its left side
			if g.data[row][col] == BoxLeft {
				total += 100*row + col //nolint:mnd // magic number
			}
		}
	}
	return total
}

// RunRobotsV2 moves the robot all the moves but treats BoxLeft and BoxRight squares as one
// ridged body.
func (g *Grid) RunRobotsV2() {
	for _, mv := range g.movesVec {
		// fmt.Println(g)
		// fmt.Println(mv)
		g.maybeMoveV2(mv)
	}
	// fmt.Println(g)
}

func (g *Grid) maybeMoveV2(mv MoveVector) {
	g.doMove(mv, g.robotLoc)
}

// moveable says if it's worth trying to move the next object
func moveable(s State) bool {
	return s == Empty || s == BoxLeft || s == BoxRight
}

// boxLoc is the two sides of a box
type boxLocation struct {
	left  Location
	right Location
}

// doMove checks if the given move is possible from that location. If it is, it first calls itself
// to move any movable object out of the way as necessary, and then does it's own move.
//
// Blocks are checked and moved as a unit.
//
// The return value indicates if the move was possible.
func (g *Grid) doMove(mv MoveVector, currentLoc Location) bool {
	currentVal := g.GetLoc(currentLoc)
	if currentVal == Wall {
		return false
	}
	if currentVal == Empty {
		return true // noop to move an empty spot, but possible, so true
	}
	nextLoc := currentLoc.Add(mv)
	nextVal := g.GetLoc(nextLoc)
	if nextVal == Wall {
		// if the next position is a wall, can't go. Do nothing and return false.
		return false
	}

	if currentVal == Robot {
		// move robot. robot one wide
		if moveable(nextVal) && g.doMove(mv, nextLoc) { // if next value could move, try to move it
			g.moveRobot(currentLoc, nextLoc) // it moved, so move the robot now
			return true
		}
		return false // not movable, or didn't move
	}

	// now for the boxes...
	var boxLoc boxLocation
	if currentVal == BoxLeft {
		boxLoc = boxLocation{left: currentLoc, right: Location{Row: currentLoc.Row, Col: currentLoc.Col + 1}}
	} else if currentVal == BoxRight {
		boxLoc = boxLocation{left: Location{Row: currentLoc.Row, Col: currentLoc.Col - 1}, right: currentLoc}
	}

	// move block
	if mv == upVec || mv == downVec {
		// check of whatever is above the right and left hands sides are moveable.
		if moveable(g.GetLoc(boxLoc.left.Add(mv))) && moveable(g.GetLoc(boxLoc.right.Add(mv))) {
			// if the value directly above the left side of the box is the left side of another box, just move that box
			if g.GetLoc(boxLoc.left.Add(mv)) == BoxLeft {
				if g.doMove(mv, boxLoc.left.Add(mv)) { // move the one box
					g.moveBoxV2(boxLoc, mv)
					return true
				}
			} else {
				// moveable, but not a left box, will need to move both sides
				if g.doMove(mv, boxLoc.left.Add(mv)) &&
					g.doMove(mv, boxLoc.right.Add(mv)) { // move both squares above
					g.moveBoxV2(boxLoc, mv)
					return true
				}
			}
		}
		return false // box didn't move, so can't move this one
	}

	// move blocks right and left
	if mv == leftVec {
		afterBoxLoc := boxLoc.left.Add(mv)
		if moveable(g.GetLoc(afterBoxLoc)) {
			if g.doMove(mv, afterBoxLoc) {
				g.moveBoxV2(boxLoc, mv)
				return true
			}
		}
		return false // box didn't move, so can't move this one
	}

	if mv == rightVec {
		afterBoxLoc := boxLoc.right.Add(mv)
		if moveable(g.GetLoc(afterBoxLoc)) {
			if g.doMove(mv, afterBoxLoc) {
				g.moveBoxV2(boxLoc, mv)
				return true
			}
		}
		return false // box didn't move, so can't move this one
	}
	// impossible
	panic("did not account for move!")
}

// moveBoxV1 moves the box according to the vector. The spot it is moving to must be empty
// or it will panic.
func (g *Grid) moveBoxV1(current Location, vec MoveVector) {
	next := current.Add(vec)
	if g.GetLoc(next) != Empty {
		panic("trying to move box up/down into an object") // programmer error
	}
	// clear previous spots
	g.data[current.Row][current.Col] = Empty
	g.data[next.Row][next.Col] = Box
}

// moveBoxV2 moves the box according to the vector. The spot it is moving to must be empty
// or it will panic.
func (g *Grid) moveBoxV2(current boxLocation, vec MoveVector) {
	nextLeft := current.left.Add(vec)
	nextRight := current.right.Add(vec)
	switch vec {
	case upVec, downVec:
		// up/down
		if g.GetLoc(nextLeft) != Empty || g.GetLoc(nextRight) != Empty {
			panic("trying to move box up/down into an object") // programmer error
		}
		// clear previous spots
		g.data[current.left.Row][current.left.Col] = Empty
		g.data[current.right.Row][current.right.Col] = Empty
	case rightVec:
		if g.GetLoc(nextRight) != Empty {
			panic("trying to move box right into an object") // programmer error
		}
		g.data[current.left.Row][current.left.Col] = Empty // clear previous spot
	case leftVec:
		if g.GetLoc(nextLeft) != Empty {
			panic("trying to move box left into an object") // programmer error
		}
		g.data[current.right.Row][current.right.Col] = Empty // clear previous spot
	}
	g.data[nextLeft.Row][nextLeft.Col] = BoxLeft
	g.data[nextRight.Row][nextRight.Col] = BoxRight
}

// moveRobot moves the robot. The spot it is moving to must be empty
// or it will panic. Check first.
func (g *Grid) moveRobot(current, next Location) {
	if g.GetLoc(next) != Empty {
		panic("trying to move robot into an object") // programmer error
	}
	g.data[next.Row][next.Col] = Robot
	g.data[current.Row][current.Col] = Empty
	g.robotLoc = next
}
