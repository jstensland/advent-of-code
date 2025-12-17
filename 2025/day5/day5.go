// Package day5 solves AoC 2025 day 5
package day5

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Part1(r io.Reader) (int, error) {
	spans, items, err := ParseIn(r)
	if err != nil {
		return 0, err
	}
	return countFresh(CombineRanges(spans), items), nil
}

func Part2(r io.Reader) (int, error) {
	spans, _, err := ParseIn(r)
	if err != nil {
		return 0, err
	}
	return countIDs(CombineRanges(spans)), nil
}

func ParseIn(r io.Reader) ([]Span, []int, error) {
	scanner := bufio.NewScanner(r)
	ranges := []Span{}

	// process the first section
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}

		const spanSegments = 2
		pieces := strings.Split(line, "-")
		if len(pieces) != spanSegments {
			return nil, nil, fmt.Errorf("unexpected range line: %v", line)
		}
		start, err := strconv.Atoi(pieces[0])
		if err != nil {
			return nil, nil, fmt.Errorf("unexpected conversion error: %w", err)
		}
		end, err := strconv.Atoi(pieces[1])
		if err != nil {
			return nil, nil, fmt.Errorf("unexpected conversion error: %w", err)
		}
		ranges = append(ranges, Span{Start: start, End: end})
	}

	// process the items
	items := []int{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		item, err := strconv.Atoi(line)
		if err != nil {
			return nil, nil, fmt.Errorf("unexpected conversion error: %w", err)
		}
		items = append(items, item)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("scanner error on input: %w", err)
	}

	return ranges, items, nil
}
