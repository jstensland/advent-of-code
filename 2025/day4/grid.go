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

// CanMove returns 1 if it's paper and has fewer than 4 other rolls around it.
// Otherwise, it returns 0
func (g *Grid) CanMove(pos position) int {
	if g.Cells[pos.row][pos.col] != PaperRoll {
		return 0
	}
	total := 0

	for _, pos := range g.surroundingPositions(pos) {
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

func (g *Grid) surroundingPositions(pos position) []position {
	out := []position{}

	// top row
	if pos.row > 0 {
		if pos.col > 0 {
			out = append(out, position{pos.row - 1, pos.col - 1})
		}
		out = append(out, position{pos.row - 1, pos.col})
		if pos.col+1 < g.width {
			out = append(out, position{pos.row - 1, pos.col + 1})
		}
	}
	// middle row
	if pos.col > 0 {
		out = append(out, position{pos.row, pos.col - 1})
	}
	if pos.col+1 < g.width {
		out = append(out, position{pos.row, pos.col + 1})
	}
	// lower row
	if pos.row+1 < g.height {
		if pos.col > 0 {
			out = append(out, position{pos.row + 1, pos.col - 1})
		}
		out = append(out, position{pos.row + 1, pos.col})
		if pos.col+1 < g.width {
			out = append(out, position{pos.row + 1, pos.col + 1})
		}
	}
	return out
}
