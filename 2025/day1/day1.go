// Package day1 solves AoC 2025 day 1
package day1

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// Direction represents the rotation direction.
type Direction int

const (
	Left Direction = iota
	Right
)

func (d Direction) String() string {
	return map[Direction]string{
		Left:  "Left",
		Right: "Right",
	}[d]
}

const (
	startingPosition = 50
	positionsTotal   = 100
)

// Move represents a rotation and distance for the safe.
type Move struct {
	Direction Direction
	Distance  int
}

// Dial represents the current position of the safe's dial.
type Dial struct {
	position int
}

// NewDial creates a new dial starting at position 50.
func NewDial() *Dial {
	return &Dial{position: startingPosition}
}

// Move updates the dial position based on the move.
// Right moves up (increases position), Left moves down (decreases position).
// The dial loops from 0 to 99.
func (d *Dial) Move(m Move) {
	if m.Direction == Right {
		d.position = (d.position + m.Distance) % positionsTotal
	} else {
		d.position = (d.position - m.Distance) % positionsTotal
		if d.position < 0 {
			d.position += positionsTotal
		}
	}
}

func (d *Dial) MoveV2(m Move) int {
	var zeros int
	if m.Direction == Right {
		d.position, zeros = MoveRight(d.position, m.Distance)
	} else {
		d.position, zeros = MoveLeft(d.position, m.Distance)
	}
	return zeros
}

func MoveRight(pos, num int) (int, int) {
	return MovePosition(pos, num), MoveZeros(pos, num)
}

func MoveLeft(pos, num int) (int, int) {
	return MovePosition(pos, -num), MoveZeros(pos, -num)
}

func MovePosition(pos, num int) int {
	finalPosition := (pos + num) % positionsTotal
	if finalPosition < 0 {
		finalPosition += positionsTotal
	}
	return finalPosition
}

func MoveZeros(pos, num int) int {
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
		dial.Move(move)
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
		fmt.Printf("Move: %v position: %v count: %v \n", move, dial.Position(), count)
	}

	return count, nil
}

// ParseMoves reads the input and returns a slice of moves.
func ParseMoves(r io.Reader) ([]Move, error) {
	var moves []Move
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			continue
		}

		var direction Direction
		switch line[0] {
		case 'L':
			direction = Left
		case 'R':
			direction = Right
		default:
			return nil, fmt.Errorf("invalid direction: %c", line[0])
		}

		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, fmt.Errorf("invalid distance in line %q: %w", line, err)
		}

		moves = append(moves, Move{
			Direction: direction,
			Distance:  distance,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return moves, nil
}
