// Package day3 solves AoC 2025 day 3
package day3

import (
	"io"
	"strconv"
	"strings"
)

// Bank is a bank of batteries. One line of the input.
type Bank []int

func Part1(r io.Reader) (int, error) {
	banks, err := ParseIn(r)
	if err != nil {
		return 0, err
	}

	total := 0
	for _, bank := range banks {
		total += Biggest(bank)
	}

	return total, nil
}

func Part2(r io.Reader) (int, error) {
	banks, err := ParseIn(r)
	if err != nil {
		return 0, err
	}
	total := 0
	for _, bank := range banks {
		total += Biggest12(bank)
	}

	return total, nil
}

func Biggest(r Bank) int {
	// Algorithm
	leftDigit := r[len(r)-2]
	rightDigit := r[len(r)-1]
	oldLeft := 0
	// scan each row backward, one digit at a time.
	for i := range r {
		if i == 0 || i == 1 {
			// start on the third digit from the right
			continue
		}
		idx := len(r) - 1 - i
		if leftDigit <= r[idx] {
			oldLeft = leftDigit
			leftDigit = r[idx]
		}

		if oldLeft > rightDigit {
			rightDigit = oldLeft
		}
	}

	out, err := strconv.Atoi(strconv.Itoa(leftDigit) + strconv.Itoa(rightDigit))
	if err != nil {
		panic(err)
	}
	return out
}

func Biggest12(r Bank) int {
	// initialize the answer with the furthers right 12 numbers
	const onBatteries = 12
	// answer := make([]int, onBatteries)
	answer := r[len(r)-onBatteries:]
	// iterate through each new left digit, returning a new answer
	// after considering each
	for idx := range r {
		if idx < onBatteries {
			continue
		}

		answer = Bigger(r[len(r)-idx-1], answer)
	}
	return convert(answer)
}

func convert(row []int) int {
	var strvalBuilder strings.Builder
	for _, val := range row {
		strvalBuilder.WriteString(strconv.Itoa(val))
	}
	strval := strvalBuilder.String()
	out, err := strconv.Atoi(strval)
	if err != nil {
		panic(err)
	}
	return out
}

func Bigger(num int, current []int) []int {
	if len(current) == 0 {
		// if current has zero length, return the original
		return current
	}
	if num < current[0] {
		// if num is smaller than the left most digit. return the original
		return current
	}

	// if it's bigger or equal
	//
	// replace the first digit with that one, and pass the old first digit to the right
	// asking if it makes the number to the right bigger
	oldFirst := current[0]
	current[0] = num
	remaining := Bigger(oldFirst, current[1:])

	// use the remaining result for the rest of the current
	for idx := range remaining {
		current[idx+1] = remaining[idx]
	}
	out := make([]int, len(current))
	copy(out, current)
	return out
}
