package main

import (
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Particle struct {
	Loc pixel.Vec
	Vel pixel.Vec
	Col pixel.RGBA
	Rad float64
}

type Particles []*Particle

func (ps Particles) Initialize(col pixel.RGBA) {
	for i := range ps {
		ps[i] = &Particle{
			Rad: 2,
			Col: col,
		}
	}
}

func (ps Particles) Draw(imd *imdraw.IMDraw) {
	for _, p := range ps {
		imd.Color = p.Col
		imd.Push(p.Loc)
		imd.Circle(p.Rad, 0)
	}
}

func (ps Particles) Boom(loc pixel.Vec, scale float64) {
	for _, p := range ps {
		p.Loc = loc
		s := rand.Float64() * scale
		p.Vel = pixel.V(rand.Float64()-0.5, rand.Float64()-0.5).Unit().Scaled(s)
	}
}
