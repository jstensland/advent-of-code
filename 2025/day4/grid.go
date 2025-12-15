package day4

import "iter"

// CellState represents the state of a location in the grid.
type CellState int

const (
	Empty CellState = iota
	PaperRoll
)

// Grid represents the parsed input grid where each location can either be a paper roll or empty.
type Grid struct {
	Cells  [][]CellState
	width  int
	height int
}

func (g *Grid) Width() int  { return g.width }
func (g *Grid) Height() int { return g.height }

type position struct {
	row int
	col int
}

// CanMove returns 1 if it's paper and has fewer than 4 other rolls around it. Otherwise, it returns 0.
func (g *Grid) CanMove(pos position) int {
	if g.Cells[pos.row][pos.col] != PaperRoll {
		return 0
	}
	total := 0

	for pos := range g.surroundingPositions(pos) {
		if g.Cells[pos.row][pos.col] == PaperRoll {
			total++
		}
	}
	const adjacentLimit = 4
	if total < adjacentLimit {
		return 1
	}
	return 0
}

func (g *Grid) TryRemoval(pos position) int {
	if g.CanMove(pos) == 0 {
		return 0
	}
	g.Cells[pos.row][pos.col] = Empty
	removed := 1
	for neighbor := range g.surroundingPositions(pos) {
		removed += g.TryRemoval(neighbor)
	}
	return removed
}

func (g *Grid) positions() iter.Seq[position] {
	return func(yield func(r position) bool) {
		for row := range g.height {
			for col := range g.width {
				if !yield(position{row, col}) {
					return
				}
			}
		}
	}
}

// surroundingPositions takes a position on the grid, and returns a slice of the surrounding positions
// it is aware of the edges of the grid.
func (g *Grid) surroundingPositions(pos position) iter.Seq[position] {
	return func(yield func(r position) bool) {
		for i := pos.row - 1; i <= pos.row+1; i++ {
			for j := pos.col - 1; j <= pos.col+1; j++ {
				if i >= 0 && i < g.height && // row on the grid
					j >= 0 && j < g.width && // col on the grid
					(i != pos.row || j != pos.col) { // not the position itself
					if !yield(position{i, j}) {
						return
					}
				}
			}
		}
	}
}
