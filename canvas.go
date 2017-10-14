package main

import (
	"honnef.co/go/js/dom"
	"github.com/gopherjs/gopherjs/js"
)

type Canvas struct {
	ctx *dom.CanvasRenderingContext2D
	ship Ship
	asteroids []*Asteroid
	explosion *js.Object
	width float64
	height float64
}

func (c *Canvas) CreateAsteroids(number float64) {
	for i := 1; i < int(number); i++ {
		asteroid := Asteroid{}
		asteroid.canvas = c
		asteroid.CreateRandom();
		c.asteroids = append(c.asteroids, &asteroid)
	}
}

func (c *Canvas) Initialize() {
	c.ship.Initialize()
	c.CreateAsteroids(10)
	c.explosion = js.Global.Get("Image").New()
	c.explosion.Set("src", "./explosion.png")
}

func (c *Canvas) Draw() {
	c.ctx.ClearRect(0, 0, int(c.width), int(c.height));
	c.ship.Draw();
	
	for _, a := range c.asteroids {
		a.Intersects(&c.ship)
		a.Draw()
	}
}