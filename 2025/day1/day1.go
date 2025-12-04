// Package day1 solves AoC 2025 day 1
package day1

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

const (
	startingPosition = 50
	positionsTotal   = 100
)

// Move represents a rotation and distance for the safe.
type Move struct {
	Distance int // negative means left
}

// Dial represents the current position of the safe's dial.
type Dial struct {
	position int
}

// NewDial creates a new dial starting at position 50.
func NewDial() *Dial {
	return &Dial{position: startingPosition}
}

func (d *Dial) MoveV2(m Move) int {
	var zeros int
	d.position, zeros = MoveDial(d.position, m.Distance)
	return zeros
}

func MoveDial(pos, num int) (int, int) {
	return movePosition(pos, num), moveZeros(pos, num)
}

func movePosition(pos, num int) int {
	finalPosition := (pos + num) % positionsTotal
	if finalPosition < 0 {
		finalPosition += positionsTotal
	}
	return finalPosition
}

func moveZeros(pos, num int) int {
	zeros := 0
	virtualPosition := pos + num
	if virtualPosition <= 0 && pos != 0 {
		// didn't start at zero, but did reach zero on first move left
		zeros++
	}
	zeros += abs(virtualPosition / positionsTotal) // added rotations
	return zeros
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Position returns the current position of the dial.
func (d *Dial) Position() int {
	return d.position
}

func Part1(r io.Reader) (int, error) {
	moves, err := ParseMoves(r)
	if err != nil {
		return 0, err
	}

	dial := NewDial()
	count := 0

	for _, move := range moves {
		_ = dial.MoveV2(move)
		if dial.Position() == 0 {
			count++
		}
	}

	return count, nil
}

func Part2(r io.Reader) (int, error) {
	moves, err := ParseMoves(r)
	if err != nil {
		return 0, err
	}

	dial := NewDial()
	count := 0

	for _, move := range moves {
		count += dial.MoveV2(move)
		// fmt.Printf("Move: %v position: %v count: %v \n", move, dial.Position(), count)
	}

	return count, nil
}

// ParseMoves reads the input and returns a slice of moves.
func ParseMoves(r io.Reader) ([]Move, error) {
	var moves []Move
	scanner := bufio.NewScanner(r)

	const minLineLen = 2
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < minLineLen {
			continue
		}

		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, fmt.Errorf("invalid distance in line %q: %w", line, err)
		}

		switch line[0] {
		case 'L':
			distance = -distance
		case 'R':
		default:
			return nil, fmt.Errorf("invalid direction: %c", line[0])
		}

		moves = append(moves, Move{
			Distance: distance,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error while scanning input: %w", err)
	}

	return moves, nil
}
