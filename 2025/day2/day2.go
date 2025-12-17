// Package day2 solves AoC 2025 day 2
package day2

import (
	"fmt"
	"io"
)

// Range is a go representation of the input.
type Range struct {
	Start ID
	End   ID
}

func Part1(r io.Reader) (int, error) {
	ranges, err := ParseIn(r)
	if err != nil {
		return 0, err
	}

	ids := []ID{}
	for _, ran := range ranges {
		invalids, err := InvalidIDs(ran)
		if err != nil {
			return 0, fmt.Errorf("error determining invalid IDs: %w", err)
		}
		ids = append(ids, invalids...)
	}

	return sumIDs(ids), nil
}

// InvalidIDs returns all invalid IDs for a given range.
//
// Invalid IDs are IDs which are made only of some sequence of digits repeated twice.
// e.g. 1414, 88, 123123, 87798779.
func InvalidIDs(inRange Range) ([]ID, error) {
	firstInvalid := findFirst(inRange.Start)
	out := []ID{}
	end := inRange.End.AsInt()
	candidate := firstInvalid
	for ID(candidate).AsInt() <= end {
		out = append(out, ID(candidate))
		candidate = candidate.NextInvalid()
	}

	return out, nil
}

func Part2(r io.Reader) (int, error) {
	_ = r
	return 0, nil
}

// func Part2(r io.Reader) (int, error) {
// 	ranges, err := ParseIn(r)
// 	if err != nil {
// 		return 0, err
// 	}
// 	ids := []ID{}
// 	for _, ran := range ranges {
// 		invalids, err := InvalidIDsV2(ran)
// 		if err != nil {
// 			return 0, fmt.Errorf("error determining invalid IDs: %w", err)
// 		}
// 		ids = append(ids, invalids...)
// 	}
//
// 	return sumIDs(ids), nil
// }

// InvalidIDsV2 returns all invalid IDs for a given range.
//
// Invalid IDs are IDs which are made only of some sequence of digits repeated
// some number of times. For example
// 12341234 (1234 two times)
// 123123123 (123 three times)
// 1212121212 (12 five times)
// 1111111 (1 seven times)
// are all invalid IDs.
func InvalidIDsV2(inRange Range) ([]ID, error) {
	firstInvalid := findFirstV2(inRange.Start)
	out := []ID{}
	end := inRange.End.AsInt()
	candidate := firstInvalid
	for ID(candidate).AsInt() <= end {
		out = append(out, ID(candidate))
		candidate = candidate.NextInvalid()
	}

	return out, nil
}

func sumIDs(in []ID) int {
	total := 0
	for _, i := range in {
		total += i.AsInt()
	}
	return total
}
