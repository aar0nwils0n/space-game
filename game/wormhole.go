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

func createWormhole(c *Canvas, x float64, y float64, rad float64, img string) *Wormhole {
	wormhole := Wormhole{}
	wormhole.canvas = c
	wormhole.radius = rad * wormhole.canvas.vh
	wormhole.x = x*wormhole.canvas.vw - wormhole.radius
	wormhole.y = y*wormhole.canvas.vh - wormhole.radius
	wormhole.img = js.Global.Get("Image").New()
	wormhole.img.Set("src", img)
	return &wormhole
}

func (w *Wormhole) intersects(s *Ship) bool {
	return intersects(w.x, w.y, w.radius, s.x, s.y, s.radius)
}

//Draw and rotate a wormhole on the canvas
func (w *Wormhole) Draw() {
	w.canvas.Ctx.Save() // save current state
	w.canvas.Ctx.Translate(int(w.x), int(w.y))
	w.canvas.Ctx.Rotate(w.rotation) // rotate
	oneAndHalfRad := Round(w.radius * 1.5)
	threeRad := Round(w.radius * 3)
	w.canvas.Ctx.Call("drawImage", w.img, -oneAndHalfRad, -oneAndHalfRad, threeRad, threeRad)
	w.canvas.Ctx.Restore()
	w.rotation += 0.1
}
