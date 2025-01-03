package day16_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jstensland/advent-of-code/2024/day16"
)

func TestRight(t *testing.T) {
	assert.Equal(t, day16.Position{0, 0, day16.East}, day16.Position{0, 0, day16.North}.Right())
	assert.Equal(t, day16.Position{0, 0, day16.South}, day16.Position{0, 0, day16.East}.Right())
	assert.Equal(t, day16.Position{0, 0, day16.West}, day16.Position{0, 0, day16.South}.Right())
	assert.Equal(t, day16.Position{0, 0, day16.North}, day16.Position{0, 0, day16.West}.Right())
}

func TestLeft(t *testing.T) {
	assert.Equal(t, day16.Position{0, 0, day16.West}, day16.Position{0, 0, day16.North}.Left())
	assert.Equal(t, day16.Position{0, 0, day16.North}, day16.Position{0, 0, day16.East}.Left())
	assert.Equal(t, day16.Position{0, 0, day16.East}, day16.Position{0, 0, day16.South}.Left())
	assert.Equal(t, day16.Position{0, 0, day16.South}, day16.Position{0, 0, day16.West}.Left())
}

func TestRightForward(t *testing.T) {
	assert.Equal(t, day16.Position{3, 4, day16.East}, day16.Position{3, 3, day16.North}.Right().Forward())
	assert.Equal(t, day16.Position{4, 3, day16.South}, day16.Position{3, 3, day16.East}.Right().Forward())
	assert.Equal(t, day16.Position{3, 2, day16.West}, day16.Position{3, 3, day16.South}.Right().Forward())
	assert.Equal(t, day16.Position{2, 3, day16.North}, day16.Position{3, 3, day16.West}.Right().Forward())
}
