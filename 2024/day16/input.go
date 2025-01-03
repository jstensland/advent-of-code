// Package day16 implements a solution for Day 16 of Advent of Code 2024
package day16

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
)

var ErrUnknownInput = errors.New("unknown input character")

type State int

const (
	Start State = iota
	End
	Wall
	Empty
)

type Orientation int

const numDirections = 4

const (
	North Orientation = iota
	East
	South
	West
)

type Position struct {
	Row       int
	Col       int
	Direction Orientation
}

// Forward will add a vector to a location and returns the resulting location.
func (loc Position) Forward() Position {
	var mv Move
	switch loc.Direction {
	case North:
		mv = Move{deltaEast: 0, deltaSouth: -1}
	case South:
		mv = Move{deltaEast: 0, deltaSouth: 1}
	case East:
		mv = Move{deltaEast: 1, deltaSouth: 0}
	case West:
		mv = Move{deltaEast: -1, deltaSouth: 0}
	}
	return Position{Row: loc.Row + mv.deltaSouth, Col: loc.Col + mv.deltaEast, Direction: loc.Direction}
}

// Right returns the same position, turned to the right.
func (loc Position) Right() Position {
	return Position{Row: loc.Row, Col: loc.Col, Direction: (loc.Direction + 1) % numDirections}
}

// Left returns the same position, turned to the left..
func (loc Position) Left() Position {
	return Position{Row: loc.Row, Col: loc.Col, Direction: ((loc.Direction - 1) + numDirections) % numDirections}
}

type Move struct {
	deltaEast  int
	deltaSouth int
}

type Grid struct {
	data [][]State
	// ReindeerLocation Location
	Start     Position
	End       Position
	Width     int
	Height    int
	leastCost int

	visited map[Position]int
}

func (g *Grid) GetLoc(l Position) State {
	// if you're off the grid some how, it's basically a wall, but maps should
	// avoid this themselves
	if l.Row < 0 || l.Row >= g.Height || l.Col < 0 || l.Col >= g.Width {
		return Wall
	}
	return g.data[l.Row][l.Col]
}

// ParseIn reads the input into a grid.
func ParseIn(in io.Reader) (*Grid, error) {
	data := make([][]State, 0)
	scanner := bufio.NewScanner(in)
	var startCol int
	var startRow int
	var endCol int
	var endRow int

	// Read the grid out of the reader
	for scanner.Scan() {
		inRow := scanner.Text()
		if len(inRow) == 0 {
			break // blank line means we hit the end
		}
		newRow, maybeStartCol, maybeEndCol, err := parseRow(inRow)
		if err != nil {
			return nil, err
		}
		if maybeStartCol != nil {
			startCol = *maybeStartCol
			startRow = len(data)
		}
		if maybeEndCol != nil {
			endCol = *maybeEndCol
			endRow = len(data)
		}
		data = append(data, newRow)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &Grid{
		data:   data,
		Start:  Position{startRow, startCol, East},
		End:    Position{endRow, endCol, North}, // end orientation does not matter
		Width:  len(data[0]),
		Height: len(data),

		visited: make(map[Position]int),
	}, nil
}

func parseRow(inRow string) ([]State, *int, *int, error) {
	var startCol *int
	var endCol *int

	newRow := make([]State, len(inRow))
	for idx, val := range inRow {
		switch val {
		case '\n':
		case '#':
			newRow[idx] = Wall
		case '.':
			newRow[idx] = Empty
		case 'S':
			newRow[idx] = Empty // set first so overwritten for offset = 0
			tmp := idx
			startCol = &tmp
		case 'E':
			newRow[idx] = Empty // set first so overwritten for offset = 0
			tmp := idx
			endCol = &tmp
		default:
			return nil, nil, nil, fmt.Errorf("%s %w", string(val), ErrUnknownInput)
		}
	}
	return newRow, startCol, endCol, nil
}
