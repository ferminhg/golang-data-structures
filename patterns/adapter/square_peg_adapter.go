package adapter

import math "math"

type RoundWithRadius interface {
	GetRadius() int
}

type RoundHole struct {
	radius int
}

func NewRoundHole(radius int) *RoundHole {
	return &RoundHole{radius: radius}
}

func (r *RoundHole) GetRadius() int {
	return r.radius
}

func (r *RoundHole) Fits(peg RoundWithRadius) bool {
	return r.radius >= peg.GetRadius()
}

type RoundPeg struct {
	radius int
}

func NewRoundPeg(radius int) *RoundPeg {
	return &RoundPeg{radius: radius}
}

func (p *RoundPeg) GetRadius() int {
	return p.radius
}

type SquarePeg struct {
	width int
}

func NewSquarePeg(width int) *SquarePeg {
	return &SquarePeg{width: width}
}

func (s *SquarePeg) GetWidth() int {
	return s.width
}

type SquarePegAdapter struct {
	peg *SquarePeg
}

func NewSquarePegAdapter(peg *SquarePeg) *SquarePegAdapter {
	return &SquarePegAdapter{peg: peg}
}

func (s *SquarePegAdapter) GetRadius() int {
	return int(float64(s.peg.GetWidth()) * math.Sqrt(2) / 2)
}
