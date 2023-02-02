package main

import (
	"math"

	"github.com/faiface/pixel"
)

type Engine struct {
	Gravity    pixel.Vec
	Edges      pixel.Rect
	Bounciness pixel.Rect
	Drag       float64
	Wind       float64
}

func (e *Engine) Apply(ps Particles) {
	for _, p := range ps {
		// apply gravity
		p.Vel = p.Vel.Add(e.Gravity)

		// apply drag
		p.Vel = p.Vel.Scaled(e.Drag)

		// apply velocity to location
		p.Loc = p.Loc.Add(p.Vel)
		p.Loc.X += e.Wind

		// handle extreme velocity by forceably slowing things down
		if math.Abs(p.Vel.Y) > e.Edges.H() {
			p.Vel.Y /= 2
		}
		if math.Abs(p.Vel.X) > e.Edges.W() {
			p.Vel.X /= 2
		}

		// bottom bouncing
		scale := pixel.V(1, 1)
		if p.Loc.Y < e.Edges.Min.Y {
			p.Loc.Y = e.Edges.Min.Y - p.Loc.Y
			scale.Y = -e.Bounciness.Min.Y
		}

		// top bounce
		if p.Loc.Y > e.Edges.Max.Y {
			p.Loc.Y = p.Loc.Y - e.Edges.Min.Y
			scale.Y = -e.Bounciness.Min.Y
		}

		// left bounce
		if p.Loc.X < e.Edges.Min.X {
			p.Loc.X = e.Edges.Min.X - p.Loc.X
			scale.X = -e.Bounciness.Min.X
		}

		// right bounce
		if p.Loc.X > e.Edges.Max.X {
			p.Loc.X = p.Loc.X - e.Edges.Min.X
			scale.X = -e.Bounciness.Min.X
		}

		// apply bounce
		p.Vel = p.Vel.ScaledXY(scale)

		// And in case something else has gone wrong, force the point back into the box
		if !e.Edges.Contains(p.Loc) {
			if p.Loc.Y > e.Edges.Max.Y {
				p.Loc.Y = e.Edges.Max.Y
			}
			if p.Loc.Y < e.Edges.Min.Y {
				p.Loc.Y = e.Edges.Min.Y
			}
			if p.Loc.X > e.Edges.Max.X {
				p.Loc.X = e.Edges.Max.X
			}
			if p.Loc.X < e.Edges.Min.X {
				p.Loc.X = e.Edges.Min.X
			}
		}
	}
}
