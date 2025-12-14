package day4_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2025/day4"
	"github.com/jstensland/advent-of-code/2025/runner"
)

var _ runner.Solver = day4.Part1

func example1() string {
	return `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
}

func example1_mapped() string {
	return `..xx.xx@x.
x@@.@.@.@@
@@@@@.x.@@
@.@@@@..@.
x@.@@@@.@x
.@@@@@@@.@
.@.@.@.@@@
x.@@@.@@@@
.@@@@@@@@.
x.x.@@@.x.`
}

func TestParseIn(t *testing.T) {
	grid, err := day4.ParseIn(bytes.NewReader([]byte(example1())))
	require.NoError(t, err)

	// Check dimensions
	assert.Equal(t, 10, grid.Width, "expected width of 10")
	assert.Equal(t, 10, grid.Height, "expected height of 10")

	// Check some sample cells
	assert.Equal(t, day4.Empty, grid.Cells[0][0], "position (0,0) should be Empty")
	assert.Equal(t, day4.Empty, grid.Cells[0][1], "position (0,1) should be Empty")
	assert.Equal(t, day4.PaperRoll, grid.Cells[0][2], "position (0,2) should be PaperRoll")
	assert.Equal(t, day4.PaperRoll, grid.Cells[1][0], "position (1,0) should be PaperRoll")
	assert.Equal(t, day4.Empty, grid.Cells[1][3], "position (1,3) should be Empty")
	assert.Equal(t, day4.PaperRoll, grid.Cells[9][0], "position (9,0) should be PaperRoll")
	assert.Equal(t, day4.Empty, grid.Cells[9][9], "position (9,9) should be Empty")
}

func TestPart1(t *testing.T) {
	answer := 0 // TODO: update to answer
	input, err := os.ReadFile("input.txt")
	require.NoError(t, err, "failed to read input.txt")

	result, err := day4.Part1(bytes.NewReader(input))

	require.NoError(t, err)
	if result != answer {
		assert.Equal(t, result, answer)
	}
}

func TestPart1_Example1(t *testing.T) {
	answer := 13

	result, err := day4.Part1(bytes.NewReader([]byte(example1())))

	require.NoError(t, err)
	assert.Equal(t, answer, result)
}

func TestPart2(t *testing.T) {
	answer := 0 // TODO: update to answer
	input, err := os.ReadFile("input.txt")
	require.NoError(t, err, "failed to read input.txt")

	result, err := day4.Part2(bytes.NewReader(input))

	require.NoError(t, err)
	if result != answer {
		assert.Equal(t, answer, result)
	}
}

func TestPart2_Example1(t *testing.T) {
	answer := 0 // TODO: update to answer

	result, err := day4.Part2(bytes.NewReader([]byte(example1())))

	require.NoError(t, err)
	assert.Equal(t, answer, result)
}
