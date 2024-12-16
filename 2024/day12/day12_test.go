package day12_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2024/day12"
)

func exampleTrivial() io.Reader {
	return strings.NewReader(`AAAA
BBCD
BBCC
EEEC`)
}

func TestRunPart1ExampleTrivial(t *testing.T) {
	total, err := day12.SolvePart1(exampleTrivial())

	require.NoError(t, err)
	assert.Equal(t, 140, total)
}

func exampleInner() io.Reader {
	return strings.NewReader(`OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`)
}

func TestRunPart1ExampleInner(t *testing.T) {
	total, err := day12.SolvePart1(exampleInner())

	require.NoError(t, err)
	assert.Equal(t, 772, total)
}

func example() io.Reader {
	return strings.NewReader(`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`)
}

func TestRunPart1Example(t *testing.T) {
	total, err := day12.SolvePart1(example())

	require.NoError(t, err)
	assert.Equal(t, 1930, total)
}

func TestRunPart1(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	total, err := day12.SolvePart1(in)

	require.NoError(t, err)
	assert.Equal(t, 1431316, total) // confirmed
}

func TestRunPart2ExampleTrivial(t *testing.T) {
	total, err := day12.SolvePart2(exampleTrivial())

	require.NoError(t, err)
	assert.Equal(t, 80, total)
}

func TestRunPart2ExampleInner(t *testing.T) {
	total, err := day12.SolvePart2(exampleInner())

	require.NoError(t, err)
	assert.Equal(t, 436, total)
}

func TestRunPart2Example(t *testing.T) {
	total, err := day12.SolvePart2(example())

	require.NoError(t, err)
	assert.Equal(t, 1206, total)
}

func TestRunPart2(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	total, err := day12.SolvePart2(in)

	require.NoError(t, err)
	assert.Equal(t, 821428, total) // ?
}
