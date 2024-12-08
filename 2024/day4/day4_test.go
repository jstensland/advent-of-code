package day4_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2024/day4"
)

func smallGrid() io.ReadCloser {
	return io.NopCloser(strings.NewReader(`XMASX
XMASM
XMASA
XMASX`))
}

func TestRunPart1Small(t *testing.T) {
	total, err := day4.SolvePart1(smallGrid())

	require.NoError(t, err)
	assert.Equal(t, 4+1+1, total)
}

func TestRunPart1(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	total, err := day4.SolvePart1(in)

	require.NoError(t, err)
	assert.Equal(t, 2434, total)
}

func grid2Example() io.ReadCloser {
	return io.NopCloser(strings.NewReader(`.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`))
}

func TestRunPart2Example(t *testing.T) {
	total, err := day4.SolvePart2(grid2Example())

	require.NoError(t, err)
	assert.Equal(t, 9, total)
}

func grid2VariousCasesExample() io.ReadCloser {
	return io.NopCloser(strings.NewReader(`.M.M......
..A..MSMS.
.M.S.MAA.A
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.S.
..........`))
}

func TestRunPart2EdgeCases(t *testing.T) {
	total, err := day4.SolvePart2(grid2VariousCasesExample())

	require.NoError(t, err)
	assert.Equal(t, 7, total)
}

func TestRunPart2(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	total, err := day4.SolvePart2(in)

	require.NoError(t, err)
	assert.Equal(t, 1835, total)
}
