package day2_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2024/day2"
)

func TestRunPart1(t *testing.T) {
	answer, err := day2.RunPart1("./input.txt")

	assert.NoError(t, err)
	assert.Equal(t, 252, answer) // confirmed
}

func TestRunPart2(t *testing.T) {
	answer, err := day2.RunPart2("./input.txt")

	assert.NoError(t, err)
	assert.Equal(t, 324, answer) // confirmed
}

func TestInstructionsPart2(t *testing.T) {
	data := []byte(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`)
	in := io.NopCloser(bytes.NewBuffer(data))

	answer, err := day2.Part2Analysis(in)

	assert.NoError(t, err)
	assert.Equal(t, 4, answer) // confirmed
}

func TestInstructionsPart2BadFirst(t *testing.T) {
	testCases := []struct {
		raw    string
		result bool
	}{
		{raw: `20 20 20 21 22`, result: false},
		{raw: `20 20 21 22`, result: true},
		{raw: `23 42 41 40 38`, result: true},
		{raw: `53 12 13 14 15`, result: true},
		{raw: `42 41 40 38 38`, result: true},
		{raw: `42 41 40 38 28`, result: true},
		{raw: `34 42 41 40 38 28`, result: false},
		{raw: `83 84 84 87 88 90 94`, result: false},
		{raw: `52 49 53 55 57 60 62`, result: true}, // switch to increasing and delete the first...
		{raw: `13 13 15 17 18 19 21`, result: true},
		{raw: `25 28 24 21 19 16`, result: true},
		{raw: `18 21 23 22 23 26 28`, result: true}, // working this one

		{raw: `1 5 6 7 8`, result: true},
		{raw: `1 5 4 3 2`, result: true},
		{raw: `1 5 2 3 4`, result: true},
		{raw: `10 5 6 7 8`, result: true},
		{raw: `10 5 4 3 2`, result: true},
		{raw: `10 5 11 12 13`, result: true},
	}

	// TODO: these cases should have names, and should be run with t.Run for better output
	for _, tc := range testCases {
		in := io.NopCloser(bytes.NewBuffer([]byte(tc.raw)))

		answer, err := day2.Part2Analysis(in)

		assert.NoError(t, err)
		if tc.result {
			assert.Equal(t, 1, answer) // confirmed
		} else {
			assert.Equal(t, 0, answer) // confirmed
		}

	}
}

func TestPossibleFixes(t *testing.T) {
	testCases := []struct {
		raw  string
		idx  int
		alts [][]int
	}{
		{raw: `1 5 6 7 8`, idx: 1, alts: [][]int{{5, 6, 7, 8}, {5, 6, 7, 8}, {1, 6, 7, 8}}},
		{raw: `1 5 4 3 2`, idx: 1, alts: [][]int{{5, 4, 3, 2}, {5, 4, 3, 2}, {1, 4, 3, 2}}},
		{raw: `1 5 2 3 4`, idx: 1, alts: [][]int{{5, 2, 3, 4}, {5, 2, 3, 4}, {1, 2, 3, 4}}},
		{raw: `1 2 3 9 4 5`, idx: 3, alts: [][]int{{2, 3, 9, 4, 5}, {1, 2, 9, 4, 5}, {1, 2, 3, 4, 5}}},
		{raw: `1 2 3 3 4 5`, idx: 3, alts: [][]int{{2, 3, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}},
		{raw: `1 3 5 7 9 8`, idx: 5, alts: [][]int{{3, 5, 7, 9, 8}, {1, 3, 5, 7, 8}, {1, 3, 5, 7, 9}}},
		{raw: `1 3 5 4 5 7`, idx: 3, alts: [][]int{{3, 5, 4, 5, 7}, {1, 3, 4, 5, 7}, {1, 3, 5, 5, 7}}},
		{raw: `34 42 41 40 38 28`, idx: 1, alts: [][]int{{42, 41, 40, 38, 28}, {42, 41, 40, 38, 28}, {34, 41, 40, 38, 28}}},
		{raw: `63 60 62 65 67 69`, idx: -1, alts: [][]int{{60, 62, 65, 67, 69}}},
	}

	for _, tc := range testCases {
		rep, err := day2.ParseReport(tc.raw)
		require.NoError(t, err)

		alts := [][]int{}
		for opt := range rep.PossibleFixes(tc.idx) {
			assert.Contains(t, tc.alts, opt.Levels(), "possible fix %v not in expected %v", opt, tc.alts)
			alts = append(alts, opt.Levels())
		}
		assert.Equal(t, tc.alts, alts, "did not get expected set of alternates to test")
	}
}
