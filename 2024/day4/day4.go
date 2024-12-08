// Package day7 solves AoC 2024 day 4.
package day4

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"log"

	"github.com/jstensland/advent-of-code/2024/runner"
)

func Run(inFile string) error {
	in := runner.Reader(inFile)
	defer in.Close() //nolint:errcheck // no need to check for error

	answer, err := SolvePart1(in)
	if err != nil {
		return err
	}
	fmt.Println("Day 4 part 1:", answer) //nolint:forbidigo // no IO CLI yet

	in2 := runner.Reader(inFile)
	defer in2.Close() //nolint:errcheck // no need to check for error
	answer, err = SolvePart2(in2)
	if err != nil {
		return err
	}
	fmt.Println("Day 4 part 2:", answer) //nolint:forbidigo // no IO CLI yet
	return nil
}

// SolvePart1 finds occurances of XMAS in a wordsearch fashion.
//
// It's a small input, so parse the whole thing into memory
// and then search for XMAS from every X, checking each of the 8
// possible directions.
func SolvePart1(in io.Reader) (int, error) {
	grid, err := ParseGrid(in)
	if err != nil {
		return 0, fmt.Errorf("error loading grid: %w", err)
	}

	return grid.XmasCount(), nil
}

func SolvePart2(in io.Reader) (int, error) {
	grid, err := ParseGrid(in)
	if err != nil {
		return 0, fmt.Errorf("error loading grid: %w", err)
	}

	return grid.XmasCount2(), nil
}

type Grid struct {
	data   [][]rune
	width  int
	height int
}

