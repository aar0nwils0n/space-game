package game

import (
	"math"

	"github.com/gopherjs/gopherjs/js"
)

//Wormhole is a rotating image that is drawn on the canvas
type Wormhole struct {
	canvas   *Canvas
	x        float64
	y        float64
	radius   float64
	rotation float64
	img      *js.Object
}

func (w *Wormhole) intersects(s *Ship) bool {
	xDistance := math.Abs(w.x - s.x)
	yDistance := math.Abs(w.y - s.y)
	hypot := math.Hypot(xDistance, yDistance)
	distance := hypot - (w.radius + s.radius)
	return distance < 0
}

func (w *Wormhole) init() {
	w.radius = 100
	w.x = 750
	w.y = 750
	w.img = js.Global.Get("Image").New()
	w.img.Set("src", "./assets/images/wormhole.png")
}

func (w *Wormhole) draw() {
	w.canvas.Ctx.Save() // save current state
	w.canvas.Ctx.Translate(int(w.x), int(w.y))
	w.canvas.Ctx.Rotate(w.rotation) // rotate
	w.canvas.Ctx.Call("drawImage", w.img, -w.radius*1.5, -w.radius*1.5, w.radius*2*1.5, w.radius*2*1.5)
	w.canvas.Ctx.Restore()
	w.rotation += 0.1
}
