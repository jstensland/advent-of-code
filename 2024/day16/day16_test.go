package day16_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2024/day16"
)

func example() io.Reader {
	return strings.NewReader(
		`###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`)
}

func example2() io.Reader {
	return strings.NewReader(`#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`)
}

func TestParseExample(t *testing.T) {
	grid, err := day16.ParseIn(example())

	require.NoError(t, err)
	assert.Equal(t, 15, grid.Width)
	assert.Equal(t, 15, grid.Height)
	assert.Equal(t, day16.Position{day16.Location{13, 1}, day16.East}, grid.Start)
	assert.Equal(t, day16.Position{day16.Location{1, 13}, day16.North}, grid.End)
}

func TestExampleCost(t *testing.T) {
	grid, err := day16.ParseIn(example())

	require.NoError(t, err)
	assert.Equal(t, 7036, grid.BestRoute())
}

func TestExample2Cost(t *testing.T) {
	grid, err := day16.ParseIn(example2())

	require.NoError(t, err)
	assert.Equal(t, 11048, grid.BestRoute())
}

func TestPart1(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	result, err := day16.SolvePart1(in)

	require.NoError(t, err)
	assert.Equal(t, 99488, result)
}

func TestBestSeatsExample(t *testing.T) {
	grid, err := day16.ParseIn(example())

	require.NoError(t, err)
	assert.Equal(t, 45, grid.BestSeats())
}

func TestBestSeatsExample2(t *testing.T) {
	grid, err := day16.ParseIn(example2())

	require.NoError(t, err)
	assert.Equal(t, 64, grid.BestSeats())
}

func TestPart2(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	result, err := day16.SolvePart2(in)

	require.NoError(t, err)
	assert.Equal(t, 516, result)
}
