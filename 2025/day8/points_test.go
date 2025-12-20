package day8_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jstensland/advent-of-code/2025/day8"
)

func TestPointFrom_Sorting(t *testing.T) {
	tests := []struct {
		name     string
		p1       day8.Point
		p2       day8.Point
		wantP1   day8.Point // expected first point after sorting
		wantP2   day8.Point // expected second point after sorting
		distance float64
	}{
		{
			name:     "x differs - x takes priority",
			p1:       day8.Point{X: 5, Y: 0, Z: 0},
			p2:       day8.Point{X: 3, Y: 100, Z: 100},
			wantP1:   day8.Point{X: 3, Y: 100, Z: 100},
			wantP2:   day8.Point{X: 5, Y: 0, Z: 0},
			distance: 27.0,
		},
		{
			name:     "x same, y differs - y takes priority",
			p1:       day8.Point{X: 5, Y: 10, Z: 0},
			p2:       day8.Point{X: 5, Y: 3, Z: 100},
			wantP1:   day8.Point{X: 5, Y: 3, Z: 100},
			wantP2:   day8.Point{X: 5, Y: 10, Z: 0},
			distance: 22.0,
		},
		{
			name:     "x and y same, z differs - z determines order",
			p1:       day8.Point{X: 5, Y: 10, Z: 20},
			p2:       day8.Point{X: 5, Y: 10, Z: 15},
			wantP1:   day8.Point{X: 5, Y: 10, Z: 15},
			wantP2:   day8.Point{X: 5, Y: 10, Z: 20},
			distance: 3.0,
		},
		{
			name:     "x dominates even with large y and z differences",
			p1:       day8.Point{X: 1, Y: 999, Z: 999},
			p2:       day8.Point{X: 2, Y: 0, Z: 0},
			wantP1:   day8.Point{X: 1, Y: 999, Z: 999},
			wantP2:   day8.Point{X: 2, Y: 0, Z: 0},
			distance: 126.0,
		},
		{
			name:     "y dominates z when x is same",
			p1:       day8.Point{X: 5, Y: 2, Z: 999},
			p2:       day8.Point{X: 5, Y: 1, Z: 0},
			wantP1:   day8.Point{X: 5, Y: 1, Z: 0},
			wantP2:   day8.Point{X: 5, Y: 2, Z: 999},
			distance: 100,
		},
		{
			name:     "negative values - x still dominates",
			p1:       day8.Point{X: -1, Y: 100, Z: 100},
			p2:       day8.Point{X: -5, Y: 0, Z: 0},
			wantP1:   day8.Point{X: -5, Y: 0, Z: 0},
			wantP2:   day8.Point{X: -1, Y: 100, Z: 100},
			distance: 27,
		},
		{
			name:     "already sorted - should remain in order",
			p1:       day8.Point{X: 1, Y: 2, Z: 3},
			p2:       day8.Point{X: 4, Y: 5, Z: 6},
			wantP2:   day8.Point{X: 4, Y: 5, Z: 6},
			wantP1:   day8.Point{X: 1, Y: 2, Z: 3},
			distance: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dist := tt.p1.From(tt.p2)
			assert.Equal(t, tt.distance, math.Round(dist.Length()))

			// Check that the points in the distance are sorted correctly
			if dist.Ends()[0] != tt.wantP1 {
				t.Errorf("first point = %v, want %v", dist.Ends()[0], tt.wantP1)
			}
			if dist.Ends()[1] != tt.wantP2 {
				t.Errorf("second point = %v, want %v", dist.Ends()[1], tt.wantP2)
			}
		})
	}
}
