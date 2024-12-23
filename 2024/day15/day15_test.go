package day15_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2024/day15"
)

func example() io.Reader {
	return strings.NewReader(`##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`)
}

func smallExample() io.Reader {
	return strings.NewReader(`########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`)
}

func TestParseSmallExample(t *testing.T) {
	grid, err := day15.ParseIn(smallExample(), 1, 0)

	require.NoError(t, err)
	assert.Equal(t, 8, grid.Width)
	assert.Equal(t, 8, grid.Height)
	assert.Equal(t, 2, grid.RobotLocationV2().Row)
	assert.Equal(t, 2, grid.RobotLocationV2().Col)
	assert.Len(t, grid.Moves(), 15)
}

func TestParseExample(t *testing.T) {
	grid, err := day15.ParseIn(example(), 1, 0)

	require.NoError(t, err)
	assert.Equal(t, 10, grid.Width)
	assert.Equal(t, 10, grid.Height)
	assert.Equal(t, 4, grid.RobotLocationV2().Row)
	assert.Equal(t, 4, grid.RobotLocationV2().Col)
	assert.Len(t, grid.Moves(), 700)
}

func TestExampleTotalGPS(t *testing.T) {
	grid, err := day15.ParseIn(strings.NewReader(`##########
#.O.O.OOO#
#........#
#OO......#
#OO@.....#
#O#.....O#
#O.....OO#
#O.....OO#
#OO....OO#
##########`), 1, 0)

	require.NoError(t, err)
	assert.Equal(t, 10092, grid.TotalGPS())
}

func TestSmallExample(t *testing.T) {
	total, err := day15.SolvePart1(smallExample())

	require.NoError(t, err)
	assert.Equal(t, 2028, total)
}

func TestPart1Example(t *testing.T) {
	total, err := day15.SolvePart1(example())

	require.NoError(t, err)
	assert.Equal(t, 10092, total)
}

func TestPart1(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	total, err := day15.SolvePart1(in)

	require.NoError(t, err)
	assert.Equal(t, 1446158, total)
}

func TestPart2Moves(t *testing.T) {
	board1 := `#####
#...#
#.@.#
#...#
#####`
	board2 := `######
#....#
#.@O.#
#....#
######`
	board3 := `#######
#.....#
#.OO@.#
#.....#
#######`

	board4 := `#######
#.....#
#@OOO.#
#.....#
#######`

	board5 := `#####
#...#
#.O.#
#.O.#
#.@.#
#...#
#####`
	board6 := `#####
#.@.#
#...#
#.O.#
#.O.#
#...#
#####`

	board7 := `#####
#...#
#.O.#
#.OO#
#.O.#
#.@.#
#...#
#####`

	board8 := `#######
#.....#
#.....#
#.OO@.#
#.O...#
#.....#
#.....#
#######`

	testCases := []struct {
		desc          string
		startingBoard string
		moves         string
		robotLoc      day15.Location
	}{
		{
			desc:          "open left",
			startingBoard: board1,
			moves:         "<",
			robotLoc:      day15.Location{2, 3},
		},
		{
			desc:          "open up",
			startingBoard: board1,
			moves:         "^",
			robotLoc:      day15.Location{1, 4},
		},
		{
			desc:          "into walls in various directions",
			startingBoard: board1,
			moves:         "^^^>>>>v<",
			robotLoc:      day15.Location{2, 6},
		},
		{
			desc:          "push box right",
			startingBoard: board2,
			moves:         ">>>>",
			robotLoc:      day15.Location{2, 7},
		},
		{
			desc:          "push two boxes left",
			startingBoard: board3,
			moves:         "<<<<",
			robotLoc:      day15.Location{2, 6},
		},
		{
			desc:          "push three boxes right",
			startingBoard: board4,
			moves:         ">>>>>",
			robotLoc:      day15.Location{2, 5},
		},
		{
			desc:          "push aligned boxes up",
			startingBoard: board5,
			moves:         "^>^<^",
			robotLoc:      day15.Location{3, 4},
		},
		{
			desc:          "push aligned boxes down",
			startingBoard: board6,
			moves:         "v>vvvv",
			robotLoc:      day15.Location{3, 5},
		},
		{
			desc:          "push boxes up not aligned",
			startingBoard: board7,
			moves:         "<^>v>^^",
			robotLoc:      day15.Location{4, 5},
		},
		{
			desc:          "push boxes down, not aligned",
			startingBoard: board8,
			moves:         "<^<<v<<<",
			robotLoc:      day15.Location{3, 4},
		},

		// Add tests for down, not aligned
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			in := strings.NewReader(tC.startingBoard + "\n\n" + tC.moves + "\n")
			grid, err := day15.ParseIn(in, 2, 1)
			require.NoError(t, err)

			grid.RunRobotsV2()

			assert.Equal(t, tC.robotLoc.Row, grid.RobotLocationV2().Row)
			assert.Equal(t, tC.robotLoc.Col, grid.RobotLocationV2().Col)
		})
	}
}

func TestExampleIn(t *testing.T) {
	total, err := day15.SolvePart2(example())

	require.NoError(t, err)
	assert.Equal(t, 9021, total)
	// ####################
	// ##[].......[].[][]##
	// ##[]...........[].##
	// ##[]........[][][]##
	// ##[]......[]....[]##
	// ##..##......[]....##
	// ##..[]............##
	// ##..@......[].[][]##
	// ##......[][]..[]..##
	// ####################
}

func TestExampleTotalGPSV2(_ *testing.T) {}

func TestPart2(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	total, err := day15.SolvePart2(in)

	require.NoError(t, err)
	// 1461173 is too high
	assert.Equal(t, 1461173, total) // WRONG!
}
