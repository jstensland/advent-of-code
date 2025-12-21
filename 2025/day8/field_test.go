package day8_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2025/day8"
)

func TestNewField(t *testing.T) {
	tests := []struct {
		name            string
		points          []day8.Point
		wantNumSets     int
		wantNumPairs    int
		wantShortestLen float64
	}{
		{
			name:            "empty field",
			points:          []day8.Point{},
			wantNumSets:     0,
			wantNumPairs:    0,
			wantShortestLen: 0,
		},
		{
			name: "single point",
			points: []day8.Point{
				{X: 0, Y: 0, Z: 0},
			},
			wantNumSets:     1,
			wantNumPairs:    1, // only self-distance
			wantShortestLen: 0.0,
		},
		{
			name: "two points",
			points: []day8.Point{
				{X: 0, Y: 0, Z: 0},
				{X: 1, Y: 0, Z: 0},
			},
			wantNumSets:     2,
			wantNumPairs:    4, // 2 self-distances + 2 between-point distances
			wantShortestLen: 0.0,
		},
		{
			name: "three points in a line",
			points: []day8.Point{
				{X: 0, Y: 0, Z: 0},
				{X: 1, Y: 0, Z: 0},
				{X: 2, Y: 0, Z: 0},
			},
			wantNumSets:     3,
			wantNumPairs:    9, // 3 self-distances + 6 between-point distances
			wantShortestLen: 0.0,
		},
		{
			name: "three points forming a triangle",
			points: []day8.Point{
				{X: 0, Y: 0, Z: 0},
				{X: 3, Y: 0, Z: 0},
				{X: 0, Y: 4, Z: 0},
			},
			wantNumSets:     3,
			wantNumPairs:    9,
			wantShortestLen: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := day8.NewField(tt.points)

			// Check number of sets (each point should be in its own set initially)
			sets := field.SortedSets()
			assert.Len(t, sets, tt.wantNumSets, "unexpected number of sets")

			// Verify each set has exactly one point
			for i, set := range sets {
				assert.Equal(t, 1, set.Size(), "set %d should have size 1", i)
			}

			// Check that FindSet works for each point
			for _, point := range tt.points {
				set := field.FindSet(point)
				assert.NotNil(t, set, "FindSet should return a set for point %v", point)
				assert.Equal(t, 1, set.Size(), "initial set for point %v should have size 1", point)
			}
		})
	}
}

func TestNewField_PairsSorted(t *testing.T) {
	// Create a field with points at varying distances
	points := []day8.Point{
		{X: 0, Y: 0, Z: 0},
		{X: 5, Y: 0, Z: 0}, // distance 5 from origin
		{X: 1, Y: 0, Z: 0}, // distance 1 from origin
	}

	field := day8.NewField(points)

	// Get initial number of sets
	initialSets := len(field.SortedSets())
	assert.Equal(t, 3, initialSets, "should start with 3 sets")

	field.ConnectN(1)

	// Should now have 2 sets (two closest points merged)
	assert.Len(t, field.SortedSets(), 2, "should have 2 sets after connecting closest pair")
}

func TestNewField_PairsSortedByLength(t *testing.T) {
	// Parse example1 points
	points, err := day8.ParseIn(bytes.NewReader([]byte(example1())))
	require.NoError(t, err, "failed to parse example1")

	// Create field which should calculate and sort all pairs
	field := day8.NewField(points)
	pairs := field.Pairs()

	// Verify we have pairs (n points create n*(n-1)/2 unique pairs since we exclude self-distances)
	expectedPairs := len(points) * (len(points) - 1) / 2
	assert.Len(t, pairs, expectedPairs, "expected n*(n-1)/2 unique pairs for n points")

	// Verify pairs are sorted from smallest to largest distance
	for i := 1; i < len(pairs); i++ {
		prevLength := pairs[i-1].Length()
		currLength := pairs[i].Length()
		assert.LessOrEqual(t, prevLength, currLength,
			"pair at index %d (length %.6f) should not be greater than pair at index %d (length %.6f)",
			i-1, prevLength, i, currLength)
	}

	// Additional sanity check: first pair should have smallest distance, last should have largest
	if len(pairs) > 0 {
		smallestDistance := pairs[0].Length()
		largestDistance := pairs[len(pairs)-1].Length()
		assert.LessOrEqual(t, smallestDistance, largestDistance,
			"first pair distance (%.6f) should be <= last pair distance (%.6f)",
			smallestDistance, largestDistance)
	}
}

func TestNewField_MultiplyLargest(t *testing.T) {
	tests := []struct {
		name        string
		points      []day8.Point
		connections int
		numLargest  int
		wantProduct int
	}{
		{
			name: "three separate points",
			points: []day8.Point{
				{X: 0, Y: 0, Z: 0},
				{X: 10, Y: 0, Z: 0},
				{X: 20, Y: 0, Z: 0},
			},
			connections: 0, // no connections made
			numLargest:  3,
			wantProduct: 1, // 1 * 1 * 1 = 1
		},
		{
			name: "connect two points",
			points: []day8.Point{
				{X: 0, Y: 0, Z: 0},
				{X: 1, Y: 0, Z: 0},
				{X: 10, Y: 0, Z: 0},
			},
			connections: 1,
			numLargest:  2,
			wantProduct: 2, // 2 * 1 = 2 (one set with 2 points, one with 1)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := day8.NewField(tt.points)
			field.ConnectN(tt.connections)
			result := field.MultiplyLargest(tt.numLargest)
			assert.Equal(t, tt.wantProduct, result)
		})
	}
}
