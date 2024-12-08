package day3_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jstensland/advent-of-code/2024/day3"
)

func TestPart1(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	answer, err := day3.SolvePart1(in)

	assert.NoError(t, err)
	assert.Equal(t, 175700056, answer) // confirmed
}

func TestPart2(t *testing.T) {
	inFile := "./input.txt"
	in, err := os.Open(inFile)
	require.NoError(t, err)

	answer, err := day3.SolvePart2(in)

	assert.NoError(t, err)
	// 102785526 // wrong answer when not taking into accoutn new line
	assert.Equal(t, 71668682, answer) //  X
}

func TestParseLine(t *testing.T) {
	testCases := []struct {
		name string
		line string
		want []day3.Op
		err  error
	}{
		{
			name: "basic",
			line: "mul(1,2)",
			want: []day3.Op{{Left: 1, Right: 2}},
			err:  nil,
		},
		{
			name: "padding",
			line: "xxxxmul(9,4)xxxx",
			want: []day3.Op{{Left: 9, Right: 4}},
			err:  nil,
		},
		{
			name: "extra mul",
			line: "mul(mul(9,4)",
			want: []day3.Op{{Left: 9, Right: 4}},
			err:  nil,
		},
		{
			name: "multiple different length ints",
			line: "mul(xxxmul(9,4)xxxxmul(20,245)xxx",
			want: []day3.Op{{Left: 9, Right: 4}, {Left: 20, Right: 245}},
			err:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ops, err := day3.ParseLine(tc.line)
			if err != tc.err {
				t.Fatalf("expected error %v, got %v", tc.err, err)
			}
			if len(ops) != len(tc.want) {
				t.Fatalf("expected %v, got %v", tc.want, ops)
			}
			for i, want := range tc.want {
				if ops[i] != want {
					t.Fatalf("expected %v, got %v", want, ops[i])
				}
			}
		})
	}
}

func TestParseLine2(t *testing.T) {
	testCases := []struct {
		name string
		line string
		want []day3.Op
		err  error
	}{
		{
			name: "basic",
			line: "mul(1,2)",
			want: []day3.Op{{Left: 1, Right: 2}},
			err:  nil,
		},
		{
			name: "padding",
			line: "xxxxmul(9,4)xxxx",
			want: []day3.Op{{Left: 9, Right: 4}},
			err:  nil,
		},
		{
			name: "extra mul",
			line: "mul(mul(9,4)",
			want: []day3.Op{{Left: 9, Right: 4}},
			err:  nil,
		},
		{
			name: "multiple different length ints",
			line: "mul(xxxmul(9,4)xxxxmul(20,245)xxx",
			want: []day3.Op{{Left: 9, Right: 4}, {Left: 20, Right: 245}},
			err:  nil,
		},
		{
			name: "don't then do",
			line: "mul(xdon't()xxmul(9,4)xxdo()xxmul(20,245)xxx",
			want: []day3.Op{{Left: 20, Right: 245}},
			err:  nil,
		},
		{
			name: "example from prompt",
			line: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			want: []day3.Op{{Left: 2, Right: 4}, {Left: 8, Right: 5}},
			err:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			parser := day3.NewParser2()
			ops, err := parser.ParseLine(tc.line)
			if err != tc.err {
				t.Fatalf("expected error %v, got %v", tc.err, err)
			}
			if len(ops) != len(tc.want) {
				t.Fatalf("expected %v, got %v", tc.want, ops)
			}
			for i, want := range tc.want {
				if ops[i] != want {
					t.Fatalf("expected %v, got %v", want, ops[i])
				}
			}
		})
	}
}

func TestParseOps2(t *testing.T) {
	in := strings.NewReader(`
		xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+
    mul(32,64](mul(11,8)undo()?mul(8,5))",
    `)
	ops := day3.ParseOps2(in)

	assert.Equal(t, []day3.Op{{2, 4}, {8, 5}}, ops)
}

func TestComputeExample(t *testing.T) {
	ops := []day3.Op{
		{Left: 2, Right: 4},
		{Left: 5, Right: 5},
		{Left: 11, Right: 8},
		{Left: 8, Right: 5},
	}

	assert.Equal(t, 161, day3.Compute(ops))
}
