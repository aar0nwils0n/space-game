package main

import (
	"github.com/gopherjs/gopherjs/js"
	"math/rand"
	"math"
)

type Asteroid struct {
	Exploder
	image string
	img *js.Object
}


func (a *Asteroid) Intersects(ship *Ship) {
	xDistance := math.Abs(a.x - ship.x)	
	yDistance := math.Abs(a.y - ship.y)
	hypot := math.Hypot(xDistance, yDistance)
	distance := hypot - (a.radius + ship.radius)
	
	if(distance < 0) {
		if(a.explodeFrame == 0) {
			a.explodeFrame = 1;
		}
		if(ship.explodeFrame == 0) {
			ship.explodeFrame = 1;
		}
	}
}

func (a *Asteroid) CreateRandom() {
	a.image = "./asteroid.png"
	a.img = js.Global.Get("Image").New()
	a.img.Set("src", a.image)
	a.radius = (rand.Float64() + 0.25) * 50
	a.x = rand.Float64() * 800;
	a.y = rand.Float64() * 800;
}

func (a *Asteroid) Draw() {
	if(a.explodeFrame == 0) {
		a.canvas.ctx.Call("drawImage", a.img, a.x, a.y, a.radius * 2, a.radius * 2)
	} else if(a.exploded() == false) {
		a.explode()
	} 
}