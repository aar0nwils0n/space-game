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
	wormhole Wormhole
	width float64
	height float64
	level float64
}

func (c *Canvas) CreateAsteroids() {
	number := c.level + 10
	for i := 1; i < int(number); i++ {
		asteroid := Asteroid{}
		asteroid.canvas = c
		asteroid.CreateRandom();
		c.asteroids = append(c.asteroids, &asteroid)
	}
}

func(c *Canvas) levelUp() {
	c.level++
	c.ship.reset()
	c.CreateAsteroids()
}

func (c *Canvas) Initialize() {
	c.ship.Initialize()
	c.CreateAsteroids()
	c.explosion = js.Global.Get("Image").New()
	c.explosion.Set("src", "./explosion.png")
	c.wormhole = Wormhole{}
	c.wormhole.canvas = c
	c.wormhole.init()
}

func (c *Canvas) Draw() {
	if(c.wormhole.intersects(&c.ship) == true) {
		c.levelUp()
	}
	c.ctx.ClearRect(0, 0, int(c.width), int(c.height));
	c.ship.Draw();
	c.wormhole.draw();
	for _, a := range c.asteroids {
		a.Intersects(&c.ship)
		a.Draw()
	}
}