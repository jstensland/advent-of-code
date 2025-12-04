package day1_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2025/day1"
	"github.com/jstensland/advent-of-code/2025/runner"
)

var _ runner.Solver = day1.Part1

func example1() string {
	return `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
}

func TestPart1(t *testing.T) {
	answer := 1102

	input, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("failed to read input.txt: %v", err)
	}

	result, err := day1.Part1(bytes.NewReader(input))
	require.NoError(t, err)

	if result != answer {
		assert.Equal(t, result, answer)
	}
}

func TestPart1_CustomInput(t *testing.T) {
	result, err := day1.Part1(bytes.NewReader([]byte(example1())))
	if err != nil {
		t.Fatalf("Part1 failed: %v", err)
	}

	assert.Equal(t, 3, result)
}

func TestPart2(t *testing.T) {
	answer := 6175

	input, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("failed to read input.txt: %v", err)
	}

	result, err := day1.Part2(bytes.NewReader(input))
	require.NoError(t, err)

	if result != answer {
		assert.Equal(t, answer, result)
	}
}

func TestPart2_CustomInput(t *testing.T) {
	result, err := day1.Part2(bytes.NewReader([]byte(example1())))
	if err != nil {
		t.Fatalf("Part2 failed: %v", err)
	}

	assert.Equal(t, 6, result)
}

func TestPart2_AllRight(t *testing.T) {
	result, err := day1.Part2(bytes.NewReader([]byte(
		`R51
R99
R500
R1
R500
`)))
	require.NoError(t, err)

	assert.Equal(t, 12, result)
}

func TestPart2_AllLeft(t *testing.T) {
	result, err := day1.Part2(bytes.NewReader([]byte(
		`L40
L10
L25
L76
L500
`)))
	require.NoError(t, err)

	assert.Equal(t, 7, result)
}

func TestMoveRight(t *testing.T) {
	tests := []struct {
		name          string
		startPos      int
		distance      int
		expectedPos   int
		expectedZeros int
	}{
		{
			name:          "land on 0 without rotation",
			startPos:      50,
			distance:      50,
			expectedPos:   0,
			expectedZeros: 1,
		},
		{
			name:          "land on 0 with one rotation",
			startPos:      10,
			distance:      90,
			expectedPos:   0,
			expectedZeros: 1,
		},
		{
			name:          "go around once from 0",
			startPos:      0,
			distance:      100,
			expectedPos:   0,
			expectedZeros: 1,
		},
		{
			name:          "go around once from 50",
			startPos:      50,
			distance:      100,
			expectedPos:   50,
			expectedZeros: 1,
		},
		{
			name:          "go around twice",
			startPos:      0,
			distance:      250,
			expectedPos:   50,
			expectedZeros: 2,
		},
		{
			name:          "go around four times",
			startPos:      25,
			distance:      375,
			expectedPos:   0,
			expectedZeros: 4,
		},
		{
			name:          "go around three times landing on 99",
			startPos:      99,
			distance:      300,
			expectedPos:   99,
			expectedZeros: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos, rotations := day1.MoveDial(tt.startPos, tt.distance)
			assert.Equal(t, tt.expectedPos, pos, "position mismatch")
			assert.Equal(t, tt.expectedZeros, rotations, "rotations mismatch")
		})
	}
}

func TestMoveLeft(t *testing.T) {
	tests := []struct {
		name          string
		startPos      int
		distance      int
		expectedPos   int
		expectedZeros int
	}{
		{
			name:          "do not pass 0",
			startPos:      50,
			distance:      -30,
			expectedPos:   20,
			expectedZeros: 0,
		},
		{
			name:          "land on 0 without rotation",
			startPos:      50,
			distance:      -50,
			expectedPos:   0,
			expectedZeros: 1,
		},
		{
			name:          "simple wraparound",
			startPos:      25,
			distance:      -50,
			expectedPos:   75,
			expectedZeros: 1,
		},
		{
			name:          "land on 0 with one rotation",
			startPos:      50,
			distance:      -150,
			expectedPos:   0,
			expectedZeros: 2,
		},
		{
			name:          "go around once from 0",
			startPos:      0,
			distance:      -100,
			expectedPos:   0,
			expectedZeros: 1,
		},
		{
			name:          "go around once from 50",
			startPos:      50,
			distance:      -100,
			expectedPos:   50,
			expectedZeros: 1,
		},
		{
			name:          "go around twice",
			startPos:      50,
			distance:      -250,
			expectedPos:   0,
			expectedZeros: 3,
		},
		{
			name:          "go around three times",
			startPos:      25,
			distance:      -300,
			expectedPos:   25,
			expectedZeros: 3,
		},
		{
			name:          "go around four times landing on 99",
			startPos:      99,
			distance:      -400,
			expectedPos:   99,
			expectedZeros: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos, rotations := day1.MoveDial(tt.startPos, tt.distance)
			assert.Equal(t, tt.expectedPos, pos, "position mismatch")
			assert.Equal(t, tt.expectedZeros, rotations, "rotations mismatch")
		})
	}
}

func TestParseMoves(t *testing.T) {
	moves, err := day1.ParseMoves(bytes.NewReader([]byte(example1())))
	if err != nil {
		t.Fatalf("ParseMoves failed: %v", err)
	}

	expected := []day1.Move{
		{Distance: -68},
		{Distance: -30},
		{Distance: 48},
		{Distance: -5},
		{Distance: 60},
		{Distance: -55},
		{Distance: -1},
		{Distance: -99},
		{Distance: 14},
		{Distance: -82},
	}

	assert.Equal(t, expected, moves)
}
