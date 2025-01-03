package day16

import (
	"fmt"
	"io"
	"slices"
	"sort"
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

// search will recursively search the next steps for each node
func (g *Grid) search(loc Position, cost int) int {
	// fmt.Println("searching", loc, "cost:", cost)
	// if it's a wall, this is an impossible move
	if g.GetLoc(loc) == Wall {
		return -1
	}
	// if you've been here before... skip it
	if lastCost, ok := g.visited[loc]; ok && lastCost < cost {
		return -1 // got here for less some other way
	}
	g.visited[loc] = cost

	// check if the cost is already beyond the known cheapest. Quit if so
	if g.leastCost > 0 && cost > g.leastCost {
		return -1
	}
	// check if you're at the ending location. Return accumulated cost if so.
	if loc == g.End {
		// fmt.Println("found end! cost:", cost)
		if cost < g.leastCost {
			g.leastCost = cost
		}
		return cost
	}

	// search cost from forward, right and left
	fwd := g.search(loc.Forward(), cost+1)
	right := g.search(loc.Right().Forward(), cost+1001)
	left := g.search(loc.Left().Forward(), cost+1001)

	// select the lowest that's not -1
	costs := []int{fwd, right, left}
	sort.Ints(costs)
	idx := slices.IndexFunc(costs, func(n int) bool {
		return n > 0
	})
	if idx == -1 {
		return -1
	}
	// fmt.Println("searching", loc, "incoming cost:", cost)
	// fmt.Println("cost is:", costs[idx])
	return costs[idx]
}
