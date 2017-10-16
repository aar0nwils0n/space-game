package game

import (
	"math"

	"github.com/gopherjs/gopherjs/js"
)

//Asteroid creates staic image on canvas
type Asteroid struct {
	Exploder
	image string
	img   *js.Object
}

func (a *Asteroid) intersects(ship *Ship) bool {
	xDistance := math.Abs(a.x - ship.x)
	yDistance := math.Abs(a.y - ship.y)
	hypot := math.Hypot(xDistance, yDistance)
	distance := hypot - a.radius - ship.radius

	if distance < 0 {
		if a.explodeFrame == 0 {
			a.explodeFrame = 1
		}
		if ship.explodeFrame == 0 {
			ship.explodeFrame = 1
		}
	}

	return distance < 0
}

func (a *Asteroid) draw() {
	if a.explodeFrame == 0 {
		a.Canvas.Ctx.Call("drawImage", a.img, a.x-a.radius, a.y-a.radius, a.radius*2, a.radius*2)
	} else if a.exploded() == false {
		a.explode()
	}
}
