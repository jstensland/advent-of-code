package day14_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2024/day14"
)

func example() io.Reader {
	return strings.NewReader(`p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`)
}

func TestParseExample(t *testing.T) {
	grid, err := day14.ParseIn(example(), 7, 11)

	require.NoError(t, err)
	assert.Len(t, grid.Robots, 12)
	// check the first robot
	assert.Equal(t, day14.Position{0, 4}, grid.Robots[0].Position)
	assert.Equal(t, day14.Velocity{3, -3}, grid.Robots[0].Velocity)
}

// p=2,4 v=2,-3

func TestTickRobot1(t *testing.T) {
	grid := day14.Grid{
		Width:  11,
		Height: 7,
		Robots: []day14.Robot{
			{
				ID:       day14.RobotID(1),
				Position: day14.Position{2, 4},
				Velocity: day14.Velocity{2, -3},
			},
		},
	}

	grid.Tick()
	assert.Equal(t, day14.Position{4, 1}, grid.Robots[0].Position, "wrong position after one tick")
	grid.Tick()
	assert.Equal(t, day14.Position{6, 5}, grid.Robots[0].Position, "wrong position after one tick")
	grid.Tick()
	assert.Equal(t, day14.Position{8, 2}, grid.Robots[0].Position, "wrong position after one tick")
	grid.Tick()
	assert.Equal(t, day14.Position{10, 6}, grid.Robots[0].Position, "wrong position after one tick")
	grid.Tick()
	assert.Equal(t, day14.Position{1, 3}, grid.Robots[0].Position, "wrong position after one tick")
}

func TestSafetyFactor(t *testing.T) {
	grid := day14.Grid{
		Width:  11,
		Height: 7,
		Robots: []day14.Robot{
			{Position: day14.Position{Col: 2, Row: 3}},  // row 3 is ignored
			{Position: day14.Position{Col: 5, Row: 4}},  // col 5 is ignored
			{Position: day14.Position{Col: 4, Row: 2}},  // top left
			{Position: day14.Position{Col: 2, Row: 2}},  // top left
			{Position: day14.Position{Col: 10, Row: 0}}, // top right
			{Position: day14.Position{Col: 6, Row: 4}},  // bottom right
			{Position: day14.Position{Col: 4, Row: 5}},  // bottom left
			{Position: day14.Position{Col: 4, Row: 5}},  // bottom left
		},
	}

	assert.Equal(t, 4, grid.SafetyFactor())
}

func TestExampleParts(t *testing.T) {
	grid, err := day14.ParseIn(example(), 7, 11)
	require.NoError(t, err)

	for range 100 {
		grid.Tick()
	}

	robotPositions := make([]day14.Position, 0, len(grid.Robots))
	for _, robot := range grid.Robots {
		robotPositions = append(robotPositions, robot.Position)
	}

	// fmt.Println(grid)
	// ......2..1.
	// ...........
	// 1..........
	// .11........
	// .....1.....
	// ...12......
	// .1....1....
	expectedPositions := []day14.Position{
		{6, 0}, {6, 0}, {9, 0}, {0, 2}, {1, 3}, {2, 3}, {5, 4}, {3, 5}, {4, 5}, {4, 5}, {1, 6}, {6, 6},
	}
	assert.ElementsMatch(t, expectedPositions, robotPositions, "listA is expected positions")
	assert.Equal(t, 12, grid.SafetyFactor())
}

func TestExampleIn(t *testing.T) {
	safetyFactor, err := day14.SolvePart1(example(), 7, 11)

	require.NoError(t, err)
	assert.Equal(t, 12, safetyFactor)
}

func TestSolvePart1(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	safetyFactor, err := day14.SolvePart1(in, 103, 101)

	require.NoError(t, err)
	assert.Equal(t, 226548000, safetyFactor)
}

func TestTreeLikeSymmetric(t *testing.T) {
	grid := day14.Grid{
		Width:  11,
		Height: 7,
		Robots: []day14.Robot{
			{Position: day14.Position{Col: 0, Row: 0}},
			{Position: day14.Position{Col: 1, Row: 0}},
			{Position: day14.Position{Col: 9, Row: 0}},
			{Position: day14.Position{Col: 10, Row: 0}},
		},
	}
	assert.True(t, grid.TreeLikeSymmetric())
}

func TestTreeLikeManyInARow(t *testing.T) {
	grid := day14.Grid{
		Width:  11,
		Height: 7,
		Robots: []day14.Robot{
			{Position: day14.Position{Col: 2, Row: 3}},
			{Position: day14.Position{Col: 3, Row: 3}},
			{Position: day14.Position{Col: 4, Row: 3}},
			{Position: day14.Position{Col: 5, Row: 3}},
			{Position: day14.Position{Col: 6, Row: 3}},
		},
	}

	assert.True(t, grid.TreeLike(5), "should see 5 in a row")

	gridGap := day14.Grid{
		Width:  11,
		Height: 7,
		Robots: []day14.Robot{
			{Position: day14.Position{Col: 1, Row: 3}},
			{Position: day14.Position{Col: 2, Row: 3}},
			{Position: day14.Position{Col: 3, Row: 3}},
			{Position: day14.Position{Col: 5, Row: 3}},
			{Position: day14.Position{Col: 6, Row: 3}},
		},
	}

	assert.False(t, gridGap.TreeLike(5), "gap should break up sequence")
}

func TestSolvePart2(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	safetyFactor, err := day14.SolvePart2(in, 103, 101)

	require.NoError(t, err)
	assert.Equal(t, 7753, safetyFactor)
}
