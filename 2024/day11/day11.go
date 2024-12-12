// Package day11 solves AoC 2024 day 10.
package day11

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

func SolvePart1(in io.Reader) (int, error) {
	stoneLine, err := ParseInput(in)
	if err != nil {
		return 0, fmt.Errorf("error parsing input file: %w", err)
	}

	blinks := 25
	for range blinks {
		stoneLine.Blink()
	}
	return len(*stoneLine), nil
}

func SolvePart2(in io.Reader) (int, error) {
	stoneLine, err := ParseInput(in)
	if err != nil {
		return 0, fmt.Errorf("error parsing input file: %w", err)
	}
	blinks := 75
	// Much slower... can make it in the 40s pretty easily, but then really bogs down
	i := 0
	for range blinks {
		i++
		fmt.Println("Blink", i)
		stoneLine.Blink()
	}
	return len(*stoneLine), nil
}

type StoneLine []int

func (sl *StoneLine) String() string {
	if sl == nil {
		return ""
	}
	var out string
	for _, stone := range *sl {
		out += strconv.Itoa(stone) + " "
	}
	// remove the last space
	out = out[:len(out)-1]
	return out
}

// Blink performs one iteration on all stones.
//   - If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
//   - If the stone is engraved with a number that has an even number of digits, it is replaced by two
//     stones. The left half of the digits are engraved on the new left stone, and the right half of
//     the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes:
//     1000 would become stones 10 and 0.)
//   - If none of the other rules apply, the stone is replaced by a new stone; the old stone's number
//     multiplied by 2024 is engraved on the new stone./
func (sl *StoneLine) Blink() {
	out := StoneLine{}
	for _, stoneVal := range *sl {
		// if i+1 >= len(out) {
		// 	fmt.Println("growing")
		// 	out = slices.Grow(out, offset)
		// }
		switch {
		case stoneVal == 0:
			out = append(out, 1)
		case len(strconv.Itoa(stoneVal))%2 == 0:
			// This is changing our result... how to ensure
			// the updated values aren't iterated on in the same iteration?
			// collect answers at this level in a new slice... or...

			left, right := Split(stoneVal)
			// out[i] = left
			// out[i+1] = right
			// offset++
			out = append(out, left, right)

		default:
			// (*sl)[i+offset] = stoneVal * 2024 //nolint:mnd // value meaningless and defined for me. It magic.
			out = append(out, stoneVal*2024) //nolint:mnd // value meaningless and defined for me. It magic.
		}
	}
	*sl = out
}

// Split modifies the slice by taking the stone at the given index. and splitting it into two stones.
func Split(val int) (int, int) {
	strVal := strconv.Itoa(val)
	// The left half of the digits are engraved on the new left stone
	left, err := strconv.Atoi(strVal[:len(strVal)/2]) // TODO: look at ways to split in half
	if err != nil {
		panic(err)
	}
	right, err := strconv.Atoi(strVal[len(strVal)/2:])
	if err != nil {
		panic(err)
	}

	return left, right
}

func ParseInput(in io.Reader) (*StoneLine, error) {
	raw, err := io.ReadAll(in)
	if err != nil {
		return nil, fmt.Errorf("error reading input: %w", err)
	}
	stones := strings.Split(strings.TrimSpace(string(raw)), " ")
	out := make(StoneLine, len(stones))
	for i, strVal := range stones {
		val, err := strconv.Atoi(strVal)
		if err != nil {
			return nil, fmt.Errorf("non int value for stone: %w", err)
		}
		out[i] = val
	}

	return &out, nil
}
