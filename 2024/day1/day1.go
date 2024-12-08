// Package day1 solves the first day of the advent of code
package day1

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"slices"
	"strconv"
	"strings"
)

// SolvePart1 sorts both lists and calculate the distance between each
// corresponding pairs between the lists
func SolvePart1(in io.Reader) (int, error) {
	left, right, err := loadInput(in)
	if err != nil {
		log.Fatalf("Error reading input: %s", err)
	}
	left.sort()
	right.sort()

	return left.distance(right), nil
}

// SolvePart2 loads the lists and calculates the frequency distance between the lists
func SolvePart2(in io.Reader) (int, error) {
	left, right, err := loadInput(in)
	if err != nil {
		log.Fatalf("Error reading input: %s", err)
	}

	return left.freqDistance(right), nil
}

type list struct {
	list    []int
	freqMap map[int]int
}

func newList() *list {
	return &list{
		freqMap: make(map[int]int),
	}
}

func (l list) sort() {
	slices.Sort(l.list)
}

func (l *list) insert(val int) {
	l.list = append(l.list, val)
	if _, ok := l.freqMap[val]; !ok {
		l.freqMap[val] = 0
	}
	l.freqMap[val]++
}

func (l list) distance(other *list) int {
	distance := 0
	for idx, val := range l.list {
		distance += abs(val - other.list[idx])
	}
	return distance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (l list) freqDistance(other *list) int {
	total := 0
	for _, val := range l.list {
		lFreq, ok := l.freqMap[val]
		if !ok {
			panic("if it's in the list, it should appear at least once")
		}
		otherFreq, ok := other.freqMap[val]
		if !ok {
			otherFreq = 0
		}

		total += val * lFreq * otherFreq
	}
	return total
}

func loadInput(in io.Reader) (*list, *list, error) {
	left := newList()
	right := newList()
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		leftVal, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, nil, fmt.Errorf("value was not an int: %w", err)
		}
		left.insert(leftVal)

		rightVal, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, nil, fmt.Errorf("value was not an int: %w", err)
		}
		right.insert(rightVal)
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error loading input: %w", err)
	}
	return left, right, nil
}
