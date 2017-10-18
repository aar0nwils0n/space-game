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

func (a *Asteroid) intersects(s *Ship) bool {
	intersecting := intersects(a.x, a.y, a.radius, s.x, s.y, s.radius)

	if intersecting {
		if a.explodeFrame == 0 {
			a.explodeFrame = 1
		}
		if s.explodeFrame == 0 {
			s.explodeFrame = 1
		}
	}

	return intersecting
}

//Draws asteroid on canvas and progresses explosion if nessecary
func (a *Asteroid) Draw() {
	if a.explodeFrame == 0 {
		a.Canvas.Ctx.Call("drawImage", a.img, a.x-a.radius, a.y-a.radius, a.radius*2, a.radius*2)
	} else if a.exploded() == false {
		a.explode()
	}
}
