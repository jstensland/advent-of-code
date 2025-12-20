package day8

import (
	"maps"
	"slices"
)

type Field struct {
	points     []Point
	pairs      []Distance
	sets       map[string]*Set
	pointToSet map[Point]*Set
}

func NewField(points []Point) *Field {
	// Calculate the distance between every point and sort them from least to greatest
	pairs := []Distance{}
	// TODO: this id doubled!
	for idx, point1 := range points {
		for _, point2 := range points[idx:] {
			if point1 != point2 {
				pairs = append(pairs, point1.From(point2))
			}
		}
	}

	slices.SortFunc(pairs, func(a, b Distance) int {
		switch {
		case a.Length() < b.Length():
			return -1
		case a.Length() > b.Length():
			return 1
		default:
			return 0
		}
	})

	out := &Field{
		points:     points,
		pairs:      pairs,
		sets:       map[string]*Set{},
		pointToSet: map[Point]*Set{},
	}

	// Put each point in it's own set
	for _, point := range points {
		set := NewSet([]Point{point})
		out.Set(&set)
	}

	return out
}

// Pairs exposes the parsed pairs. It should probably make a copy instead
//
// TODO: currently for testing only.
func (f *Field) Pairs() []Distance {
	return f.pairs
}

func (f *Field) ConnectN(num int) {
	for idx := range num {
		if idx >= len(f.pairs) {
			panic("tried to make more connections than there are pairs to connect")
		}
		shortestDistance := f.pairs[idx]
		f.connect(shortestDistance.Ends())
	}
}

func (f *Field) ConnectAll() Distance {
	idx := 0
	var shortestDistance Distance
	for f.NumSets() > 1 {
		shortestDistance = f.pairs[idx]
		f.connect(shortestDistance.Ends())
		idx++
	}
	return shortestDistance
}

func (f *Field) Del(s *Set) {
	delete(f.sets, s.String())
	for _, p := range s.Points() {
		delete(f.pointToSet, p)
	}
}

func (f *Field) Set(s *Set) {
	f.sets[s.String()] = s
	for _, p := range s.Points() {
		f.pointToSet[p] = s
	}
}

func (f *Field) FindSet(p Point) *Set {
	return f.pointToSet[p]
}

func (f *Field) NumSets() int {
	return len(f.sets)
}

// SortedSets returns the sets sorted in reverse order by size.
func (f *Field) SortedSets() []*Set {
	sortedSets := slices.Collect(maps.Values(f.sets))
	slices.SortFunc(sortedSets, func(a, b *Set) int {
		return b.Size() - a.Size() // reverse sort
	})
	return sortedSets
}

// MultiplyLargest returns the sizes of the largest num sets.
func (f *Field) MultiplyLargest(num int) int {
	largest := f.SortedSets()[:num]
	out := 1
	for _, value := range largest {
		out *= value.Size()
	}
	return out
}

// Connect mutates the field by connecting the num of closets points.
func (f *Field) connect(ps [2]Point) {
	// find the set for p1 and p2
	p1, p2 := ps[0], ps[1]
	s1 := f.FindSet(p1)
	s2 := f.FindSet(p2) // TODO: some how on the second iteration of the test, the p2 is not finding a set already!

	// TODO: I could do better than this with a method.... s1.Equals(s2)
	if s1.String() == s2.String() {
		return
	}
	s3 := s1.Merge(s2)
	f.Del(s1)
	f.Del(s2)
	f.Set(s3)
}
