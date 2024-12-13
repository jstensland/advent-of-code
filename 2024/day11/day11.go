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

// SolvePart2 solves without accounting for order
func SolvePart2(in io.Reader, rounds int) (int64, error) {
	stoneLine, err := ParseInput(in)
	if err != nil {
		return 0, fmt.Errorf("error parsing input file: %w", err)
	}

	stoneSet := NewStoneSet(stoneLine)
	i := 0
	for range rounds {
		i++
		stoneSet.Blink()
	}
	return stoneSet.Length(), nil
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
//
// Much slower... can make it in the 40s pretty easily, but then really bogs down
// keeps the order of the stones by keeping them all in a slice which grows exponentially
func (sl *StoneLine) Blink() {
	out := StoneLine{}
	for _, stoneVal := range *sl {
		switch {
		case stoneVal == 0:
			out = append(out, 1)
		case len(strconv.Itoa(stoneVal))%2 == 0:
			left, right := Split(stoneVal)
			out = append(out, left, right)
		default:
			out = append(out, stoneVal*2024) //nolint:mnd // value meaningless and defined for me. It magic.
		}
	}
	*sl = out
}

func (sl *StoneLineSet) String() string {
	out := ""
	for stone, count := range sl.Data {
		out += fmt.Sprintf("%d:%d ", stone, count)
	}
	return out
}

func (sl *StoneLineSet) Blink() {
	// make a copy to not affect the current one
	original := make(map[StoneNumber]int64, len(sl.Data))
	for k, v := range sl.Data {
		original[k] = v
	}

	for stone, count := range original {
		switch {
		case stone == 0:
			sl.Sub(stone, count) // changing all original 0s to 1s
			sl.Add(1, count)
		case len(strconv.FormatInt(int64(stone), 10))%2 == 0:
			sl.Sub(stone, count) // change all originals
			// add each split
			left, right := SplitStone(stone)
			sl.Add(left, count)
			sl.Add(right, count)
		default:
			sl.Sub(stone, count)      // changing this one
			sl.Add(stone*2024, count) //nolint:mnd // value meaningless and defined for me. It magic.
		}
	}
}

func (sl *StoneLineSet) Sub(stone StoneNumber, num int64) {
	if count, ok := sl.Data[stone]; ok {
		if num > count {
			// should not happen
			panic(fmt.Sprintf("removing more of stone %v than exists: %d of %d\n", stone, num, count))
		}

		if num >= count {
			delete(sl.Data, stone)
			return
		}
		sl.Data[stone] -= num
		return
	}
	sl.Data[stone] = 1
}

func (sl *StoneLineSet) Add(stone StoneNumber, num int64) {
	if _, ok := sl.Data[stone]; ok {
		// count += num // does this work or not stored?
		sl.Data[stone] += num
		return
	}
	sl.Data[stone] = num
}

func (sl *StoneLineSet) Length() int64 {
	out := int64(0)
	for _, count := range sl.Data {
		out += count
	}
	return out
}

// SplitStone modifies the slice by taking the stone at the given index. and splitting it into two stones.
func SplitStone(val StoneNumber) (StoneNumber, StoneNumber) {
	strVal := strconv.FormatInt(int64(val), 10)
	// The left half of the digits are engraved on the new left stone
	left, err := strconv.ParseInt(strVal[:len(strVal)/2], 10, 64)
	if err != nil {
		panic(err)
	}
	right, err := strconv.ParseInt(strVal[len(strVal)/2:], 10, 64)
	if err != nil {
		panic(err)
	}

	return StoneNumber(left), StoneNumber(right)
}

// Split modifies the slice by taking the stone at the given index. and splitting it into two stones.
func Split(val int) (int, int) {
	strVal := strconv.Itoa(val)
	// The left half of the digits are engraved on the new left stone
	// IMPROVEMENT: better way to split in half?
	left, err := strconv.Atoi(strVal[:len(strVal)/2])
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

type StoneNumber int64

type StoneLineSet struct {
	Data map[StoneNumber]int64
}

func NewStoneSet(sl *StoneLine) *StoneLineSet {
	if sl == nil {
		return nil
	}
	out := StoneLineSet{
		Data: map[StoneNumber]int64{},
	}
	for _, stone := range *sl {
		out.Add(StoneNumber(stone), 1)
	}
	return &out
}
