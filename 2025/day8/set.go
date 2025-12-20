package day8

import (
	"maps"
	"slices"
	"strings"
)

type Set struct {
	points map[Point]bool
}

func NewSet(ps []Point) Set {
	points := map[Point]bool{}
	for _, p := range ps {
		points[p] = true
	}
	return Set{points}
}

func (s *Set) Add(p Point) {
	s.points[p] = true
}

func (s *Set) Merge(s2 *Set) *Set {
	newSet := NewSet(s.Points())
	for p := range s2.points {
		newSet.Add(p)
	}
	return &newSet
}

// Points returns the points in the set as a slice. Order is not maintained.
func (s *Set) Points() []Point {
	return slices.Collect(maps.Keys(s.points))
}

func (s *Set) Size() int {
	return len(s.points)
}

func (s *Set) String() string {
	points := slices.Collect(maps.Keys(s.points))
	slices.SortFunc(points, comparePoints)
	var sb strings.Builder
	for i, p := range points {
		if i > 0 {
			sb.WriteString("-")
		}
		sb.WriteString(p.String())
	}
	return sb.String()
}
