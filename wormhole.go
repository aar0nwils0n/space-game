package main

import (
	"math"
	"math/rand"
	"github.com/gopherjs/gopherjs/js"
)

type Wormhole struct {
	canvas *Canvas
	x float64
	y float64
	radius float64
	rotation float64
	img *js.Object
}

func (w *Wormhole) intersects(s *Ship) bool {
	xDistance := math.Abs(w.x - s.x)	
	yDistance := math.Abs(w.y - s.y)
	hypot := math.Hypot(xDistance, yDistance)
	distance := hypot - (w.radius + s.radius)
	return distance < 0;
}

func (w *Wormhole) init() {
	w.radius = 100
	w.x = rand.Float64() * 800
	w.y = rand.Float64() * 800
	w.img = js.Global.Get("Image").New()
	w.img.Set("src", "./wormhole.png")
}

func (w *Wormhole) draw() {
	w.canvas.ctx.Save(); // save current state
	w.canvas.ctx.Translate(int(w.x), int(w.y));
	w.canvas.ctx.Rotate(w.rotation); // rotate
	w.canvas.ctx.Call("drawImage", w.img, -w.radius, -w.radius,  w.radius * 2, w.radius * 2)
	w.canvas.ctx.Restore();
	w.rotation += 0.1
}