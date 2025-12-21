// Package day9 solves AoC 2025 day 9
package day9

import (
	"fmt"
	"io"
	"slices"
)

func Part1(r io.Reader) (int, error) {
	in, err := ParseIn(r)
	if err != nil {
		return 0, err
	}
	// Algo 1
	// Iterate through the grid from the outside in a spiral
	pool := NewCandidatePool()

	fmt.Println("in:", in)
	minX, maxX, minY, maxY := FindExtremes(in)
	for p := range SpiralIn(minX, maxX, minY, maxY) {
		// TODO: add break condition
		// if the biggest square from the current point to a corner is SMALLER than our current biggest, then quit
		if biggestToCorner(minX, maxX, minY, maxY, p) < pool.Biggest().Area() {
			break
		}

		// check the point is a flag, and add it, if so
		if slices.ContainsFunc(in, func(flag Point) bool { return flag.X == p.X && flag.Y == p.Y }) {
			pool.Add(p)
		}
	}

	// Each time a flag is found, add it to the set of candidates
	// Maintain the largest square size for all candidates as you go, as well as the largest square total and which 2
	//   points make it
	// Stop early when the area from the current point to all corners, is less than the currently largest option

	fmt.Println("biggest square", pool.Biggest())

	return pool.Biggest().Area(), nil
}

func biggestToCorner(minX, maxX, minY, maxY int, p Point) int {
	return max(
		Square{Flag1: p, Flag2: Point{minX, minY}}.Area(),
		Square{Flag1: p, Flag2: Point{minX, maxY}}.Area(),
		Square{Flag1: p, Flag2: Point{maxX, minY}}.Area(),
		Square{Flag1: p, Flag2: Point{maxX, maxY}}.Area(),
	)
}

// Algo 2
// Read in all the pairs, keeping track of the greatest and least X and Y coordinate and calculating the distances
// from each flag to each corner
//
// Pick the flags that are closest to each corner, and try the combos, recording the greatest
//
// Potentially will need to woork one's way in a little bit on those closest to the corners though, similar to the
// spiral

type Square struct {
	Flag1 Point
	Flag2 Point
}

func (s Square) Area() int {
	width := max(s.Flag1.X, s.Flag2.X) - min(s.Flag1.X, s.Flag2.X) + 1
	height := max(s.Flag1.Y, s.Flag2.Y) - min(s.Flag1.Y, s.Flag2.Y) + 1
	return width * height
}

type CandidatePool struct {
	biggest Square
	points  map[Point]bool
	combos  []Square
}

func NewCandidatePool() *CandidatePool {
	return &CandidatePool{
		biggest: Square{Flag1: Point{}, Flag2: Point{}},
		points:  map[Point]bool{},
		combos:  []Square{},
	}
}

func (cp *CandidatePool) Add(flag Point) {
	for point := range cp.points {
		newSquare := Square{point, flag}
		fmt.Printf("new square added with point %v and %v with area %v\n", flag, point, newSquare.Area())
		if newSquare.Area() > cp.biggest.Area() {
			cp.biggest = newSquare
		}
		cp.combos = append(cp.combos, newSquare)
	}
	cp.points[flag] = true
}

func (cp *CandidatePool) Biggest() Square { return cp.biggest }

// 	points map[Point]bool
// 	combos [][2]Point
// }
//
// func NewSet() Set {
// 	return Set{
// 		points: map[Point]bool{},
// 		combos: [][2]Point{},
// 	}
// }
//
// func (set Set) Add(flag Point) {
// 	// calculate all combos
// 	//
// }

func Part2(r io.Reader) (int, error) {
	_, err := ParseIn(r)
	if err != nil {
		return 0, err
	}
	// TODO: solve part 2
	answer := 0

	return answer, nil
}

// func FindCorners(points []Point) (Point, Point, Point, Point) {
// 	p1 := points[0]
// 	minX, maxX, minY, maxY := p1.X, p1.X, p1.Y, p1.Y
//
// 	for _, p := range points[1:] {
// 		minX = min(minX, p.X)
// 		maxX = max(maxX, p.X)
// 		minY = min(minY, p.Y)
// 		maxY = max(maxY, p.Y)
// 	}
//
// 	return Point{minX, minY}, Point{minX, maxY}, Point{maxX, maxY}, Point{maxX, minY}
// }
