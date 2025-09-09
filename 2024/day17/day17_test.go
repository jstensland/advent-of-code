package day17_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2024/day17"
)

func example() io.Reader {
	return strings.NewReader(`Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`)
}

func TestParseInExample(t *testing.T) {
	computer, err := day17.ParseIn(example())

	require.NoError(t, err)
	assert.Equal(t, 729, computer.RegisterA())
	assert.Equal(t, 0, computer.RegisterB())
	assert.Equal(t, 0, computer.RegisterC())
	assert.Equal(t, []uint8{0, 1, 5, 4, 3, 0}, computer.GetData())
}

func TestSolveExample(t *testing.T) {
	out, err := day17.SolvePart1(example())

	require.NoError(t, err)
	assert.Equal(t, "4,6,3,5,6,3,5,2,1,0", out)
}

func TestSolvePart1(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	out, err := day17.SolvePart1(in)

	require.NoError(t, err)
	assert.Equal(t, "3,7,1,7,2,1,0,6,3", out)
}

func TestAdv(t *testing.T) {
	c := day17.NewComputer(16, 0, 0, nil)

	day17.Adv(c, 0)
	assert.Equal(t, 16, c.RegisterA()) // 16/1
	day17.Adv(c, 1)
	assert.Equal(t, 8, c.RegisterA()) // 16/2
	day17.Adv(c, 2)
	assert.Equal(t, 2, c.RegisterA()) // 8/4

	c.SetRegisterA(24)
	c.SetRegisterB(3)
	day17.Adv(c, 5)                   // use register B
	assert.Equal(t, 3, c.RegisterA()) // 24/(2^3)

	c.SetRegisterA(15)
	c.SetRegisterC(2)
	day17.Adv(c, 6)                   // use register C
	assert.Equal(t, 3, c.RegisterA()) // 15/(2^2) = 3.75 -> 3
}

func TestBdv(t *testing.T) {
	c := day17.NewComputer(16, 0, 0, nil)

	day17.Bdv(c, 0)
	assert.Equal(t, 16, c.RegisterB()) // 16/1
	day17.Bdv(c, 1)
	assert.Equal(t, 8, c.RegisterB()) // 16/2

	c.SetRegisterA(8)
	day17.Bdv(c, 2)
	assert.Equal(t, 2, c.RegisterB()) // 8/4

	c.SetRegisterA(24)
	c.SetRegisterB(3)
	day17.Bdv(c, 5)                   // use register B
	assert.Equal(t, 3, c.RegisterB()) // 24/(2^3)

	c.SetRegisterA(15)
	c.SetRegisterC(2)
	day17.Bdv(c, 6)                   // use register C
	assert.Equal(t, 3, c.RegisterB()) // 15/(2^2) = 3.75 -> 3
}

func TestExampleInstruction1(t *testing.T) {
	c := day17.NewComputer(0, 0, 9, []uint8{2, 6})

	c.RunProgram()

	assert.Equal(t, 1, c.RegisterB())
}

func TestExampleInstruction2(t *testing.T) {
	c := day17.NewComputer(10, 0, 0, []uint8{5, 0, 5, 1, 5, 4})

	c.RunProgram()

	assert.Equal(t, "0,1,2", c.Result())
}

func TestExampleInstruction3(t *testing.T) {
	c := day17.NewComputer(2024, 0, 0, []uint8{0, 1, 5, 4, 3, 0})

	c.RunProgram()

	assert.Equal(t, "4,2,5,6,7,7,7,7,3,1,0", c.Result())
	assert.Equal(t, 0, c.RegisterA())
}

func TestExampleInstruction4(t *testing.T) {
	c := day17.NewComputer(0, 29, 0, []uint8{1, 7})

	c.RunProgram()

	assert.Equal(t, 26, c.RegisterB())
}

func TestExampleInstruction5(t *testing.T) {
	c := day17.NewComputer(0, 2024, 43690, []uint8{4, 0})

	c.RunProgram()

	assert.Equal(t, 44354, c.RegisterB())
}

func Part2Example() io.Reader {
	return strings.NewReader(`Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`)
}

func TestSolvePart2Example(t *testing.T) {
	// sanity checks
	computer, err := day17.ParseIn(Part2Example())
	require.NoError(t, err)
	assert.Equal(t, []uint8{0, 3, 5, 4, 3, 0}, computer.GetData())
	assert.Equal(t, "0,3,5,4,3,0", computer.Program.DataString())

	out, err := day17.SolvePart2BruteForce(Part2Example())
	require.NoError(t, err)

	// Answer
	assert.Equal(t, 117440, out)
}

func TestSolvePart2(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	out, err := day17.SolvePart2BruteForce(in) // TOO SLOW

	require.NoError(t, err)
	assert.Equal(t, "?", out)
}

func TestSolvePart2Example_3(t *testing.T) {
	// sanity checks
	computer, err := day17.ParseIn(Part2Example())
	require.NoError(t, err)
	assert.Equal(t, []uint8{0, 3, 5, 4, 3, 0}, computer.GetData())
	assert.Equal(t, "0,3,5,4,3,0", computer.Program.DataString())

	out, err := day17.SolvePart2Dynamic(Part2Example())
	require.NoError(t, err)

	// Answer
	assert.Equal(t, 117440, out)
}

// func TestSolvePart2_3(t *testing.T) {
// 	inFile := "./input.txt"
// 	in, err := os.Open(inFile)
// 	require.NoError(t, err)
//
// 	out, err := day17.SolvePart2BruteForce(in) // TOO SLOW
//
// 	require.NoError(t, err)
// 	assert.Equal(t, "?", out)
// }
