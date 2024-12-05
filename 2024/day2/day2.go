package day2

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"os"
	"slices"
	"strconv"
	"strings"
)

func RunPart1(inFile string) (int, error) {
	reader, err := os.Open(inFile)
	if err != nil {
		return 0, fmt.Errorf("failed to open file %s: %w", inFile, err)
	}
	defer reader.Close()

	safeCount := 0
	reports, err := loadReports(reader)
	if err != nil {
		return 0, fmt.Errorf("failed to load reports: %w", err)
	}

	for _, report := range reports {
		// for report := range readReports(reader) {
		if report.Safe() {
			safeCount++
		}
	}
	return safeCount, nil
}

func RunPart2(inFile string) (int, error) {
	reader, err := os.Open(inFile)
	if err != nil {
		return 0, fmt.Errorf("failed to open file %s: %w", inFile, err)
	}
	return Part2Analysis(reader)
}

func Part2Analysis(in io.ReadCloser) (int, error) {
	safeCount := 0
	reports, err := loadReports(in)
	if err != nil {
		return 0, fmt.Errorf("failed to load reports: %w", err)
	}
	defer in.Close()

	for _, report := range reports {
		if report.SafeDampened() {
			safeCount++
		}
	}
	return safeCount, nil
}

type Report struct {
	// levels is an ordered list of levels read in the report
	levels []int
}

// Levels returns a copy of the current report for inspection
// this should really be moved to an export_test.go or rethought
// to not be needed.
func (r Report) Levels() []int {
	return slices.Clone(r.levels)
}

// Safe says if the report is safe or not
//
// safe report
// - levels are all increasing
// - levels differ, but by no more than 3
func (r Report) Safe() bool {
	if len(r.levels) < 3 { // what about 0?... skip or safe?
		return true
	}

	// need a change on all. first change defines direction
	if r.levels[0] == r.levels[1] {
		return false
	}

	increasing := r.levels[1] > r.levels[0]
	for idx, level := range r.levels[1:] {
		change := delta(level, r.levels[idx], increasing)
		if change > 3 || change < 1 {
			return false
		}
	}
	return true
}

// SafeDampened allows the removal of one number to create a safe report
func (r Report) SafeDampened() bool {
	if len(r.levels) < 3 { // what about 0?... skip or safe?
		fmt.Println("there were 0 levels. look into this") // TODO: remove
		return true
	}
	safe := true
	// TODO: move idx determination into PossibleFixes
	increasing := r.levels[1] > r.levels[0]
	for idx, level := range r.levels {
		if idx == 0 {
			continue
		}
		change := delta(level, r.levels[idx-1], increasing)
		if change > 3 || change < 1 {
			safe = false
			// if this idx is bad, try some alternatives
			for alt := range r.PossibleFixes(idx) {
				if alt.Safe() {
					return true
				}
			}
			// no need to check the rest of the levels. did not find a safe alternative
			// for the issue at idx
			break
		}
	}
	return safe
}

// possibleFixes recieves the index where the first non-conforming value is for the
// repot, and returns possible deletions
func (r Report) PossibleFixes(problemIdx int) iter.Seq[Report] {
	firstRemoved := false
	return func(yield func(r Report) bool) {
		// problemIdx is never the first, but we should try removing the first
		// since it could change the direction for increase/decreasing
		// 3 2 4 5 6 7
		if !firstRemoved {
			original := slices.Clone(r.levels)
			newLevels := slices.Delete(original, 0, 1)
			if !yield(Report{levels: newLevels}) {
				return
			}
			firstRemoved = true
		}

		// problemIdx is only 1, if the jump was too big. e.g.
		// 1 5 6 7 8
		// 1 5 4 3 2
		// 1 5 2 3 4
		// 10 5 6 7 8
		// 10 5 4 3 2
		// 10 5 11 12 13
		// For these cases, elminating the first or second might help
		if problemIdx == 1 || problemIdx == 2 {
			for i := 0; i < 2; i++ {
				original := slices.Clone(r.levels)
				// delete the first? use i
				newLevels := slices.Delete(original, problemIdx+i-1, problemIdx+i)
				// fmt.Println(newLevels)
				if !yield(Report{levels: newLevels}) {
					return
				}
			}
		}

		// problemIdx is the 3rd or beyond
		// 1 2 3 9 4 5 // delete the problemIdx
		// 1 2 3 3 4 5 // delete the problemIdx or the problemIdx - 1
		// 1 3 5 7 9 8 // delete the problemIdx
		// 1 3 5 4 5 7 // must delete the problemIdx - 1
		if problemIdx > 2 {
			// offer the deltion othe the problemidx or th eone before it
			for i := 0; i < 2; i++ { // 0, 1
				original := slices.Clone(r.levels)
				newLevels := slices.Delete(original, problemIdx+i-1, problemIdx+i)
				// fmt.Println(newLevels)
				if !yield(Report{newLevels}) {
					return
				}
			}
		}
	}
}

// delta calculates the change in the given direction.
func delta(snd, fst int, increasing bool) int {
	if increasing {
		return snd - fst
	}
	return fst - snd
}

func loadReports(in io.Reader) ([]Report, error) {
	out := []Report{}
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		rep, err := ParseReport(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("failed to parse report: %w", err)
		}
		out = append(out, rep)

	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read reports: %w", err)
	}
	return out, nil
}

func ParseReport(raw string) (Report, error) {
	fields := strings.Fields(raw)
	levels := make([]int, len(fields))
	var err error
	for i, field := range fields {
		levels[i], err = strconv.Atoi(field)
		if err != nil {
			return Report{}, fmt.Errorf("failed to read reports: %w", err)
		}
	}

	return Report{levels}, nil
}

// refactor out the reading and converting...
// since I don't want to hold things in memory, just opt
// for scanning normally, and then a function that converts
// each line to a report
func readReports(in io.Reader) iter.Seq[Report] {
	scanner := bufio.NewScanner(in)
	return func(yield func(r Report) bool) {
		for scanner.Scan() {
			if err := scanner.Err(); err != nil {
				// don't panic... but how to handle an error
				// in an iterator...
				panic(err)
			}

			fields := strings.Fields(scanner.Text())
			levels := make([]int, len(fields))
			var err error
			for i, field := range fields {
				levels[i], err = strconv.Atoi(field)
				if err != nil {
					panic(err)
				}
			}
			if !yield(Report{levels}) {
				return
			}
		}
	}
}
