package day8_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2024/day8"
)

func TestPart1Input(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	answer, err := day8.SolvePart1(in)

	require.NoError(t, err)
	assert.Equal(t, 396, answer) // confirmed
}

func exampleIn() io.Reader {
	return strings.NewReader(`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`)
}

func TestPart1Example(t *testing.T) {
	answer, err := day8.SolvePart1(exampleIn())

	require.NoError(t, err)
	assert.Equal(t, 14, answer)
}

func example2In() io.Reader {
	return strings.NewReader(
		`T.........
...T......
.T........
..........
..........
..........
..........
..........
..........
..........`)
}

func TestPart2Input(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	answer, err := day8.SolvePart2(in)

	require.NoError(t, err)
	assert.Equal(t, 1200, answer) //
}

func TestPart2Example1(t *testing.T) {
	answer, err := day8.SolvePart2(exampleIn())

	require.NoError(t, err)
	assert.Equal(t, 34, answer)
}

func TestPart2Example2(t *testing.T) {
	answer, err := day8.SolvePart2(example2In())

	require.NoError(t, err)
	assert.Equal(t, 9, answer)
}
