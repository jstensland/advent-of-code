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