func ParseGrid(in io.Reader) (Grid, error) {
	scanner := bufio.NewScanner(in)
	grid := Grid{[][]rune{}, 0, 0}
	for scanner.Scan() {
		row := []rune(scanner.Text())
		grid.data = append(grid.data, row)
		grid.height++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	grid.width = len(grid.data[0])
	return grid, nil
}

func (g Grid) Width() int  { return g.width }
func (g Grid) Height() int { return g.height }

type position struct {
	row int
	col int
}

func (g Grid) positions() iter.Seq[position] {
	return func(yield func(r position) bool) {
		for row := 0; row < g.height; row++ {
			for col := 0; col < g.width; col++ {
				if !yield(position{row, col}) {
					return
				}
			}
		}
	}
}

// XmasCount counts the number of XMAS in a row.
// Each X could have up to 8 XMAS.
func (g Grid) XmasCount() int {
	count := 0
	for loc := range g.positions() {
		// fmt.Printf("looking at row %v and column %v\n", loc.row, loc.col)
		count += g.checkPosition(loc)
	}

	return count
}

// TODO: write specific tests for each to avoid print statement visual parsing
// for issues

func (g Grid) checkPosition(loc position) int {
	// Check if it's an X. If not, continue.
	if g.data[loc.row][loc.col] != 'X' {
		return 0
	}

	total := 0
	if g.checkUp(loc) {
		// fmt.Println("matches up:", loc)
		total++
	}
	if g.checkRightUp(loc) {
		// fmt.Println("matches right up:", loc)
		total++
	}
	if g.checkRight(loc) {
		// fmt.Println("matches right:", loc)
		total++
	}
	if g.checkRightDown(loc) {
		// fmt.Println("matches right down:", loc)
		total++
	}
	if g.checkDown(loc) {
		// fmt.Println("matches down:", loc)
		total++
	}
	if g.checkLeftDown(loc) {
		// fmt.Println("matches left down:", loc)
		total++
	}
	if g.checkLeft(loc) {
		// fmt.Println("matches left:", loc)
		total++
	}
	if g.checkLeftUp(loc) {
		// fmt.Println("matches left up:", loc)
		total++
	}
	return total
}

// These should be reduced... the same function could take
// the loc, the word, and the direction...

func (g Grid) checkRight(loc position) bool {
	if loc.col+3 >= g.width {
		return false
	}
	return g.data[loc.row][loc.col+1] == 'M' &&
		g.data[loc.row][loc.col+2] == 'A' &&
		g.data[loc.row][loc.col+3] == 'S'
}

func (g Grid) checkRightUp(loc position) bool {
	if loc.col+3 >= g.width || loc.row-3 < 0 {
		return false
	}
	return g.data[loc.row-1][loc.col+1] == 'M' &&
		g.data[loc.row-2][loc.col+2] == 'A' &&
		g.data[loc.row-3][loc.col+3] == 'S'
}

func (g Grid) checkRightDown(loc position) bool {
	if loc.col+3 >= g.width || loc.row+3 >= g.height {
		return false
	}
	return g.data[loc.row+1][loc.col+1] == 'M' &&
		g.data[loc.row+2][loc.col+2] == 'A' &&
		g.data[loc.row+3][loc.col+3] == 'S'
}

func (g Grid) checkUp(loc position) bool {
	if loc.row-3 < 0 {
		return false
	}
	return g.data[loc.row-1][loc.col] == 'M' &&
		g.data[loc.row-2][loc.col] == 'A' &&
		g.data[loc.row-3][loc.col] == 'S'
}

func (g Grid) checkDown(loc position) bool {
	if loc.row+3 >= g.height {
		return false
	}
	return g.data[loc.row+1][loc.col] == 'M' &&
		g.data[loc.row+2][loc.col] == 'A' &&
		g.data[loc.row+3][loc.col] == 'S'
}

func (g Grid) checkLeftDown(loc position) bool {
	if loc.col-3 < 0 || loc.row+3 >= g.height {
		return false
	}
	return g.data[loc.row+1][loc.col-1] == 'M' &&
		g.data[loc.row+2][loc.col-2] == 'A' &&
		g.data[loc.row+3][loc.col-3] == 'S'
}

func (g Grid) checkLeft(loc position) bool {
	if loc.col-3 < 0 {
		return false
	}
	return g.data[loc.row][loc.col-1] == 'M' &&
		g.data[loc.row][loc.col-2] == 'A' &&
		g.data[loc.row][loc.col-3] == 'S'
}

func (g Grid) checkLeftUp(loc position) bool {
	if loc.row-3 < 0 || loc.col-3 < 0 {
		return false
	}
	return g.data[loc.row-1][loc.col-1] == 'M' &&
		g.data[loc.row-2][loc.col-2] == 'A' &&
		g.data[loc.row-3][loc.col-3] == 'S'
}

// XmasCount2 looks for these shapes
// M.S
// .A.
// M.S
func (g Grid) XmasCount2() int {
	count := 0
	for loc := range g.positions() {
		// fmt.Printf("looking at row %v and column %v\n", loc.row, loc.col)
		count += g.checkPosition2(loc)
	}

	return count
}

func (g Grid) checkPosition2(loc position) int {
	// Orient around A characters. If it's not an A, return false
	if g.data[loc.row][loc.col] != 'A' {
		return 0
	}
	// if it's on the edge, return false
	if loc.col == 0 || loc.row == 0 || loc.col == g.width-1 || loc.row == g.height-1 {
		return 0
	}

	// There are 4 orientations, which are all rotations
	// each A either is or is not a an XMAS

	// M.S
	// .A.
	// M.S

	// S.S
	// .A.
	// M.M

	// S.M
	// .A.
	// S.M

	// M.M
	// .A.
	// S.S

	// just check that corner are opposites
	upperLeft := g.data[loc.row-1][loc.col-1]
	upperRight := g.data[loc.row-1][loc.col+1]
	lowerLeft := g.data[loc.row+1][loc.col-1]
	lowerRight := g.data[loc.row+1][loc.col+1]

	if oppositeSandM(upperLeft, lowerRight) && oppositeSandM(lowerLeft, upperRight) {
		return 1
	}
	return 0
}

func oppositeSandM(x, y rune) bool {
	if x == 'S' {
		return y == 'M'
	}
	if x == 'M' {
		return y == 'S'
	}
	return false
}
