package day11_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2024/day11"
)

func exampleEasyIn() io.Reader {
	return strings.NewReader("0 1 10 99 999")
}

func TestPart1ExampleEasy(t *testing.T) {
	stoneLine, err := day11.ParseInput(exampleEasyIn())
	require.NoError(t, err)

	stoneLine.Blink()

	assert.Equal(t, "1 2024 1 0 9 9 2021976", stoneLine.String())
}

func exampleIn() io.Reader {
	return strings.NewReader("125 17")
}

func TestPart1ExampleSteps(t *testing.T) {
	stoneLine, err := day11.ParseInput(exampleIn())
	require.NoError(t, err)

	assert.Equal(t, "125 17", stoneLine.String())
	stoneLine.Blink()
	assert.Equal(t, "253000 1 7", stoneLine.String())
	stoneLine.Blink()
	assert.Equal(t, "253 0 2024 14168", stoneLine.String())
	stoneLine.Blink()
	assert.Equal(t, "512072 1 20 24 28676032", stoneLine.String())
	stoneLine.Blink()
	assert.Equal(t, "512 72 2024 2 0 2 4 2867 6032", stoneLine.String())
	stoneLine.Blink()
	assert.Equal(t, "1036288 7 2 20 24 4048 1 4048 8096 28 67 60 32", stoneLine.String())
	stoneLine.Blink()
	assert.Equal(t, "2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2", stoneLine.String())
}

func TestPart1Example(t *testing.T) {
	answer, err := day11.SolvePart1(exampleIn())

	require.NoError(t, err)
	assert.Equal(t, 55312, answer)
}

func TestPart1Input(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	answer, err := day11.SolvePart1(in)

	require.NoError(t, err)
	assert.Equal(t, 185894, answer) // confirmed
}

func TestPart2Input(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	answer, err := day11.SolvePart2(in)

	require.NoError(t, err)
	assert.Equal(t, 1, answer) // ?
}
