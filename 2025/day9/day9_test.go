package day9_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2025/day9"
	"github.com/jstensland/advent-of-code/2025/runner"
)

var _ runner.Solver = day9.Part1

func example1() string {
	return `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`
}

func TestPart1_Example1(t *testing.T) {
	answer := 50

	result, err := day9.Part1(bytes.NewReader([]byte(example1())))

	require.NoError(t, err)
	assert.Equal(t, answer, result)
}

func TestPart1(t *testing.T) {
	answer := 4733727792
	input, err := os.ReadFile("input.txt")
	require.NoError(t, err, "failed to read input.txt")

	result, err := day9.Part1(bytes.NewReader(input))

	require.NoError(t, err)
	assert.Equal(t, answer, result)
}

func TestPart2_Example1(t *testing.T) {
	answer := 24

	result, err := day9.Part2(bytes.NewReader([]byte(example1())))

	require.NoError(t, err)
	assert.Equal(t, answer, result)
}

func TestPart2(t *testing.T) {
	answer := 0 // TODO: update to answer
	input, err := os.ReadFile("input.txt")
	require.NoError(t, err, "failed to read input.txt")

	result, err := day9.Part2(bytes.NewReader(input))

	require.NoError(t, err)
	assert.Equal(t, answer, result)
}
