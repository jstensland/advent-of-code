package day3_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2025/day3"
	"github.com/jstensland/advent-of-code/2025/runner"
)

var _ runner.Solver = day3.Part1

func example1() string {
	return `987654321111111
811111111111119
234234234234278
818181911112111`
}

func example2() string {
	return `2712233521522212239633525221424223292522332923342263323223226223332531222232333293222213262324223122
6443848847769438847664244676354493684774344514352544147447353899987328644494647946726462934543554474
3332433333233333633323332331333325311333423333343552333343435324242433322332243454433334433332333241
6534444453544144554565453444455534444525545454555544534444654445544543445544455555444224441534555454
2216522322423236322222455428123424233411332323212234432246622229282239575322311611343244322231423352`
}

func TestPart1(t *testing.T) {
	answer := 17427
	// 17212 is too low
	// 17434 is too high

	input, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("failed to read input.txt: %v", err)
	}

	result, err := day3.Part1(bytes.NewReader(input))
	require.NoError(t, err)

	if result != answer {
		assert.Equal(t, answer, result)
	}
}

func TestPart1_Example1(t *testing.T) {
	answer := 357
	result, err := day3.Part1(bytes.NewReader([]byte(example1())))
	require.NoError(t, err, "Part1 failed")

	assert.Equal(t, answer, result)
}

func TestPart2_Example2(t *testing.T) {
	answer := 99 + 99 + 65 + 66 + 99
	result, err := day3.Part1(bytes.NewReader([]byte(example2())))
	require.NoError(t, err, "Part1 failed")

	assert.Equal(t, answer, result)
}

func TestPart2(t *testing.T) {
	answer := 0 // TODO: update to answer
	input, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("failed to read input.txt: %v", err)
	}

	result, err := day3.Part2(bytes.NewReader(input))
	require.NoError(t, err)

	if result != answer {
		assert.Equal(t, answer, result)
	}
}

func TestPart2_Example1(t *testing.T) {
	answer := 3121910778619
	result, err := day3.Part2(bytes.NewReader([]byte(example1())))
	if err != nil {
		t.Fatalf("Part2 failed: %v", err)
	}
	assert.Equal(t, answer, result)
}

func TestBiggestPerRow(t *testing.T) {
	tests := []struct {
		name     string
		row      day3.Bank
		expected int
	}{
		{
			name:     "example1 row 1: 987654321111111",
			row:      day3.Bank{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
			expected: 98,
		},
		{
			name:     "example1 row 2: 811111111111119",
			row:      day3.Bank{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
			expected: 89,
		},
		{
			name:     "example1 row 3: 234234234234278",
			row:      day3.Bank{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
			expected: 78,
		},
		{
			name:     "example1 row 4: 818181911112111",
			row:      day3.Bank{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
			expected: 92,
		},
		{
			name:     "AI found my bug and gave me this case",
			row:      day3.Bank{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9, 8},
			expected: 98,
		},

		{
			name: "example2 row 1",
			row: day3.Bank{
				2, 7, 1, 2, 2, 3, 3, 5, 2, 1, 5, 2, 2, 2, 1, 2, 2, 3, 9, 6, 3, 3, 5, 2, 5, 2, 2, 1, 4, 2, 4, 2, 2, 3, 2, 9, 2,
				5, 2, 2, 3, 3, 2, 9, 2, 3, 3, 4, 2, 2, 6, 3, 3, 2, 3, 2, 2, 3, 2, 2, 6, 2, 2, 3, 3, 3, 2, 5, 3, 1, 2, 2, 2, 2,
				3, 2, 3, 3, 3, 2, 9, 3, 2, 2, 2, 2, 1, 3, 2, 6, 2, 3, 2, 4, 2, 2, 3, 1, 2, 2,
			},
			expected: 99,
		},
		{
			name: "example2 row 2",
			row: day3.Bank{
				6, 4, 4, 3, 8, 4, 8, 8, 4, 7, 7, 6, 9, 4, 3, 8, 8, 4, 7, 6, 6, 4, 2, 4, 4, 6, 7, 6, 3, 5, 4, 4, 9, 3, 6, 8, 4,
				7, 7, 4, 3, 4, 4, 5, 1, 4, 3, 5, 2, 5, 4, 4, 1, 4, 7, 4, 4, 7, 3, 5, 3, 8, 9, 9, 9, 8, 7, 3, 2, 8, 6, 4, 4, 4,
				9, 4, 6, 4, 7, 9, 4, 6, 7, 2, 6, 4, 6, 2, 9, 3, 4, 5, 4, 3, 5, 5, 4, 4, 7, 4,
			},
			expected: 99,
		},
		{
			name: "example2 row 3",
			row: day3.Bank{
				3, 3, 3, 2, 4, 3, 3, 3, 3, 3, 2, 3, 3, 3, 3, 3, 6, 3, 3, 3, 2, 3, 3, 3, 2, 3, 3, 1, 3, 3, 3, 3, 2, 5, 3, 1, 1,
				3, 3, 3, 4, 2, 3, 3, 3, 3, 3, 4, 3, 5, 5, 2, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 2, 4, 2, 4, 2, 4, 3, 3, 3, 2, 2, 3,
				3, 2, 2, 4, 3, 4, 5, 4, 4, 3, 3, 3, 3, 4, 4, 3, 3, 3, 3, 2, 3, 3, 3, 2, 4, 1,
			},
			expected: 65,
		},
		{
			name: "example2 row 4",
			row: day3.Bank{
				6, 5, 3, 4, 4, 4, 4, 4, 5, 3, 5, 4, 4, 1, 4, 4, 5, 5, 4, 5, 6, 5, 4, 5, 3, 4, 4, 4, 4, 5, 5, 5, 3, 4, 4, 4, 4,
				5, 2, 5, 5, 4, 5, 4, 5, 4, 5, 5, 5, 5, 4, 4, 5, 3, 4, 4, 4, 4, 6, 5, 4, 4, 4, 5, 5, 4, 4, 5, 4, 3, 4, 4, 5, 5,
				4, 4, 4, 5, 5, 5, 5, 5, 4, 4, 4, 2, 2, 4, 4, 4, 1, 5, 3, 4, 5, 5, 5, 4, 5, 4,
			},
			expected: 66,
		},
		{
			name: "example2 row 5",
			row: day3.Bank{
				2, 2, 1, 6, 5, 2, 2, 3, 2, 2, 4, 2, 3, 2, 3, 6, 3, 2, 2, 2, 2, 2, 4, 5, 5, 4, 2, 8, 1, 2, 3, 4, 2, 4, 2, 3, 3,
				4, 1, 1, 3, 3, 2, 3, 2, 3, 2, 1, 2, 2, 3, 4, 4, 3, 2, 2, 4, 6, 6, 2, 2, 2, 2, 9, 2, 8, 2, 2, 3, 9, 5, 7, 5, 3,
				2, 2, 3, 1, 1, 6, 1, 1, 3, 4, 3, 2, 4, 4, 3, 2, 2, 2, 3, 1, 4, 2, 3, 3, 5, 2,
			},
			expected: 99,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := day3.Biggest(tt.row)
			assert.Equal(t, tt.expected, result)
		})
	}
}
