package day9_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2024/day9"
)

func exampleIn() io.Reader {
	return strings.NewReader("2333133121414131402")
}

func TestPart1Steps(t *testing.T) {
	blocks, err := day9.ParseInput(exampleIn())

	require.NoError(t, err)
	assert.Equal(t, "00...111...2...333.44.5555.6666.777.888899", blocks.String())

	blocks.MoveFileSegments()
	assert.Equal(t, "0099811188827773336446555566..............", blocks.String())
	assert.Equal(t, 1928, blocks.CheckSum()) // confirmed
}

func TestPart1Example(t *testing.T) {
	answer, err := day9.SolvePart1(exampleIn())

	require.NoError(t, err)
	assert.Equal(t, 1928, answer)
}

func TestPart1Input(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	answer, err := day9.SolvePart1(in)

	require.NoError(t, err)
	assert.Equal(t, 6337921897505, answer) // confirmed
}

func TestPart2Steps(t *testing.T) {
	blocks, err := day9.ParseInput(exampleIn())
	require.NoError(t, err)
	assert.Equal(t, "00...111...2...333.44.5555.6666.777.888899", blocks.String())

	blocks.MoveFiles()

	assert.Equal(t, "00992111777.44.333....5555.6666.....8888..", blocks.String())
	assert.Equal(t, 2858, blocks.CheckSum()) // confirmed
}

func TestPart2Example(t *testing.T) {
	answer, err := day9.SolvePart2(exampleIn())

	require.NoError(t, err)
	assert.Equal(t, 2858, answer)
}

// slow to run with `-race` and coverage, but otherwise fast
func TestPart2Input(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	answer, err := day9.SolvePart2(in)

	require.NoError(t, err)
	assert.Equal(t, 6362722604045, answer) // confirmed
}
