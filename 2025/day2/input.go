package day2

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

func ParseIn(r io.Reader) ([]Range, error) {
	scanner := bufio.NewScanner(r)

	// Read first line
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("scanning input: %w", err)
		}
		return nil, errors.New("no input found")
	}

	line := strings.TrimSpace(scanner.Text())

	// Check if there's more than one line
	if scanner.Scan() {
		return nil, errors.New("expected only one line of input, found multiple")
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanning for additional lines: %w", err)
	}

	// Split by comma to get individual ranges
	rawRanges := strings.Split(line, ",")
	ranges := make([]Range, 0, len(rawRanges))

	const numRangeParts = 2
	for _, rawRange := range rawRanges {
		// Split by dash to get start and end
		rangeParts := strings.Split(strings.TrimSpace(rawRange), "-")
		if len(rangeParts) != numRangeParts {
			continue // Skip malformed ranges
		}

		// Convert each digit to an int for start
		starts := make([]int, 0, len(rangeParts[0]))
		for _, digit := range rangeParts[0] {
			starts = append(starts, int(digit-'0')) // convert numeric rune to int equivalent
		}

		// Convert each digit to an int for end
		ends := make([]int, 0, len(rangeParts[1]))
		for _, digit := range rangeParts[1] {
			ends = append(ends, int(digit-'0'))
		}

		ranges = append(ranges, Range{Start: starts, End: ends})
	}

	return ranges, nil
}
