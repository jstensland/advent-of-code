// Package day13 solves AoC 2024 day 13
package day13

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"regexp"
	"strconv"
)

const (
	ACost = 3
	BCost = 1
)

var ErrUnsolvable = errors.New("unsolvable game")

// SolvePart1 finds the right combo of buttons to reach the prize.
func SolvePart1(in io.Reader) (int, error) {
	games, err := ParseIn(in)
	if err != nil {
		return 0, fmt.Errorf("error loading input: %w", err)
	}

	total := 0 // total tokens needed for all  wins
	for _, game := range games {
		cost, err := game.Solve()
		if errors.Is(err, ErrUnsolvable) {
			continue
		}
		total += cost
	}
	return total, nil
}

func SolvePart2(in io.Reader) (int, error) {
	games, err := ParseIn(in)
	if err != nil {
		return 0, fmt.Errorf("error loading input: %w", err)
	}

	total := 0 // total tokens needed for all  wins
	for _, game := range games {
		game.Correct()
		cost, err := game.Solve()
		if errors.Is(err, ErrUnsolvable) {
			continue
		}
		total += cost
	}
	return total, nil
}

type Coordinate struct {
	X int
	Y int
}

type Button struct {
	Cost   int
	Count  int
	XDelta int
	YDelta int
}

func (b *Button) Slope() float64 {
	return float64(b.YDelta) / float64(b.XDelta)
}

func (b *Button) Price() int {
	return b.Cost * b.Count
}

// UnitPrice calculates the cost of pressing the button per unit travelled.
func (b *Button) UnitPrice() float64 {
	return float64(b.Cost) / math.Sqrt(float64(b.XDelta*b.XDelta+b.YDelta*b.YDelta))
}

type Game struct {
	A       Button
	B       Button
	Prize   Coordinate
	Current Coordinate
}

func (g *Game) Correct() {
	// part2 says the prizes should be 10000000000000 further out
	g.Prize.X += 10000000000000
	g.Prize.Y += 10000000000000
}

// Solve is a revamp applicable for small and large numbers
func (g *Game) Solve() (int, error) {
	slopeToPrize := float64(g.Prize.Y) / float64(g.Prize.X)

	// - check if the prize is too high or too low to be reached, drop case
	if slopeToPrize < g.B.Slope() && slopeToPrize < g.A.Slope() ||
		slopeToPrize > g.B.Slope() && slopeToPrize > g.A.Slope() {
		return 0, ErrUnsolvable
	}

	// Solve for g.A.Count and g.B.Count in the following system
	//
	// g.A.Count*g.A.XDelta + g.B.Count*g.B.XDelta = g.Prize.X
	// g.A.Count*g.A.YDelta + g.B.Count*g.B.YDelta = g.Prize.Y

	g.A.Count = (g.B.XDelta*g.Prize.Y - g.B.YDelta*g.Prize.X) /
		(g.B.XDelta*g.A.YDelta - g.B.YDelta*g.A.XDelta)

	if (g.B.XDelta*g.Prize.Y-g.B.YDelta*g.Prize.X)%
		(g.B.XDelta*g.A.YDelta-g.B.YDelta*g.A.XDelta) != 0 {
		// cannot be done if there is some left over
		return 0, ErrUnsolvable
	}

	g.B.Count = (g.A.XDelta*g.Prize.Y - g.A.YDelta*g.Prize.X) /
		(g.A.XDelta*g.B.YDelta - g.B.XDelta*g.A.YDelta)

	if (g.A.XDelta*g.Prize.Y-g.A.YDelta*g.Prize.X)%
		(g.A.XDelta*g.B.YDelta-g.B.XDelta*g.A.YDelta) != 0 {
		// cannot be done if there is some left over
		return 0, ErrUnsolvable
	}

	return g.A.Price() + g.B.Price(), nil
}

