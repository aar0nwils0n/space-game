package main

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
	"math/rand"
	"math"
)

type Asteroid struct {
	radius float64
	image string
	img *js.Object
	ctx *dom.CanvasRenderingContext2D
	x float64
	y float64
	exploded bool
}


func (a *Asteroid) Intersects(ship Ship) {
	xDistance := math.Abs(a.x - ship.x)	
	yDistance := math.Abs(a.y - ship.y)
	hypot := math.Hypot(xDistance, yDistance)
	distance := hypot - (a.radius + ship.radius) * 0.25
	
	if(distance < 0) {
		a.exploded = true;
	}
}

func (a *Asteroid) CreateRandom() {
	a.image = "./asteroid.png"
	a.img = js.Global.Get("Image").New()
	a.img.Set("src", a.image)
	a.radius = rand.Float64() * 200
	a.x = rand.Float64() * 800;
	a.y = rand.Float64() * 800;
}

func (a *Asteroid) Draw() {
	if(a.exploded != true) {
		a.ctx.Call("drawImage", a.img, a.x, a.y, a.radius, a.radius)
	}
}