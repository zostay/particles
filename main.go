package main

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Particles",
		Bounds: pixel.R(0, 0, 2800, 1400),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())

	pss := make([]*Particles, 4)
	for i := range pss {
		ps := make(Particles, 200)
		ps.Initialize(pixel.RGB(rand.Float64(), rand.Float64(), rand.Float64()))

		cx := rand.Float64()*win.Bounds().W() + win.Bounds().Min.X
		cy := rand.Float64()*win.Bounds().H() + win.Bounds().Min.Y
		size := rand.Float64()*5 + 10
		ps.Boom(pixel.V(cx, cy), size)

		pss[i] = &ps
	}

	// ps := Particles{
	// 	{
	// 		Loc: win.Bounds().Center(),
	// 		Rad: 1,
	// 		Vel: pixel.V(12, 0),
	// 		Col: pixel.RGB(1, 1, 1),
	// 	},
	// }

	e := Engine{
		Gravity:    pixel.V(0, -0.1),
		Edges:      win.Bounds(),
		Bounciness: pixel.R(1, 0.9, 1, 1),
		Drag:       0.997,
		Wind:       0,
	}

	imd := imdraw.New(nil)
	for !win.Closed() {
		win.Clear(pixel.RGB(0, 0, 0))
		imd.Clear()

		for _, ps := range pss {
			ps.Draw(imd)
			e.Apply(*ps)
		}
		imd.Draw(win)

		win.Update()

		// fmt.Printf("tick vel.y=%f\n", ps[0].Vel.Y)
	}
}
