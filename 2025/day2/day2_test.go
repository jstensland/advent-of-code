package day2_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2025/day2"
	"github.com/jstensland/advent-of-code/2025/runner"
)

var _ runner.Solver = day2.Part1

func example1() string {
	return `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,` +
		`565653-565659,824824821-824824827,2121212118-2121212124`
}

func TestPart1(t *testing.T) {
	answer := 43952536386 // TODO: update to answer
	input, err := os.ReadFile("input.txt")
	require.NoError(t, err, "failed to read input.txt")

	result, err := day2.Part1(bytes.NewReader(input))

	require.NoError(t, err)
	assert.Equal(t, result, answer)
}

func TestPart1_Example1(t *testing.T) {
	answer := 1227775554

	result, err := day2.Part1(bytes.NewReader([]byte(example1())))

	require.NoError(t, err, "Part1 failed")
	assert.Equal(t, answer, result)
}

// func TestPart2(t *testing.T) {
// 	answer := 0 // TODO: update to answer
// 	input, err := os.ReadFile("input.txt")
// 	if err != nil {
// 		t.Fatalf("failed to read input.txt: %v", err)
// 	}
//
// 	result, err := day2.Part2(bytes.NewReader(input))
// 	require.NoError(t, err)
//
// 	if result != answer {
// 		assert.Equal(t, answer, result)
// 	}
// }
//
// func TestPart2_Example1(t *testing.T) {
// 	answer := 0 // TODO: update to answer
// 	result, err := day2.Part2(bytes.NewReader([]byte(example1())))
// 	if err != nil {
// 		t.Fatalf("Part2 failed: %v", err)
// 	}
// 	assert.Equal(t, answer, result)
// }

func TestParseIn(t *testing.T) {
	result, err := day2.ParseIn(bytes.NewReader([]byte(`11-22,95-115,1188511880-1188511890`)))
	require.NoError(t, err)

	expected := []day2.Range{
		{Start: []int{1, 1}, End: []int{2, 2}},
		{Start: []int{9, 5}, End: []int{1, 1, 5}},
		{Start: []int{1, 1, 8, 8, 5, 1, 1, 8, 8, 0}, End: []int{1, 1, 8, 8, 5, 1, 1, 8, 9, 0}},
	}

	assert.Equal(t, expected, result)
}