// ParseIn reads in all the games. For solving.
func ParseIn(in io.Reader) ([]Game, error) {
	scanner := bufio.NewScanner(in)
	scanner.Split(splitOnDoubleCR)
	games := []Game{}
	// may want to read them all and split on double new lines, or just read until newline
	// parsing as you go
	buttonAMatch := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
	buttonBMatch := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
	prizeMatch := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	for scanner.Scan() {
		// read the first line as Button A
		puzzle := scanner.Text()
		aM := buttonAMatch.FindStringSubmatch(puzzle)
		xADelta, err := strconv.Atoi(aM[1])
		if err != nil {
			return nil, fmt.Errorf("error parsing xDelta: %w", err)
		}
		yADelta, err := strconv.Atoi(aM[2])
		if err != nil {
			return nil, fmt.Errorf("error parsing yDelta: %w", err)
		}

		bM := buttonBMatch.FindStringSubmatch(puzzle)
		xBDelta, err := strconv.Atoi(bM[1])
		if err != nil {
			return nil, fmt.Errorf("error parsing xDelta: %w", err)
		}
		yBDelta, err := strconv.Atoi(bM[2])
		if err != nil {
			return nil, fmt.Errorf("error parsing yDelta: %w", err)
		}

		prizeM := prizeMatch.FindStringSubmatch(puzzle)
		prizeX, err := strconv.Atoi(prizeM[1])
		if err != nil {
			return nil, fmt.Errorf("error parsing xDelta: %w", err)
		}
		prizeY, err := strconv.Atoi(prizeM[2])
		if err != nil {
			return nil, fmt.Errorf("error parsing yDelta: %w", err)
		}

		game := Game{
			A:     Button{Cost: ACost, XDelta: xADelta, YDelta: yADelta},
			B:     Button{Cost: BCost, XDelta: xBDelta, YDelta: yBDelta},
			Prize: Coordinate{X: prizeX, Y: prizeY},
		}
		games = append(games, game)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return games, nil
}

// implements Splitfunc for the scanner. https://pkg.go.dev/bufio#SplitFunc
func splitOnDoubleCR(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil // should this return ErrFinalToken?
	}

	// Find the next double newline
	splitToken := []byte("\n\n")
	if i := bytes.Index(data, splitToken); i >= 0 {
		// We found a double carriage return.
		// Return the token up to that point.
		return i + len(splitToken), data[:i], nil
	}

	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}

	// Request more data.
	return 0, nil, nil
}

// SolveSlow solves the problems by taking individual steps. It was a first iteration and
// is not used
//
//nolint:cyclop // improved version is "Solve"
func (g *Game) SolveSlow() (int, error) {
	slopeToPrize := float64(g.Prize.Y) / float64(g.Prize.X)

	// - check if the prize is too high or too low to be reached, drop case
	if slopeToPrize < g.B.Slope() && slopeToPrize < g.A.Slope() ||
		slopeToPrize > g.B.Slope() && slopeToPrize > g.A.Slope() {
		return 0, ErrUnsolvable
	}

	// Determine which button is cheaper for distance (cost/hypotenuse)
	var cheaperButton *Button
	var pricierButton *Button
	if g.A.UnitPrice() < g.B.UnitPrice() {
		cheaperButton = &g.A
		pricierButton = &g.B
	} else {
		cheaperButton = &g.B
		pricierButton = &g.A
	}

	for !g.CheckFinalRun(cheaperButton) && g.Playing() {
		g.Press(pricierButton)
	}

	// use cheaper button only in the end
	for g.Playing() {
		g.Press(cheaperButton)
	}

	// check we solved it
	if g.Current.X != g.Prize.X || g.Current.Y != g.Prize.Y {
		return 0, ErrUnsolvable
	}

	return g.A.Price() + g.B.Price(), nil
}

func (g *Game) Playing() bool {
	if g.Current.X == g.Prize.X && g.Current.Y == g.Prize.Y {
		return false // succeeded
	}
	if g.Current.X > g.Prize.X || g.Current.Y > g.Prize.Y {
		return false // failed
	}
	return true
}

func (g *Game) Press(b *Button) {
	g.Current.X += b.XDelta
	g.Current.Y += b.YDelta
	b.Count++
}

// CheckFinalRun checks if you can make it to the end with this button alone now
func (g *Game) CheckFinalRun(b *Button) bool {
	xToGo := g.Prize.X - g.Current.X
	yToGo := g.Prize.Y - g.Current.Y

	if xToGo/b.XDelta == yToGo/b.YDelta &&
		xToGo%b.XDelta == 0 && yToGo%b.YDelta == 0 {
		return true
	}
	return false
}
