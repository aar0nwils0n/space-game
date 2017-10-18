package game

import (
	"math"
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
	vh        float64
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

			x := float64(c.Height)/number*float64(j) - 12*c.vh + rand.Float64()*24*c.vh
			y := float64(c.Height)/number*float64(i) - 12*c.vh + rand.Float64()*24*c.vh

			if x < 15*c.vh && y < 15*c.vh {
				continue
			}

			if x > 85*c.vh && y > 85*c.vh {
				continue
			}
			a := Asteroid{}
			a.img = img
			a.Canvas = c
			a.radius = (rand.Float64() + 0.25) * 3 * c.vh
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
	c.vh = c.Height / 100
	c.Ship.Initialize()
	c.createAsteroids()
	c.explosion = js.Global.Get("Image").New()
	c.explosion.Set("src", "./assets/images/explosion.png")
	c.wormhole = Wormhole{}
	c.wormhole.canvas = c
	c.wormhole.init()
}

func intersects(x1 float64, y1 float64, r1 float64, x2 float64, y2 float64, r2 float64) bool {
	xDistance := math.Abs(x1 - x2)
	yDistance := math.Abs(y1 - y2)
	hypot := math.Hypot(xDistance, yDistance)
	distance := hypot - (r1 + r2)
	return distance < 0
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