func TestInvalidIDs(t *testing.T) {
	tests := []struct {
		name     string
		r        day2.Range
		expected []day2.ID
	}{
		{
			name:     "11-22 has two invalid IDs, 11 and 22",
			r:        day2.Range{Start: day2.ID{1, 1}, End: day2.ID{2, 2}},
			expected: []day2.ID{{1, 1}, {2, 2}},
		},
		{
			name:     "95-115 has one invalid ID, 99",
			r:        day2.Range{Start: day2.ID{9, 5}, End: day2.ID{1, 1, 5}},
			expected: []day2.ID{{9, 9}},
		},
		{
			name:     "998-1012 has one invalid ID, 1010",
			r:        day2.Range{Start: day2.ID{9, 9, 8}, End: day2.ID{1, 0, 1, 2}},
			expected: []day2.ID{{1, 0, 1, 0}},
		},
		{
			name: "1188511880-1188511890 has one invalid ID, 1188511885",
			r: day2.Range{
				Start: day2.ID{1, 1, 8, 8, 5, 1, 1, 8, 8, 0},
				End:   day2.ID{1, 1, 8, 8, 5, 1, 1, 8, 9, 0},
			},
			expected: []day2.ID{{1, 1, 8, 8, 5, 1, 1, 8, 8, 5}},
		},
		{
			name:     "222220-222224 has one invalid ID, 222222",
			r:        day2.Range{Start: day2.ID{2, 2, 2, 2, 2, 0}, End: day2.ID{2, 2, 2, 2, 2, 4}},
			expected: []day2.ID{{2, 2, 2, 2, 2, 2}},
		},
		{
			name:     "1698522-1698528 contains no invalid IDs",
			r:        day2.Range{Start: day2.ID{1, 6, 9, 8, 5, 2, 2}, End: day2.ID{1, 6, 9, 8, 5, 2, 8}},
			expected: []day2.ID{},
		},
		{
			name:     "446443-446449 has one invalid ID, 446446",
			r:        day2.Range{Start: day2.ID{4, 4, 6, 4, 4, 3}, End: day2.ID{4, 4, 6, 4, 4, 9}},
			expected: []day2.ID{{4, 4, 6, 4, 4, 6}},
		},
		{
			name:     "38593856-38593862 has one invalid ID, 38593859",
			r:        day2.Range{Start: day2.ID{3, 8, 5, 9, 3, 8, 5, 6}, End: day2.ID{3, 8, 5, 9, 3, 8, 6, 2}},
			expected: []day2.ID{{3, 8, 5, 9, 3, 8, 5, 9}},
		},
		{
			name:     "565653-565659 contains no invalid IDs",
			r:        day2.Range{Start: day2.ID{5, 6, 5, 6, 5, 3}, End: day2.ID{5, 6, 5, 6, 5, 9}},
			expected: []day2.ID{},
		},
		{
			name:     "824824821-824824827 contains no invalid IDs",
			r:        day2.Range{Start: day2.ID{8, 2, 4, 8, 2, 4, 8, 2, 1}, End: day2.ID{8, 2, 4, 8, 2, 4, 8, 2, 7}},
			expected: []day2.ID{},
		},
		{
			name: "2121212118-2121212124 contains no invalid IDs",
			r: day2.Range{
				Start: day2.ID{2, 1, 2, 1, 2, 1, 2, 1, 1, 8},
				End:   day2.ID{2, 1, 2, 1, 2, 1, 2, 1, 2, 4},
			},
			expected: []day2.ID{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := day2.InvalidIDs(tt.r)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// func TestInvalidIDsV2(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		r        day2.Range
// 		expected []day2.ID
// 	}{
// 		{
// 			name:     "11-22 still has two invalid IDs, 11 and 22",
// 			r:        day2.Range{Start: day2.ID{1, 1}, End: day2.ID{2, 2}},
// 			expected: []day2.ID{{1, 1}, {2, 2}},
// 		},
// 		{
// 			name:     "95-115 now has two invalid IDs, 99 and 111",
// 			r:        day2.Range{Start: day2.ID{9, 5}, End: day2.ID{1, 1, 5}},
// 			expected: []day2.ID{{9, 9}, {1, 1, 1}},
// 		},
// 		{
// 			name:     "998-1012 now has two invalid IDs, 999 and 1010",
// 			r:        day2.Range{Start: day2.ID{9, 9, 8}, End: day2.ID{1, 0, 1, 2}},
// 			expected: []day2.ID{{9, 9, 9}, {1, 0, 1, 0}},
// 		},
// 		{
// 			name: "1188511880-1188511890 still has one invalid ID, 1188511885",
// 			r: day2.Range{
// 				Start: day2.ID{1, 1, 8, 8, 5, 1, 1, 8, 8, 0},
// 				End:   day2.ID{1, 1, 8, 8, 5, 1, 1, 8, 9, 0},
// 			},
// 			expected: []day2.ID{{1, 1, 8, 8, 5, 1, 1, 8, 8, 5}},
// 		},
// 		{
// 			name:     "222220-222224 still has one invalid ID, 222222",
// 			r:        day2.Range{Start: day2.ID{2, 2, 2, 2, 2, 0}, End: day2.ID{2, 2, 2, 2, 2, 4}},
// 			expected: []day2.ID{{2, 2, 2, 2, 2, 2}},
// 		},
// 		{
// 			name:     "1698522-1698528 still contains no invalid IDs",
// 			r:        day2.Range{Start: day2.ID{1, 6, 9, 8, 5, 2, 2}, End: day2.ID{1, 6, 9, 8, 5, 2, 8}},
// 			expected: []day2.ID{},
// 		},
// 		{
// 			name:     "446443-446449 still has one invalid ID, 446446",
// 			r:        day2.Range{Start: day2.ID{4, 4, 6, 4, 4, 3}, End: day2.ID{4, 4, 6, 4, 4, 9}},
// 			expected: []day2.ID{{4, 4, 6, 4, 4, 6}},
// 		},
// 		{
// 			name:     "38593856-38593862 still has one invalid ID, 38593859",
// 			r:        day2.Range{Start: day2.ID{3, 8, 5, 9, 3, 8, 5, 6}, End: day2.ID{3, 8, 5, 9, 3, 8, 6, 2}},
// 			expected: []day2.ID{{3, 8, 5, 9, 3, 8, 5, 9}},
// 		},
// 		{
// 			name:     "565653-565659 now has one invalid ID, 565656",
// 			r:        day2.Range{Start: day2.ID{5, 6, 5, 6, 5, 3}, End: day2.ID{5, 6, 5, 6, 5, 9}},
// 			expected: []day2.ID{{5, 6, 5, 6, 5, 6}},
// 		},
// 		{
// 			name:     "824824821-824824827 now has one invalid ID, 824824824",
// 			r:        day2.Range{Start: day2.ID{8, 2, 4, 8, 2, 4, 8, 2, 1}, End: day2.ID{8, 2, 4, 8, 2, 4, 8, 2, 7}},
// 			expected: []day2.ID{{8, 2, 4, 8, 2, 4, 8, 2, 4}},
// 		},
// 		{
// 			name: "2121212118-2121212124 now has one invalid ID, 2121212121",
// 			r: day2.Range{
// 				Start: day2.ID{2, 1, 2, 1, 2, 1, 2, 1, 1, 8},
// 				End:   day2.ID{2, 1, 2, 1, 2, 1, 2, 1, 2, 4},
// 			},
// 			expected: []day2.ID{{2, 1, 2, 1, 2, 1, 2, 1, 2, 1}},
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			result, err := day2.InvalidIDsV2(tt.r)
// 			require.NoError(t, err)
// 			assert.Equal(t, tt.expected, result)
// 		})
// 	}
// }
