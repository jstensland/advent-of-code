package day8_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2025/day8"
	"github.com/jstensland/advent-of-code/2025/runner"
)

var _ runner.Solver = day8.Part1

func example1() string {
	return `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`
}

func TestPart1_Example1(t *testing.T) {
	answer := 40
	result, err := day8.Part1N(bytes.NewReader([]byte(example1())), 10)

	require.NoError(t, err)
	assert.Equal(t, answer, result)
}

func TestPart1(t *testing.T) {
	answer := 42315
	input, err := os.ReadFile("input.txt")
	require.NoError(t, err, "failed to read input.txt")

	result, err := day8.Part1(bytes.NewReader(input))

	require.NoError(t, err)
	assert.Equal(t, answer, result)
}

func TestPart2_Example1(t *testing.T) {
	answer := 25272

	result, err := day8.Part2(bytes.NewReader([]byte(example1())))

	require.NoError(t, err)
	assert.Equal(t, answer, result)
}

func TestPart2(t *testing.T) {
	answer := 8079278220
	input, err := os.ReadFile("input.txt")
	require.NoError(t, err, "failed to read input.txt")

	result, err := day8.Part2(bytes.NewReader(input))

	require.NoError(t, err)
	assert.Equal(t, answer, result)
}
