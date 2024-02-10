package adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSquarePegAdapter(t *testing.T) {
	hole := NewRoundHole(5)
	peg := NewRoundPeg(5)
	assert.True(t, hole.Fits(peg))

	smallSqpeg := NewSquarePeg(5)
	largeSqpeg := NewSquarePeg(10)
	sAdapter := NewSquarePegAdapter(smallSqpeg)
	lAdapter := NewSquarePegAdapter(largeSqpeg)
	assert.True(t, hole.Fits(sAdapter))
	assert.False(t, hole.Fits(lAdapter))
}
