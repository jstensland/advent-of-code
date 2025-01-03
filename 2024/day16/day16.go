package day16

import (
	"fmt"
	"io"
	"slices"
	"sort"
)

const (
	turnCost = 1000
	moveCost = 1
)

func SolvePart1(in io.Reader) (int, error) {
	grid, err := ParseIn(in)
	if err != nil {
		return 0, fmt.Errorf("error loading input: %w", err)
	}

	return grid.BestRoute(), nil
}

// BestRoute looks for the best route from the start to the end and returns
// its "score." A move forward costs 1, and a turn costs 1000
//
// The general algorithm is to take a step and searh all paths from there, start
// with the cheapest. Once a successful path is found, subsequent searches will give
// up after they have a total cost equal to the currently best path
//
// Optimizations... if you get to a position that's been checked, stop.
// - if cost is less, calculate the new best
// - if the cost is worse, ignore
//
// May want to track successful paths for debugging purposes...
func (g *Grid) BestRoute() int {
	// to start, we need to check all 4 directions, so check one ahead
	backward := g.search(g.Start.Right().Right().Forward(), 0)
	minCost := g.search(g.Start, 0)
	if backward > 0 && backward < minCost {
		return backward
	}
	return minCost
}

func SolvePart2(in io.Reader) (int, error) {
	grid, err := ParseIn(in)
	if err != nil {
		return 0, fmt.Errorf("error loading input: %w", err)
	}

	return grid.BestSeats(), nil
}

func (g *Grid) BestSeats() int {
	// to start, we need to check all 4 directions, so check one ahead
	backward := g.search(g.Start.Right().Right().Forward(), 0)
	minCost := g.search(g.Start, 0)
	if backward > 0 && backward < minCost {
		minCost = backward
	}

	// now that we know the min cost, do the same thing, but record all the
	// locations on the min cost route...
	// This definitely relies on internal state of the grid, which isn't ideal, but is
	// quite efficient the second time because the only the leastCost route continues
	// due to accumlated state from the previous run
	g.minAnswer = minCost
	g.search(g.Start.Right().Right().Forward(), 0)
	g.search(g.Start, 0)
	return len(g.counted)
}

// search will recursively search the next steps for each node and return the cost
// of the current position plus getting to the next place
func (g *Grid) search(pos Position, cost int) int {
	// if it's a wall, this is an impossible move
	if g.GetLoc(pos) == Wall {
		return -1
	}
	// if you've been here before for less, skip it
	if lastCost, ok := g.visited[pos]; ok && lastCost < cost {
		return -1 // got here for less some other way
	}
	g.visited[pos] = cost

	// check if the cost is already beyond the known cheapest. Quit if so
	if g.leastCost > 0 && cost > g.leastCost {
		return -1
	}
	// check if you're at the ending location. Return accumulated cost if so.
	if pos == g.End {
		if cost < g.leastCost {
			g.leastCost = cost
		}

		// for the second run...
		if cost == g.minAnswer {
			g.counted[pos.Location] = true // count the ending spot
		}

		return cost
	}

	// search cost from forward, right and left and select the lowest
	costs := []int{
		g.search(pos.Forward(), cost+moveCost),
		g.search(pos.Right().Forward(), cost+moveCost+turnCost),
		g.search(pos.Left().Forward(), cost+moveCost+turnCost),
	}
	nextCost := leastCost(costs)
	if nextCost == g.minAnswer {
		g.counted[pos.Location] = true
	}
	return nextCost
}

// leastCost finds the samllest cost that's not -1. Returns -1 if no such cost
func leastCost(costs []int) int {
	tmp := make([]int, len(costs)) // sort a copy for no side effects
	copy(tmp, costs)
	sort.Ints(tmp)
	idx := slices.IndexFunc(tmp, func(n int) bool {
		return n > 0
	})
	if idx == -1 {
		return -1
	}
	return tmp[idx]
}
