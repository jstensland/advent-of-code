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
	return strings.NewReader(`###############
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

func TestParseExample(t *testing.T) {
	grid, err := day16.ParseIn(example())

	require.NoError(t, err)
	assert.Equal(t, 15, grid.Width)
	assert.Equal(t, 15, grid.Height)
	assert.Equal(t, day16.Position{13, 1, day16.East}, grid.Start)
	assert.Equal(t, day16.Position{1, 13, day16.North}, grid.End)
}

func TestExampleCost(t *testing.T) {
	grid, err := day16.ParseIn(example())

	require.NoError(t, err)
	assert.Equal(t, 7036, grid.BestRoute())
}

func TestExample2Cost(t *testing.T) {
	grid, err := day16.ParseIn(strings.NewReader(`#################
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
#################`))

	require.NoError(t, err)
	assert.Equal(t, 11048, grid.BestRoute())
}

func TestPart2(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	result, err := day16.SolvePart1(in)

	require.NoError(t, err)
	assert.Equal(t, 99488, result)
}
