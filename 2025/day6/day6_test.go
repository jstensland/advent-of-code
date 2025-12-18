package day6_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2025/day6"
	"github.com/jstensland/advent-of-code/2025/runner"
)

var _ runner.Solver = day6.Part1

func example1() string {
	return `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `
}

func TestPart1(t *testing.T) {
	answer := 5346286649122
	input, err := os.ReadFile("input.txt")
	require.NoError(t, err, "failed to read input.txt")

	result, err := day6.Part1(bytes.NewReader(input))

	require.NoError(t, err)
	if result != answer {
		assert.Equal(t, answer, result)
	}
}

func TestPart1_Example1(t *testing.T) {
	answer := 4277556

	result, err := day6.Part1(bytes.NewReader([]byte(example1())))

	require.NoError(t, err)
	assert.Equal(t, answer, result)
}

func TestPart2(t *testing.T) {
	answer := 0 // TODO: update to answer
	input, err := os.ReadFile("input.txt")
	require.NoError(t, err, "failed to read input.txt")

	result, err := day6.Part2(bytes.NewReader(input))

	require.NoError(t, err)
	if result != answer {
		assert.Equal(t, answer, result)
	}
}

func TestPart2_Example1(t *testing.T) {
	answer := 0 // TODO: update to answer

	result, err := day6.Part2(bytes.NewReader([]byte(example1())))

	require.NoError(t, err)
	assert.Equal(t, answer, result)
}
