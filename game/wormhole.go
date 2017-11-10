package game

import (
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

func createWormhole(c *Canvas) *Wormhole {
	wormhole := Wormhole{}
	wormhole.canvas = c
	wormhole.init()
	return &wormhole
}

func (w *Wormhole) intersects(s *Ship) bool {
	return intersects(w.x, w.y, w.radius, s.x, s.y, s.radius)
}

func (w *Wormhole) init() {
	w.radius = 10 * w.canvas.vh
	w.x = 100*w.canvas.vw - w.radius
	w.y = 100*w.canvas.vh - w.radius
	w.img = js.Global.Get("Image").New()
	w.img.Set("src", "./assets/images/wormhole.png")
}

//Draw and rotate a wormhole on the canvas
func (w *Wormhole) Draw() {
	w.canvas.Ctx.Save() // save current state
	w.canvas.Ctx.Translate(int(w.x), int(w.y))
	w.canvas.Ctx.Rotate(w.rotation) // rotate
	w.canvas.Ctx.Call("drawImage", w.img, -w.radius*1.5, -w.radius*1.5, w.radius*2*1.5, w.radius*2*1.5)
	w.canvas.Ctx.Restore()
	w.rotation += 0.1
}
