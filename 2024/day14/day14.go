// Package day14 solves AoC 2024 day 14
package day14

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"slices"
	"strconv"
)

type Quadrant int

const (
	quadrant0 Quadrant = iota
	quadrant1
	quadrant2
	quadrant3
)

const treePicThreshold = 10

// RobotID uniquely identifies a robot. There are only 500.
type RobotID int

type Robot struct {
	ID       RobotID
	Position Position
	Velocity Velocity
}

type Grid struct {
	Robots []Robot
	Width  int
	Height int
}

type Position struct {
	// Col starts at zero and goes left to right
	Col int
	// Row starts at zero and goes top to bottom
	Row int
}

type Velocity struct {
	// DeltaCol is how many columns to the right to go if positive
	DeltaCol int
	// DeltaRow is how many rows to move down, if positive
	DeltaRow int
}

func SolvePart1(in io.Reader, height, width int) (int, error) {
	grid, err := ParseIn(in, height, width)
	if err != nil {
		return 0, fmt.Errorf("error loading input: %w", err)
	}

	for range 100 {
		grid.Tick()
	}
	return grid.SafetyFactor(), nil
}

// SolvePart2 find a tree...
//
// First I had height and width switched, so got discouraged. Looked up what a tree should
// look like. Tried searching outputs for long sequences of XXXXXX and that worked
func SolvePart2(in io.Reader, height, width int) (int, error) {
	grid, err := ParseIn(in, height, width)
	if err != nil {
		return 0, fmt.Errorf("error loading input: %w", err)
	}

	var seconds int
	i := 1
	for {
		grid.Tick()
		// if i%1_000 == 0 {
		// 	// show progress
		// 	fmt.Println("tick:", i)
		// }

		if grid.TreeLike(treePicThreshold) {
			// fmt.Println("tree like tick:", i) // see the answer
			// fmt.Println(grid) // see the trees
			seconds = i
			break
		}
		i++
	}
	return seconds, nil
}

func (g *Grid) Tick() {
	for i := range g.Robots {
		g.Robots[i].Position.Col = (g.Width + g.Robots[i].Position.Col + g.Robots[i].Velocity.DeltaCol) % g.Width
		g.Robots[i].Position.Row = (g.Height + g.Robots[i].Position.Row + g.Robots[i].Velocity.DeltaRow) % g.Height
	}
}

// IMPROVEMENT: look at other image processing approaches. Just looking for a lot of ### in a row for now

func (g *Grid) TreeLike(consecutive int) bool {
	// sort the robots by row and then column
	slices.SortFunc(g.Robots, func(a, b Robot) int {
		if a.Position.Row == b.Position.Row {
			return a.Position.Col - b.Position.Col
		}
		return a.Position.Row - b.Position.Row
	})

	count := 1
	prevPos := Position{0, 0}
	for _, robot := range g.Robots {
		// count columns in a row
		if prevPos.Row == robot.Position.Row && prevPos.Col == robot.Position.Col-1 {
			count++
		} else {
			// reset if not in a row
			count = 1
		}
		prevPos = robot.Position
		if count >= consecutive {
			return true
		}
	}
	return false
}

func (g *Grid) TreeLikeSymmetric() bool {
	// init grid
	grid := make([][]bool, g.Height)
	for row := range g.Height {
		grid[row] = make([]bool, g.Width)
	}

	// place them on a grid
	for _, robot := range g.Robots {
		grid[robot.Position.Row][robot.Position.Col] = true
	}

	// check the gird
	for row := range g.Height / 2 {
		for col := range g.Width {
			if grid[row][col] != grid[row][g.Width-col-1] {
				return false
			}
		}
	}
	return true
}

func (g *Grid) String() string {
	// sort the robots by row and then column
	slices.SortFunc(g.Robots, func(a, b Robot) int {
		if a.Position.Row == b.Position.Row {
			return a.Position.Col - b.Position.Col
		}
		return a.Position.Row - b.Position.Row
	})

	// output each row with a dot for no robot and an X if there is one
	out := ""

	robotIndex := 0
	row := 0
	for row < g.Height {
		col := 0
		for col < g.Width {
			if robotIndex < len(g.Robots) && g.Robots[robotIndex].Position.Row == row &&
				g.Robots[robotIndex].Position.Col == col {
				for robotIndex < len(g.Robots) &&
					g.Robots[robotIndex].Position.Row == row &&
					g.Robots[robotIndex].Position.Col == col {
					robotIndex++ // remove dups
				}
				out += "X"
			} else {
				out += "."
			}
			col++
		}
		out += "\n"
		row++
	}
	return out
}

// Quadrant returns different quadrants for each quadrant, or -1 if not in a quadrant
func (g *Grid) Quadrant(p Position) Quadrant {
	if p.Col == g.Width/2 || p.Row == g.Height/2 {
		return -1
	}

	if p.Col <= g.Width/2 {
		if p.Row <= g.Height/2 {
			return quadrant0
		}
		return quadrant1
	}
	if p.Row <= g.Height/2 {
		return quadrant2
	}
	return quadrant3
}

func (g *Grid) SafetyFactor() int {
	if g.Width%2 == 0 || g.Height%2 == 0 {
		// fmt.Println("width:", g.Width)
		// fmt.Println("height:", g.Height)
		panic("even dimensions not safe! panic!")
	}

	counts := make(map[Quadrant]int)
	for _, robot := range g.Robots {
		quad := g.Quadrant(robot.Position)
		if quad == -1 {
			continue
		}
		counts[g.Quadrant(robot.Position)]++
	}
	factor := 1
	for _, count := range counts {
		factor *= count
	}
	return factor
}

// ParseIn reads the input into a grid of robots
func ParseIn(in io.Reader, height, width int) (*Grid, error) {
	scanner := bufio.NewScanner(in)
	robotID := 0
	robots := []Robot{}

	// p=2,0 v=2,-1
	robotInfo := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
	for scanner.Scan() {
		robotID++
		line := scanner.Text()
		info := robotInfo.FindStringSubmatch(line)
		px, err := strconv.Atoi(info[1])
		if err != nil {
			return nil, fmt.Errorf("error parsing xDelta: %w", err)
		}
		py, err := strconv.Atoi(info[2])
		if err != nil {
			return nil, fmt.Errorf("error parsing yDelta: %w", err)
		}
		vx, err := strconv.Atoi(info[3])
		if err != nil {
			return nil, fmt.Errorf("error parsing xDelta: %w", err)
		}
		vy, err := strconv.Atoi(info[4])
		if err != nil {
			return nil, fmt.Errorf("error parsing yDelta: %w", err)
		}

		robots = append(robots, Robot{
			ID:       RobotID(robotID),
			Position: Position{px, py},
			Velocity: Velocity{vx, vy},
		})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &Grid{Robots: robots, Height: height, Width: width}, nil
}
