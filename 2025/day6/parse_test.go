package day6_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2025/day6"
)

func TestParseIn(t *testing.T) {
	input := example1()
	sheet, err := day6.ParseIn(bytes.NewReader([]byte(input)))

	require.NoError(t, err)
	require.NotNil(t, sheet)

	// Verify we have 4 problems (one per column)
	problems := sheet.Problems()
	require.Len(t, problems, 4, "should have 4 problems (one per column)")

	// Verify column 0: [123, 45, 6] with multiply operator
	require.Equal(t, []int{123, 45, 6}, problems[0].Operands())
	require.Equal(t, day6.Multiple, problems[0].Operator())

	// Verify column 1: [328, 64, 98] with add operator
	require.Equal(t, []int{328, 64, 98}, problems[1].Operands())
	require.Equal(t, day6.Add, problems[1].Operator())

	// Verify column 2: [51, 387, 215] with multiply operator
	require.Equal(t, []int{51, 387, 215}, problems[2].Operands())
	require.Equal(t, day6.Multiple, problems[2].Operator())

	// Verify column 3: [64, 23, 314] with add operator
	require.Equal(t, []int{64, 23, 314}, problems[3].Operands())
	require.Equal(t, day6.Add, problems[3].Operator())
}

func TestParsePart2(t *testing.T) {
	input := example1()
	sheet, err := day6.ParseInPart2(bytes.NewReader([]byte(input)))

	require.NoError(t, err)
	require.NotNil(t, sheet)

	problems := sheet.Problems()
	// require.Len(t, problems, 4, "should have 4 problems (one per column)")
	// The rightmost problem is 4 + 431 + 623 = 1058
	require.Equal(t, []int{4, 431, 623}, problems[0].Operands())
	require.Equal(t, day6.Add, problems[0].Operator())
	// The second problem from the right is 175 * 581 * 32 = 3253600
	require.Equal(t, []int{175, 581, 32}, problems[1].Operands())
	require.Equal(t, day6.Multiple, problems[1].Operator())
	// The third problem from the right is 8 + 248 + 369 = 625
	require.Equal(t, []int{8, 248, 369}, problems[2].Operands())
	require.Equal(t, day6.Add, problems[2].Operator())
	// Finally, the leftmost problem is 356 * 24 * 1 = 8544
	require.Equal(t, []int{356, 24, 1}, problems[3].Operands())
	require.Equal(t, day6.Multiple, problems[3].Operator())
}
