package day7

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// CellState represents the state of a location in the grid.
type CellState rune

const (
	Start    CellState = 'S'
	Empty    CellState = '.'
	Splitter CellState = '^'
	Beam     CellState = '|'
)

// Grid represents the parsed input, and its evolution.
type Grid struct {
	grid       [][]CellState
	width      int
	height     int
	iteration  int
	splitCount int
}

func (g *Grid) Width() int  { return g.width }
func (g *Grid) Height() int { return g.height }
func (g *Grid) String() string {
	var sb strings.Builder
	for _, row := range g.grid {
		sb.WriteString(string(row))
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (g *Grid) SplitCount() int {
	return g.splitCount
}

func ParseIn(r io.Reader) (*Grid, error) {
	cells := [][]CellState{}
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		cells = append(cells, []CellState(line))
	}

	return &Grid{
		grid:      cells,
		width:     len(cells[0]),
		height:    len(cells),
		iteration: 0,
	}, nil
}

// cache for memoization
//
//nolint:gochecknoglobals // cache for memoization
var cache = map[string]int{}

func (g *Grid) ProgressTimeline(rowIdx, colIdx int) int {
	key := strconv.Itoa(rowIdx) + "_" + strconv.Itoa(colIdx)
	if val, ok := cache[key]; ok {
		return val
	}
	// base cases
	// if we're out of rows, return 1
	if rowIdx == g.height-1 {
		return 1
	}
	// column idx shouldn't be able to go off the grid. skipping condition

	if g.grid[rowIdx][colIdx] == Empty {
		// add nothing, no split, just keep going
		answer := g.ProgressTimeline(rowIdx+1, colIdx)
		cache[key] = answer
		return answer
	}

	if g.grid[rowIdx][colIdx] == Splitter {
		// return the addition of the add the number of possibility on the right path to
		// the number of possibilities on the left
		lanswer := g.ProgressTimeline(rowIdx+1, colIdx-1)
		cache[key] = lanswer
		ranswer := g.ProgressTimeline(rowIdx+1, colIdx+1)
		cache[key] = ranswer
		return lanswer + ranswer
	}
	panic(fmt.Sprintf("AHHH what did I hit?! %v", g.grid[rowIdx][colIdx]))
}

func (g *Grid) Start() (int, int) {
	for idx, cell := range g.grid[0] {
		if cell == Start {
			return 0, idx
		}
	}
	panic("no start in the first row!")
}

func (g *Grid) Progress() {
	if g.iteration == g.height-1 {
		return
	}
	for colIdx, cell := range g.grid[g.iteration] {
		switch cell {
		case Start, Beam:
			g.advance(colIdx)
		case Splitter:
			g.clear(colIdx)
		case Empty:
		}
	}
	g.iteration++
}

func (g *Grid) advance(col int) {
	switch g.grid[g.iteration+1][col] {
	case Empty, Start, Beam:
		g.grid[g.iteration+1][col] = Beam
	case Splitter:
		if col-1 >= 0 {
			g.grid[g.iteration+1][col-1] = Beam
		}
		if col+1 < g.width {
			g.grid[g.iteration+1][col+1] = Beam
		}
		g.splitCount++
	default:
		panic("uh oh, how did I get here?")
	}
}

func (g *Grid) clear(col int) {
	g.grid[g.iteration+1][col] = Empty
}
