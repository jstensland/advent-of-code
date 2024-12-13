package day11_test

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2024/day11"
)

func exampleEasyIn() io.Reader {
	return strings.NewReader("0 1 10 99 999")
}

func TestPart1ExampleEasy(t *testing.T) {
	stoneLine, err := day11.ParseInput(exampleEasyIn())
	require.NoError(t, err)

	stoneLine.Blink()

	assert.Equal(t, "1 2024 1 0 9 9 2021976", stoneLine.String())
}

func exampleIn() io.Reader {
	return strings.NewReader("125 17")
}

func TestPart1ExampleSteps(t *testing.T) {
	stoneLine, err := day11.ParseInput(exampleIn())
	require.NoError(t, err)

	assert.Equal(t, "125 17", stoneLine.String())
	stoneLine.Blink()
	assert.Equal(t, "253000 1 7", stoneLine.String())
	stoneLine.Blink()
	assert.Equal(t, "253 0 2024 14168", stoneLine.String())
	stoneLine.Blink()
	assert.Equal(t, "512072 1 20 24 28676032", stoneLine.String())
	stoneLine.Blink()
	assert.Equal(t, "512 72 2024 2 0 2 4 2867 6032", stoneLine.String())
	stoneLine.Blink()
	assert.Equal(t, "1036288 7 2 20 24 4048 1 4048 8096 28 67 60 32", stoneLine.String())
	stoneLine.Blink()
	assert.Equal(t, "2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2", stoneLine.String())
}

func TestPart1Example(t *testing.T) {
	answer, err := day11.SolvePart1(exampleIn())

	require.NoError(t, err)
	assert.Equal(t, 55312, answer)
}

func TestPart1Input(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	answer, err := day11.SolvePart1(in)

	require.NoError(t, err)
	assert.Equal(t, 185894, answer) // confirmed
}

func TestPart2Input(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	answer, err := day11.SolvePart2(in)

	require.NoError(t, err)
	assert.Equal(t, 221632504974231, answer) // ?
}

func TestPart2ExampleSteps(t *testing.T) {
	testCases := []struct {
		rounds int
		freq   map[day11.StoneNumber]int64
		stones int64
	}{
		{0, map[day11.StoneNumber]int64{125: 1, 17: 1}, 2},
		{1, map[day11.StoneNumber]int64{253000: 1, 1: 1, 7: 1}, 3},
		{2, map[day11.StoneNumber]int64{253: 1, 0: 1, 2024: 1, 14168: 1}, 4},
		{3, map[day11.StoneNumber]int64{512072: 1, 1: 1, 20: 1, 24: 1, 28676032: 1}, 5},
		{4, map[day11.StoneNumber]int64{512: 1, 72: 1, 2024: 1, 2: 2, 0: 1, 4: 1, 2867: 1, 6032: 1}, 9},
		{5, map[day11.StoneNumber]int64{
			1036288: 1, 7: 1, 2: 1, 20: 1, 24: 1, 4048: 2, 1: 1, 8096: 1, 28: 1, 67: 1, 60: 1, 32: 1,
		}, 13},
		{6, map[day11.StoneNumber]int64{
			2097446912: 1, 14168: 1, 4048: 1, 0: 2, 2: 4, 4: 1, 40: 2, 48: 2,
			2024: 1, 80: 1, 96: 1, 8: 1, 6: 2, 7: 1, 3: 1,
		}, 22},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d rounds", tc.rounds), func(t *testing.T) {
			stoneLine, err := day11.ParseInput(exampleIn())
			require.NoError(t, err)
			stoneSet := day11.NewStoneSet(stoneLine)

			for range tc.rounds {
				stoneSet.Blink()
			}

			if !reflect.DeepEqual(tc.freq, stoneSet.Data) {
				t.Errorf("Maps are not equal: %v != %v", tc.freq, stoneSet.Data)
			}

			require.NoError(t, err)
			assert.Equal(t, tc.stones, stoneSet.Length(), "number of stones is unexpected")
		})
	}
}

func TestPart2Example(t *testing.T) {
	answer, err := day11.SolvePart2Rounds(exampleIn(), 25)

	require.NoError(t, err)
	assert.Equal(t, int64(55312), answer)
}
