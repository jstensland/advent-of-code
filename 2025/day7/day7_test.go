package day7_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2025/day7"
	"github.com/jstensland/advent-of-code/2025/runner"
)

var _ runner.Solver = day7.Part1

func example1() string {
	return `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`
}

func TestParseIn(t *testing.T) {
	r := bytes.NewReader([]byte(example1()))

	grid, err := day7.ParseIn(r)

	require.NoError(t, err)
	assert.Equal(t, 15, grid.Width())
	assert.Equal(t, 16, grid.Height())

	assert.Equal(t, example1()+"\n", grid.String())
}

func TestPart1(t *testing.T) {
	answer := 1658
	input, err := os.ReadFile("input.txt")
	require.NoError(t, err, "failed to read input.txt")

	result, err := day7.Part1(bytes.NewReader(input))

	require.NoError(t, err)
	if result != answer {
		assert.Equal(t, answer, result)
	}
}

func TestPart1_Example1(t *testing.T) {
	answer := 21
	result, err := day7.Part1(bytes.NewReader([]byte(example1())))

	require.NoError(t, err)
	assert.Equal(t, answer, result)
}

func TestPart2(t *testing.T) {
	answer := 0 // TODO: update to answer
	input, err := os.ReadFile("input.txt")
	require.NoError(t, err, "failed to read input.txt")

	result, err := day7.Part2(bytes.NewReader(input))

	require.NoError(t, err)
	if result != answer {
		assert.Equal(t, answer, result)
	}
}

func TestPart2_Example1(t *testing.T) {
	answer := 0 // TODO: update to answer

	result, err := day7.Part2(bytes.NewReader([]byte(example1())))

	require.NoError(t, err)
	assert.Equal(t, answer, result)
}
