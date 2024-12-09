package day1_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2024/day1"
)

func TestPart1Input(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	answer, err := day1.SolvePart1(in)

	require.NoError(t, err)
	assert.Equal(t, 1646452, answer) // confirmed
}

func exampleIn() io.Reader {
	return strings.NewReader(`3   4
4   3
2   5
1   3
3   9
3   3`)
}

func TestPart1Example(t *testing.T) {
	answer, err := day1.SolvePart1(exampleIn())

	require.NoError(t, err)
	assert.Equal(t, 11, answer)
}

func TestPart2Input(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	answer, err := day1.SolvePart2(in)

	require.NoError(t, err)
	assert.Equal(t, 23609874, answer) // confirmed
}

func TestPart2Example1(t *testing.T) {
	answer, err := day1.SolvePart2(exampleIn())

	require.NoError(t, err)
	assert.Equal(t, 31, answer)
}
