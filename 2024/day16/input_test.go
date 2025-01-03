package day16_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jstensland/advent-of-code/2024/day16"
)

func TestRight(t *testing.T) {
	zeroLoc := day16.Location{0, 0}
	assert.Equal(t, day16.Position{zeroLoc, day16.East}, day16.Position{zeroLoc, day16.North}.Right())
	assert.Equal(t, day16.Position{zeroLoc, day16.South}, day16.Position{zeroLoc, day16.East}.Right())
	assert.Equal(t, day16.Position{zeroLoc, day16.West}, day16.Position{zeroLoc, day16.South}.Right())
	assert.Equal(t, day16.Position{zeroLoc, day16.North}, day16.Position{zeroLoc, day16.West}.Right())
}

func TestLeft(t *testing.T) {
	zeroLoc := day16.Location{0, 0}
	assert.Equal(t, day16.Position{zeroLoc, day16.West}, day16.Position{zeroLoc, day16.North}.Left())
	assert.Equal(t, day16.Position{zeroLoc, day16.North}, day16.Position{zeroLoc, day16.East}.Left())
	assert.Equal(t, day16.Position{zeroLoc, day16.East}, day16.Position{zeroLoc, day16.South}.Left())
	assert.Equal(t, day16.Position{zeroLoc, day16.South}, day16.Position{zeroLoc, day16.West}.Left())
}

func TestRightForward(t *testing.T) {
	startLoc := day16.Location{3, 3}
	assert.Equal(t,
		day16.Position{day16.Location{3, 4}, day16.East},
		day16.Position{startLoc, day16.North}.Right().Forward(),
	)
	assert.Equal(t,
		day16.Position{day16.Location{4, 3}, day16.South},
		day16.Position{startLoc, day16.East}.Right().Forward(),
	)
	assert.Equal(t,
		day16.Position{day16.Location{3, 2}, day16.West},
		day16.Position{startLoc, day16.South}.Right().Forward(),
	)
	assert.Equal(t,
		day16.Position{day16.Location{2, 3}, day16.North},
		day16.Position{startLoc, day16.West}.Right().Forward(),
	)
}
