package day13_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2024/day13"
)

func example() io.Reader {
	return strings.NewReader(`Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`)
}

func TestParseExample(t *testing.T) {
	games, err := day13.ParseIn(example())

	require.NoError(t, err)
	assert.Len(t, games, 4)

	// check the second game
	assert.Equal(t, 26, games[1].A.XDelta)
	assert.Equal(t, 66, games[1].A.YDelta)
	assert.Equal(t, 67, games[1].B.XDelta)
	assert.Equal(t, 21, games[1].B.YDelta)
	assert.Equal(t, day13.Coordinate{12748, 12176}, games[1].Prize)

	// check the last game
	assert.Equal(t, 69, games[3].A.XDelta)
	assert.Equal(t, 23, games[3].A.YDelta)
	assert.Equal(t, 27, games[3].B.XDelta)
	assert.Equal(t, 71, games[3].B.YDelta)
	assert.Equal(t, day13.Coordinate{18641, 10279}, games[3].Prize)
}

func TestSolveUnsolveable(t *testing.T) {
	gameTooLow := day13.Game{
		A: day13.Button{
			XDelta: 1,
			YDelta: 2,
		},
		B: day13.Button{
			XDelta: 1,
			YDelta: 4,
		},
		Prize: day13.Coordinate{
			X: 5,
			Y: 25,
		},
	}

	cost, err := gameTooLow.Solve()

	require.ErrorIs(t, err, day13.ErrUnsolvable)
	assert.Equal(t, 0, cost)

	gameTooHigh := day13.Game{
		A: day13.Button{
			XDelta: 1,
			YDelta: 1,
		},
		B: day13.Button{
			XDelta: 4,
			YDelta: 3,
		},
		Prize: day13.Coordinate{
			X: 30,
			Y: 1,
		},
	}
	cost, err = gameTooHigh.Solve()

	require.ErrorIs(t, err, day13.ErrUnsolvable)
	assert.Equal(t, 0, cost)
}

func TestExampleGame1(t *testing.T) {
	gameTooHigh := day13.Game{
		A: day13.Button{
			Cost:   day13.ACost,
			XDelta: 94,
			YDelta: 34,
		},
		B: day13.Button{
			Cost:   day13.BCost,
			XDelta: 22,
			YDelta: 67,
		},
		Prize: day13.Coordinate{
			X: 8400,
			Y: 5400,
		},
	}
	cost, err := gameTooHigh.Solve()

	require.NoError(t, err)
	assert.Equal(t, 80, gameTooHigh.A.Count)
	assert.Equal(t, 40, gameTooHigh.B.Count)
	assert.Equal(t, 280, cost)
}

func TestRunPart1Example(t *testing.T) {
	total, err := day13.SolvePart1(example())

	require.NoError(t, err)
	assert.Equal(t, 480, total)
}

func TestRunPart1(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	total, err := day13.SolvePart1(in)

	require.NoError(t, err)
	assert.Equal(t, 40069, total) // confirmed
}

func TestRunPart2Example(t *testing.T) {
	total, err := day13.SolvePart2(example())

	require.NoError(t, err)
	assert.Equal(t, 875318608908, total)
}

func TestRunPart2(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	total, err := day13.SolvePart2(in)

	require.NoError(t, err)
	assert.Equal(t, 71493195288102, total) // confirmed
}
