package day7_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2024/day7"
)

func TestPart1Input(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	answer, err := day7.RunPart1(in)

	require.NoError(t, err)
	assert.Equal(t, 5837374519342, answer) //?
}

func exampleIn() io.Reader {
	return strings.NewReader(`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`)
}

func TestPart1Example(t *testing.T) {
	answer, err := day7.RunPart1(exampleIn())

	require.NoError(t, err)
	assert.Equal(t, 3749, answer)
}

func TestPart2Example(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	answer, err := day7.RunPart2(in)

	require.NoError(t, err)
	assert.Equal(t, 492383931650959, answer)
}

func TestPermsOne(t *testing.T) {
	perms := day7.Perms(1, []day7.BinaryOp{day7.Add, day7.Multiple})

	assert.Len(t, perms, 2)

	// for _, perm := range perms {
	// 	fmt.Println("perm:", perm)
	// }
}

func TestPermsTwo(t *testing.T) {
	expected := [][]day7.BinaryOp{
		{day7.Add, day7.Add},
		{day7.Multiple, day7.Add},
		{day7.Add, day7.Multiple},
		{day7.Multiple, day7.Multiple},
	}

	perms := day7.Perms(2, []day7.BinaryOp{day7.Add, day7.Multiple})

	assert.Len(t, perms, 4)

	for i, perm := range perms {
		// fmt.Println("perm:", perm)

		for j, op := range perm {
			assert.Equal(t, expected[i][j].String(), op.String())
		}
	}
}

func TestPermsThree(t *testing.T) {
	expected := [][]day7.BinaryOp{
		{day7.Add, day7.Add, day7.Add},
		{day7.Multiple, day7.Add, day7.Add},
		{day7.Add, day7.Multiple, day7.Add},
		{day7.Multiple, day7.Multiple, day7.Add},
		{day7.Add, day7.Add, day7.Multiple},
		{day7.Multiple, day7.Add, day7.Multiple},
		{day7.Add, day7.Multiple, day7.Multiple},
		{day7.Multiple, day7.Multiple, day7.Multiple},
	}

	perms := day7.Perms(3, []day7.BinaryOp{day7.Add, day7.Multiple})

	assert.Len(t, perms, 8)

	for i, perm := range perms {
		// fmt.Println("perm:", perm)

		for j, op := range perm {
			assert.Equal(t, expected[i][j].String(), op.String())
		}
	}
}
