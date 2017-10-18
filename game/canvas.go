package game

import (
	"math/rand"

	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

//Canvas holds all items to be drawn along with the state of the current level and context. Holds logic to generate items
type Canvas struct {
	Ctx       *dom.CanvasRenderingContext2D
	Ship      Ship
	asteroids []*Asteroid
	explosion *js.Object
	wormhole  Wormhole
	Width     float64
	Height    float64
	level     float64
}

// CreateAsteroids creates random asteroid field based on Canvas.level
func (c *Canvas) createAsteroids() {
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
			a.Canvas = c
			a.radius = (rand.Float64() + 0.25) * 50
			a.y = y
			a.x = x
			c.asteroids = append(c.asteroids, &a)
		}
	}
}

func (c *Canvas) levelUp() {
	c.level++
	c.Ship.reset()
	c.createAsteroids()
}

//Initialize creates all elements within canvas
func (c *Canvas) Initialize() {
	c.Ship.Initialize()
	c.createAsteroids()
	c.explosion = js.Global.Get("Image").New()
	c.explosion.Set("src", "./assets/images/explosion.png")
	c.wormhole = Wormhole{}
	c.wormhole.canvas = c
	c.wormhole.init()
}

//Draw checks for intersecting items then draws them on the canvas
func (c *Canvas) Draw() {
	if c.wormhole.intersects(&c.Ship) == true {
		c.levelUp()
	}
	c.Ctx.ClearRect(0, 0, int(c.Width), int(c.Height))
	c.Ship.draw()
	c.wormhole.draw()
	for _, a := range c.asteroids {
		a.intersects(&c.Ship)
		a.draw()
	}
}
