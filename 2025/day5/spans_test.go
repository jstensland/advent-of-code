package day5_test

import (
	"testing"

	"github.com/jstensland/advent-of-code/2025/day5"
)

func TestCombineRanges(t *testing.T) {
	tests := []struct {
		name  string
		spans []day5.Span
		want  []day5.Span
	}{
		{
			name:  "empty list",
			spans: []day5.Span{},
			want:  []day5.Span{},
		},
		{
			name:  "single span",
			spans: []day5.Span{{Start: 1, End: 5}},
			want:  []day5.Span{{Start: 1, End: 5}},
		},
		{
			name: "two non-overlapping spans (sorted)",
			spans: []day5.Span{
				{Start: 1, End: 3},
				{Start: 5, End: 7},
			},
			want: []day5.Span{
				{Start: 1, End: 3},
				{Start: 5, End: 7},
			},
		},
		{
			name: "two non-overlapping spans (unsorted)",
			spans: []day5.Span{
				{Start: 5, End: 7},
				{Start: 1, End: 3},
			},
			want: []day5.Span{
				{Start: 1, End: 3},
				{Start: 5, End: 7},
			},
		},
		{
			name: "two overlapping spans",
			spans: []day5.Span{
				{Start: 1, End: 5},
				{Start: 3, End: 7},
			},
			want: []day5.Span{
				{Start: 1, End: 7},
			},
		},
		{
			name: "two overlapping spans (reverse order)",
			spans: []day5.Span{
				{Start: 3, End: 7},
				{Start: 1, End: 5},
			},
			want: []day5.Span{
				{Start: 1, End: 7},
			},
		},
		{
			name: "one span encompasses another",
			spans: []day5.Span{
				{Start: 1, End: 10},
				{Start: 3, End: 5},
			},
			want: []day5.Span{
				{Start: 1, End: 10},
			},
		},
		{
			name: "one span encompasses another (reverse order)",
			spans: []day5.Span{
				{Start: 3, End: 5},
				{Start: 1, End: 10},
			},
			want: []day5.Span{
				{Start: 1, End: 10},
			},
		},
		{
			name: "multiple overlapping spans merge into one",
			spans: []day5.Span{
				{Start: 1, End: 3},
				{Start: 2, End: 5},
				{Start: 4, End: 8},
			},
			want: []day5.Span{
				{Start: 1, End: 8},
			},
		},
		{
			name: "adjacent spans (should NOT merge)",
			spans: []day5.Span{
				{Start: 1, End: 3},
				{Start: 4, End: 6},
			},
			want: []day5.Span{
				{Start: 1, End: 3},
				{Start: 4, End: 6},
			},
		},
		{
			name: "touching spans at boundary",
			spans: []day5.Span{
				{Start: 1, End: 5},
				{Start: 5, End: 10},
			},
			want: []day5.Span{
				{Start: 1, End: 10},
			},
		},
		{
			name: "mix of overlapping and non-overlapping spans",
			spans: []day5.Span{
				{Start: 1, End: 3},
				{Start: 5, End: 7},
				{Start: 2, End: 4},
				{Start: 10, End: 12},
			},
			want: []day5.Span{
				{Start: 1, End: 4},
				{Start: 5, End: 7},
				{Start: 10, End: 12},
			},
		},
		{
			name: "duplicate spans",
			spans: []day5.Span{
				{Start: 1, End: 5},
				{Start: 1, End: 5},
			},
			want: []day5.Span{
				{Start: 1, End: 5},
			},
		},
		{
			name: "many spans with complex overlaps",
			spans: []day5.Span{
				{Start: 1, End: 2},
				{Start: 10, End: 15},
				{Start: 3, End: 5},
				{Start: 12, End: 20},
				{Start: 4, End: 6},
				{Start: 25, End: 30},
			},
			want: []day5.Span{
				{Start: 1, End: 2},
				{Start: 3, End: 6},
				{Start: 10, End: 20},
				{Start: 25, End: 30},
			},
		},
		{
			name: "span added before all existing",
			spans: []day5.Span{
				{Start: 10, End: 15},
				{Start: 1, End: 3},
			},
			want: []day5.Span{
				{Start: 1, End: 3},
				{Start: 10, End: 15},
			},
		},
		{
			name: "span added after all existing",
			spans: []day5.Span{
				{Start: 1, End: 3},
				{Start: 10, End: 15},
			},
			want: []day5.Span{
				{Start: 1, End: 3},
				{Start: 10, End: 15},
			},
		},
		{
			name: "span inserted in the middle (no overlap)",
			spans: []day5.Span{
				{Start: 1, End: 3},
				{Start: 10, End: 15},
				{Start: 5, End: 7},
			},
			want: []day5.Span{
				{Start: 1, End: 3},
				{Start: 5, End: 7},
				{Start: 10, End: 15},
			},
		},
		{
			name: "all spans overlap into one giant span",
			spans: []day5.Span{
				{Start: 1, End: 5},
				{Start: 3, End: 8},
				{Start: 7, End: 12},
				{Start: 10, End: 15},
			},
			want: []day5.Span{
				{Start: 1, End: 15},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day5.CombineRanges(tt.spans)
			if !spansEqual(got, tt.want) {
				t.Errorf("CombineRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func spansEqual(a, b []day5.Span) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Start != b[i].Start || a[i].End != b[i].End {
			return false
		}
	}
	return true
}
