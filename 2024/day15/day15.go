package day15

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
)

var ErrUnknownInput = errors.New("unknown input character")

func SolvePart1(in io.Reader) (int, error) {
	grid, err := ParseIn(in, 1, 0)
	if err != nil {
		return 0, fmt.Errorf("error loading input: %w", err)
	}
	// fmt.Println(grid)
	grid.RunRobots()
	// fmt.Println(grid)
	return grid.TotalGPS(), nil
}

type State int

func (s State) String() string {
	return map[State]string{
		Robot:    "@",
		Wall:     "#",
		Box:      "O",
		BoxLeft:  "[",
		BoxRight: "]",
		Empty:    ".",
	}[s]
}

const (
	Robot State = iota
	Wall
	Box
	BoxLeft  // for wide boxes
	BoxRight // for wide boxes
	Empty
)

type Grid struct {
	data     [][]State
	robotLoc Location
	movesVec []MoveVector
	Width    int
	Height   int
}

func (g *Grid) GetLoc(l Location) State {
	return g.data[l.Row][l.Col]
}

type MoveVector struct {
	deltaRow int
	deltaCol int
}

//nolint:gochecknoglobals // reference sentinel values for possilbe moves
var (
	upVec    = MoveVector{-1, 0}
	rightVec = MoveVector{0, 1}
	downVec  = MoveVector{1, 0}
	leftVec  = MoveVector{0, -1}
)

func (mv MoveVector) String() string {
	return map[MoveVector]string{
		upVec:    "^",
		rightVec: ">",
		downVec:  "v",
		leftVec:  "<",
	}[mv]
}

func (g *Grid) String() string {
	out := ""
	for row := range g.Height {
		for col := range g.Width {
			out += g.data[row][col].String()
		}
		out += "\n"
	}
	return out
}

// RunRobots moves the robot all the moves. affecting the grid
func (g *Grid) RunRobots() {
	for _, mv := range g.movesVec {
		// fmt.Println(g)
		// fmt.Println(mv)
		g.maybeMoveV1(mv)
	}
}

func (g *Grid) maybeMoveV1(mv MoveVector) {
	g.doMoveV1(mv, g.robotLoc)
}

// doMoveV1 checks if the given move is possible from that location. If it is, it first calls itself
// to move any movable object out of the way as necessary, and then does it's own move.
//
// Blocks are checked and moved as a unit.
//
// The return value indicates if the move was possible.
func (g *Grid) doMoveV1(mv MoveVector, currentLoc Location) bool {
	currentVal := g.GetLoc(currentLoc)
	if currentVal == Wall {
		return false // can't move a wall
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
		// fmt.Println("start with robot move")
		// move robot. robot one wide
		if g.doMoveV1(mv, nextLoc) { // if next value could move, try to move it
			g.moveRobot(currentLoc, nextLoc) // it moved, so move the robot now
			return true
		}
		return false // not movable, or didn't move
	}

	// now for the boxes...
	if currentVal == Box {
		if g.doMoveV1(mv, nextLoc) { // move the next box
			g.moveBoxV1(currentLoc, mv)
			return true
		}
		return false // box didn't move, so can't move this one
	}
	// impossible
	panic(fmt.Sprintf("did not account for move! %v", currentVal))
}

// TotalGPS calculates the total "GPS" of all the botxes
func (g *Grid) TotalGPS() int {
	// fmt.Println(g)
	total := 0
	// for each box
	for row := range g.Height {
		for col := range g.Width {
			if g.data[row][col] == Box {
				total += 100*row + col //nolint:mnd // magic number
			}
		}
	}
	return total
}

// ParseIn reads the input into a grid of robots
func ParseIn(in io.Reader, widthFactor, offset int) (*Grid, error) {
	grid := make([][]State, 0)
	var robotRow int
	var robotCol int
	scanner := bufio.NewScanner(in)

	// Read the grid out of the reader
	for scanner.Scan() {
		inRow := scanner.Text()
		if len(inRow) == 0 {
			break // blank line indicates we're moving to the robot moves
		}
		newRow, maybeCol, err := parseRow(inRow, widthFactor, offset)
		if err != nil {
			return nil, err
		}
		if maybeCol != nil {
			robotCol = *maybeCol
			robotRow = len(grid)
		}
		grid = append(grid, newRow)
	}

	// Read the robot moves in
	robotVectMoves, err := parseMoves(scanner)
	if err != nil {
		return nil, err
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &Grid{
		data:     grid,
		robotLoc: Location{robotRow, robotCol},
		movesVec: robotVectMoves,
		Width:    len(grid[0]),
		Height:   len(grid),
	}, nil
}

func parseRow(inRow string, widthFactor, offset int) ([]State, *int, error) {
	var robotCol *int

	newRow := make([]State, len(inRow)*widthFactor)
	for idx, val := range inRow {
		switch val {
		case '\n':
		case '#':
			newRow[idx*widthFactor] = Wall
			newRow[idx*widthFactor+offset] = Wall
		case '.':
			newRow[idx*widthFactor] = Empty
			newRow[idx*widthFactor+offset] = Empty
		case '@':
			newRow[idx*widthFactor+offset] = Empty // set first so overwritten for offset = 0
			newRow[idx*widthFactor] = Robot
			tmp := idx * widthFactor
			robotCol = &tmp
		case 'O':
			if widthFactor == 1 {
				newRow[idx*widthFactor] = Box
			} else {
				newRow[idx*widthFactor] = BoxLeft
				// IMPROVEMENT not handling middle parts of the box right now for bigger factors
				// restrict factor to 1 or 2 if that's all I support!
				newRow[idx*widthFactor+offset] = BoxRight
			}
		default:
			return nil, nil, fmt.Errorf("%s %w", string(val), ErrUnknownInput)
		}
	}
	return newRow, robotCol, nil
}

// parseMoves reads the robot moves from the rest of the scanner
func parseMoves(scanner *bufio.Scanner) ([]MoveVector, error) {
	// Read the robot moves in
	moveVecMap := map[rune]MoveVector{
		'^': upVec,
		'>': rightVec,
		'<': leftVec,
		'v': downVec,
	}
	robotMoveVector := []MoveVector{}
	for scanner.Scan() {
		allMoves := scanner.Text()
		for _, chr := range allMoves {
			mvVec, ok := moveVecMap[chr]
			if !ok {
				return nil, fmt.Errorf("unexpected robot move %v: %w", chr, ErrUnknownInput)
			}
			robotMoveVector = append(robotMoveVector, mvVec)
		}
	}
	return robotMoveVector, nil
}
