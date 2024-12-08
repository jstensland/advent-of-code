package day6_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2024/day6"
	"github.com/jstensland/advent-of-code/2024/runner"
)

func TestPart1Input(t *testing.T) {
	// inFile := "./input.txt"
	// in, err := os.Open(inFile)
	// require.NoError(t, err)

	// answer, err := day6.RunPart1(in)
	answer, err := day6.SolvePart1(runner.Reader("./input.txt"))

	require.NoError(t, err)
	assert.Equal(t, 4903, answer) // confirmed
}

// too slow...
func TestPart2Input(t *testing.T) {
	answer, err := day6.SolvePart2(runner.Reader("./input.txt"))

	require.NoError(t, err)
	assert.Equal(t, 1911, answer)
}

func exampleIn() io.Reader {
	return strings.NewReader(
		`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`)
}

func TestPart1Example(t *testing.T) {
	answer, err := day6.SolvePart1(exampleIn())

	require.NoError(t, err)
	assert.Equal(t, 41, answer)
}

func TestPart2Example(t *testing.T) {
	answer, err := day6.SolvePart2(exampleIn())

	require.NoError(t, err)
	assert.Equal(t, 6, answer)
}

func TestExampleGuardLoc(t *testing.T) {
	layout, err := day6.ParseInput(exampleIn())

	require.NoError(t, err)
	assert.Equal(t, day6.Location{6, 4}, layout.GuardLocation())
}

func TestTurn(t *testing.T) {
	assert.Equal(t, day6.Right, day6.Turn(day6.Up))
	assert.Equal(t, day6.Down, day6.Turn(day6.Right))
	assert.Equal(t, day6.Left, day6.Turn(day6.Down))
	assert.Equal(t, day6.Up, day6.Turn(day6.Left))
}
