package day9

import (
	"iter"
)

func FindExtremes(points []Point) (int, int, int, int) {
	p1 := points[0]
	minX, maxX, minY, maxY := p1.X, p1.X, p1.Y, p1.Y

	for _, p := range points[1:] {
		minX = min(minX, p.X)
		maxX = max(maxX, p.X)
		minY = min(minY, p.Y)
		maxY = max(maxY, p.Y)
	}

	return minX, maxX, minY, maxY
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func Turn(d Direction) Direction {
	const numDirs = 4
	return (d + 1) % numDirs
}

// FurtherFromCenter returns a sort function that puts the points further from the center of the provided min/max
// values first.
func FurtherFromCenter(minX, maxX, minY, maxY int) func(p1, p2 Point) int {
	// TODO: this isn't quite the middle... integer division
	const numNums = 2
	center := Point{
		X: (minX + maxX) / numNums,
		Y: (minY + maxY) / numNums,
	}

	return func(p1, p2 Point) int {
		// should be negative when p1 < p2 for forward sort
		// we want furthest first
		// negative when p2 is closer to center
		p2FromCenter := distance(p2, center)
		p1FromCenter := distance(p1, center)
		if p2FromCenter < p1FromCenter {
			return -1
		}
		return 1
	}
}

// SpiralIn was not used in the end, testing too many points, but keeping it for future reference
// it does require some refactor still though.
//
//nolint:cyclop // skipping refactor as not used
func SpiralIn(minX, maxX, minY, maxY int) iter.Seq[Point] {
	x, y := minX, minY
	facing := Up
	return func(yield func(r Point) bool) {
		for minX < maxX && minY < maxY {
			if !yield(Point{x, y}) {
				return
			}
			switch facing {
			case Up:
				if y < maxY {
					y++
					continue
				}
			case Right:
				if x < maxX {
					x++
					continue
				}
			case Down:
				if y > minY {
					y--
					continue
				}
			case Left:
				if x > minX {
					x--
					continue
				}
				// made it back to the start
				minX++
				maxX--
				minY++
				maxY--
				y++
				x++
			}
			facing = Turn(facing)
			switch facing {
			case Up:
				y++
			case Right:
				x++
			case Down:
				y--
			case Left:
				x--
			}
		}
	}
}

func biggestToCorner(minX, maxX, minY, maxY int, p Point) int {
	return max(
		Square{Flag1: p, Flag2: Point{minX, minY}}.Area(),
		Square{Flag1: p, Flag2: Point{minX, maxY}}.Area(),
		Square{Flag1: p, Flag2: Point{maxX, minY}}.Area(),
		Square{Flag1: p, Flag2: Point{maxX, maxY}}.Area(),
	)
}

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
		// fmt.Printf("new square added with point %v and %v with area %v\n", flag, point, newSquare.Area())
		if newSquare.Area() > cp.biggest.Area() {
			cp.biggest = newSquare
		}
		cp.combos = append(cp.combos, newSquare)
	}
	cp.points[flag] = true
}

func (cp *CandidatePool) Biggest() Square { return cp.biggest }
