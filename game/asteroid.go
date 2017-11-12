package game

import (
	"github.com/gopherjs/gopherjs/js"
)

//Asteroid creates staic image on canvas
type Asteroid struct {
	Exploder
	image string
	img   *js.Object
}

func (a *Asteroid) intersects(ship *Ship) bool {
	intersects := intersects(a.x, a.y, a.radius, ship.x, ship.y, ship.radius)

	if intersects {
		if a.explodeFrame == 0 {
			a.startExplosion()
		}
		if ship.explodeFrame == 0 {
			ship.startExplosion()
		}
	}

	return intersects
}

//Draw an asteroid on the canvas and explode if needed
func (a *Asteroid) Draw() {
	if a.explodeFrame == 0 {
		x := Round(a.x - a.radius)
		y := Round(a.y - a.radius)
		diameter := a.radius * 2
		a.Canvas.Ctx.Call("drawImage", a.img, x, y, diameter, diameter)
	} else if a.exploded() == false {
		a.explode()
	}
}
