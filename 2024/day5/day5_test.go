package day5_test

import (
	"advent2024/day5"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func exampleInput() io.ReadCloser {
	return io.NopCloser(strings.NewReader(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`))
}

func TestRunPart1Example(t *testing.T) {
	total, err := day5.RunPart1(exampleInput())

	require.NoError(t, err)
	assert.Equal(t, 143, total)
}

func TestRunPart1(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	total, err := day5.RunPart1(in)

	require.NoError(t, err)
	assert.Equal(t, 7074, total) // confired
}

func TestRunPart2(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	total, err := day5.RunPart2(in)

	require.NoError(t, err)
	assert.Equal(t, 4828, total) // confirmed
}
