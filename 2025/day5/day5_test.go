package day5_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2025/day5"
	"github.com/jstensland/advent-of-code/2025/runner"
)

var _ runner.Solver = day5.Part1

func example1() string {
	return `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
}

func TestPart1(t *testing.T) {
	answer := 775
	input, err := os.ReadFile("input.txt")
	require.NoError(t, err, "failed to read input.txt")

	result, err := day5.Part1(bytes.NewReader(input))

	require.NoError(t, err)
	if result != answer {
		assert.Equal(t, answer, result)
	}
}

func TestPart1_Example1(t *testing.T) {
	answer := 3

	result, err := day5.Part1(bytes.NewReader([]byte(example1())))

	require.NoError(t, err)
	assert.Equal(t, answer, result)
}

func TestPart2(t *testing.T) {
	answer := 350684792662845
	input, err := os.ReadFile("input.txt")
	require.NoError(t, err, "failed to read input.txt")

	result, err := day5.Part2(bytes.NewReader(input))

	require.NoError(t, err)
	if result != answer {
		assert.Equal(t, answer, result)
	}
}

func TestPart2_Example1(t *testing.T) {
	answer := 14

	result, err := day5.Part2(bytes.NewReader([]byte(example1())))

	require.NoError(t, err)
	assert.Equal(t, answer, result)
}
