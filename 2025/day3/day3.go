// Package day3 solves AoC 2025 day 3
package day3

import (
	"io"
	"strconv"
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
	answer := 0
	_, err := ParseIn(r)
	if err != nil {
		return 0, err
	}
	// TODO: solve part 2

	return answer, nil
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
