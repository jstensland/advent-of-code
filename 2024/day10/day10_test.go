package day10_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2024/day10"
)

func exampleTrivialIn() io.Reader {
	return strings.NewReader(
		`0123
0654
0789`)
}

func TestPart1ExampleTrivial(t *testing.T) {
	answer, err := day10.SolvePart1(exampleTrivialIn())

	require.NoError(t, err)
	assert.Equal(t, 1, answer)
}

func exampleEasyIn() io.Reader {
	return strings.NewReader(
		`7770777
7771777
7772777
6543456
7777777
8777778
9777779`)
}

func TestPart1ExampleEasy(t *testing.T) {
	answer, err := day10.SolvePart1(exampleEasyIn())

	require.NoError(t, err)
	assert.Equal(t, 2, answer)
}

func exampleIn() io.Reader {
	return strings.NewReader(
		`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`)
}

func TestPart1Example(t *testing.T) {
	answer, err := day10.SolvePart1(exampleIn())

	require.NoError(t, err)
	assert.Equal(t, 36, answer)
}

func TestPart1Input(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	answer, err := day10.SolvePart1(in)

	require.NoError(t, err)
	assert.Equal(t, 472, answer) // confirmed
}

func TestPart2EasyMap(t *testing.T) {
	in := strings.NewReader(
		`.....0.
..4321.
..5..2.
..6543.
..7..4.
..8765.
..9....
`)

	answer, err := day10.SolvePart2(in)

	require.NoError(t, err)
	assert.Equal(t, 3, answer)
}

func TestPart2MediumMap(t *testing.T) {
	in := strings.NewReader(
		`012345
123456
234567
345678
4.6789
56789.`)

	answer, err := day10.SolvePart2(in)

	require.NoError(t, err)
	assert.Equal(t, 227, answer)
}

// slow to run with `-race` and coverage, but otherwise fast
func TestPart2Input(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	answer, err := day10.SolvePart2(in)

	require.NoError(t, err)
	assert.Equal(t, 969, answer) // confirmed
}
