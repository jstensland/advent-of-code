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

func (g *Grid) maybeMoveDoubleBox(currentLoc Location, mv MoveVector) bool {
	currentVal := g.GetLoc(currentLoc)
	// now for the boxes...
	var boxLoc boxLocation
	if currentVal == BoxLeft {
		boxLoc = boxLocation{left: currentLoc, right: Location{Row: currentLoc.Row, Col: currentLoc.Col + 1}}
	} else {
		boxLoc = boxLocation{left: Location{Row: currentLoc.Row, Col: currentLoc.Col - 1}, right: currentLoc}
	}

	// move block
	if mv == upVec || mv == downVec {
		return g.maybeMoveDoubleBoxVert(boxLoc, mv)
	}

	// move blocks right and left
	var afterBoxLoc Location
	if mv == leftVec {
		afterBoxLoc = boxLoc.left.Add(mv)
	} else {
		afterBoxLoc = boxLoc.right.Add(mv)
	}
	if g.doMove(mv, afterBoxLoc) {
		g.moveBoxV2(boxLoc, mv)
		return true
	}
	return false // box didn't move, so can't move this one
}

func (g *Grid) maybeMoveDoubleBoxVert(boxLoc boxLocation, mv MoveVector) bool {
	nextLeft := boxLoc.left.Add(mv)
	nextRight := boxLoc.right.Add(mv)
	// if the value directly above the left side of the box is the left side of another box, boxes
	// are aligned and there is only one to move
	if g.GetLoc(nextLeft) == BoxLeft {
		if g.doMove(mv, nextLeft) { // move the one box
			g.moveBoxV2(boxLoc, mv)
			return true
		}
		return false
	}
	// If there are too possible pieces that could move. Must check that they can BOTH move before
	// moving either one.
	if doable(g, mv, boxLoc) && g.doMove(mv, nextLeft) && g.doMove(mv, nextRight) {
		g.moveBoxV2(boxLoc, mv)
		return true
	}
	return false // box didn't move, so can't move this one
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
		if g.doMove(mv, nextLoc) { // if next value could move, try to move it
			g.moveRobot(currentLoc, nextLoc) // it moved, so move the robot now
			return true
		}
		return false // not movable, or didn't move
	}

	return g.maybeMoveDoubleBox(currentLoc, mv)
}

// doable makes a copy of the grid and applies the move, returning if it worked.
//
// IMPROVEMENT: it could just check, rather than do, which would avoid the copy, but would
// need recursive logic similar to doMove
func doable(g *Grid, mv MoveVector, boxLoc boxLocation) bool {
	dup := g.Copy()
	return dup.doMove(mv, boxLoc.left.Add(mv)) && dup.doMove(mv, boxLoc.right.Add(mv))
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
