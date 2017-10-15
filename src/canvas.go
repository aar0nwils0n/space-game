package main

import (
	"math/rand"

	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

type Canvas struct {
	ctx       *dom.CanvasRenderingContext2D
	ship      Ship
	asteroids []*Asteroid
	explosion *js.Object
	wormhole  Wormhole
	width     float64
	height    float64
	level     float64
}

// CreateAsteroids creates random asteroid field based on Canvas.level
func (c *Canvas) CreateAsteroids() {
	number := c.level + 3
	imageURL := "./assets/images/asteroid.png"
	img := js.Global.Get("Image").New()
	img.Set("src", imageURL)

	for i := 0; i < int(number); i++ {
		for j := 0; j < int(number); j++ {

			if rand.Float64() > 0.5 {
				continue
			}

			x := float64(800)/number*float64(j) - 100 + rand.Float64()*200
			y := float64(800)/number*float64(i) - 100 + rand.Float64()*200

			if x < 200 && y < 200 {
				continue
			}

			if x > 750 && y > 750 {
				continue
			}
			a := Asteroid{}
			a.img = img
			a.canvas = c
			a.radius = (rand.Float64() + 0.25) * 50
			a.y = y
			a.x = x
			c.asteroids = append(c.asteroids, &a)
		}
	}
}

func (c *Canvas) levelUp() {
	c.level++
	c.ship.reset()
	c.CreateAsteroids()
}

func (c *Canvas) Initialize() {
	c.ship.Initialize()
	c.CreateAsteroids()
	c.explosion = js.Global.Get("Image").New()
	c.explosion.Set("src", "./assets/images/explosion.png")
	c.wormhole = Wormhole{}
	c.wormhole.canvas = c
	c.wormhole.init()
}

func (c *Canvas) Draw() {
	if c.wormhole.intersects(&c.ship) == true {
		c.levelUp()
	}
	c.ctx.ClearRect(0, 0, int(c.width), int(c.height))
	c.ship.Draw()
	c.wormhole.draw()
	for _, a := range c.asteroids {
		a.intersects(&c.ship)
		a.Draw()
	}
}
