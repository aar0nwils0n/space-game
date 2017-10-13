package main

import (
	"honnef.co/go/js/dom"
	"fmt"
	
)

type Canvas struct {
	ctx *dom.CanvasRenderingContext2D
	ship Ship
	asteroids []Asteroid
}

func (c *Canvas) CreateAsteroids(number float64) {
	fmt.Println("in create ast")
	for i := 1; i < int(number); i++ {
		asteroid := Asteroid{ctx: c.ctx}
		asteroid.CreateRandom();
		c.asteroids = append(c.asteroids, asteroid)
	}
}

func (c *Canvas) Initialize() {
	c.ship.Initialize()
	c.CreateAsteroids(10)
}

func (c *Canvas) Draw() {
	c.ctx.ClearRect(0, 0, 800, 800);
	c.ship.Draw();
	
	
	for _, a := range c.asteroids {
		a.Intersects(c.ship)
		a.Draw()
		fmt.Println(a.exploded)
	}
}