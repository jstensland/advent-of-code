package day9

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p Point) From(p2 Point) Distance {
	p1 := Point{
		X: p.X,
		Y: p.Y,
	}
	// sort p1 and p2, as p1 -> p2 == p2 -> p1
	if comparePoints(p1, p2) > 0 {
		p1, p2 = p2, p1
	}
	points := [2]Point{p1, p2}

	return Distance{
		ends:   points,
		length: distance(p1, p2),
	}
}

func distance(p1, p2 Point) float64 {
	const squarePower = 2
	const oneHalf = 1.0 / 2
	return math.Pow(
		math.Pow(float64(p1.X-p2.X), squarePower)+
			math.Pow(float64(p1.Y-p2.Y), squarePower), oneHalf)
}

func (p Point) String() string {
	return fmt.Sprintf("%v,%v", p.X, p.Y)
}

func comparePoints(p1, p2 Point) int {
	// p1 < p2 => -1
	// x is more important than y is more important than z
	if p1.X > p2.X {
		return 1
	}
	if p1.X < p2.X {
		return -1
	}
	if p1.Y > p2.Y {
		return 1
	}
	return -1
}

type Distance struct {
	ends   [2]Point
	length float64
}

func (d Distance) Length() float64 {
	return d.length
}

func (d Distance) Ends() [2]Point {
	return d.ends
}

func ParseIn(r io.Reader) ([]Point, error) {
	scanner := bufio.NewScanner(r)

	points := []Point{}
	for scanner.Scan() {
		line := scanner.Text()
		coords := []int{}
		for coordinate := range strings.SplitSeq(line, ",") {
			num, err := strconv.Atoi(coordinate)
			if err != nil {
				return nil, fmt.Errorf("invalid number on line %v, %v: %w", line, num, err)
			}
			coords = append(coords, num)
		}
		const dimensions = 2
		if len(coords) != dimensions {
			return nil, fmt.Errorf("invalid line without 3 coordinates: %v", line)
		}
		points = append(points, Point{
			X: coords[0],
			Y: coords[1],
		})
	}
	return points, nil
}
