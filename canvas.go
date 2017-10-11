package main

import (
	"honnef.co/go/js/dom"
)

type Canvas struct {
	ctx *dom.CanvasRenderingContext2D
	ship Ship
	asteroids []Asteroid
}

func (c *Canvas) Draw() {
	c.ship.Draw();

	for _, a := range c.asteroids {
		a.Draw()
	}
}