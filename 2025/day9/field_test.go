package day9_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jstensland/advent-of-code/2025/day9"
)

func TestTurn(t *testing.T) {
	testCases := []struct {
		desc    string
		initDir day9.Direction
		endDir  day9.Direction
	}{
		{
			desc:    "start Up",
			initDir: day9.Up,
			endDir:  day9.Right,
		},
		{
			desc:    "start Right",
			initDir: day9.Right,
			endDir:  day9.Down,
		},
		{
			desc:    "start Down",
			initDir: day9.Down,
			endDir:  day9.Left,
		},
		{
			desc:    "start Left",
			initDir: day9.Left,
			endDir:  day9.Up,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			assert.Equal(t, tc.endDir, day9.Turn(tc.initDir))
		})
	}
}

func TestSpiralIn(t *testing.T) {
	output := []day9.Point{}
	for p := range day9.SpiralIn(0, 5, 0, 5) {
		output = append(output, p)
	}
	expected := []day9.Point{
		{0, 0},
		{0, 1},
		{0, 2},
		{0, 3},
		{0, 4},
		{0, 5},
		{1, 5},
		{2, 5},
		{3, 5},
		{4, 5},
		{5, 5},
		{5, 4},
		{5, 3},
		{5, 2},
		{5, 1},
		{5, 0},
		{4, 0},
		{3, 0},
		{2, 0},
		{1, 0},
		{0, 0},
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 4},
		{3, 4},
		{4, 4},
		{4, 3},
		{4, 2},
		{4, 1},
		{3, 1},
		{2, 1},
		{1, 1},
		{2, 3},
		{3, 3},
		{3, 2},
		{2, 2},
	}

	assert.Equal(t, expected, output)
}

func TestDistanceFromCenter(t *testing.T) {
	cmp := day9.FurtherFromCenter(0, 4, 0, 4) // center is (2, 2)
	center := day9.Point{X: 2, Y: 2}          // distance 0
	near := day9.Point{X: 2, Y: 3}            // distance 1
	far := day9.Point{X: 0, Y: 0}             // distance ~2.83

	assert.Equal(t, -1, cmp(far, near), "far point should come before near point")
	assert.Equal(t, -1, cmp(far, center), "far point should come before center point")
	assert.Equal(t, -1, cmp(near, center), "near point should come before center point")
	assert.Equal(t, 1, cmp(near, far), "near point should come after far point")
}
