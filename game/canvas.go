package game

import (
	"math"
	"math/rand"

	"github.com/gopherjs/gopherjs/js"
	"github.com/haronius/space-ship/audio"
	"honnef.co/go/js/dom"
)

//Canvas holds all items to be drawn along with the state of the current level and context. Holds logic to generate items
type Canvas struct {
	Ctx       *dom.CanvasRenderingContext2D
	Ship      *Ship
	asteroids []*Asteroid
	explosion *js.Object
	wormhole  *Wormhole
	blackhole *Wormhole
	Sprites   []Sprite
	Width     int
	Height    int
	Level     int
	vh        float64
	vw        float64
	audio     *audio.Store
}

//Sprite is an object that is drawn on the canvas
type Sprite interface {
	Draw()
}

// CreateAsteroids creates random asteroid field based on Canvas.level
func (c *Canvas) createAsteroids() {
	imageURL := "./assets/images/asteroid.png"
	img := js.Global.Get("Image").New()
	img.Set("src", imageURL)
	audio := audio.CreateStore()
	audio.Add("explosion", "./assets/audio/explosion-short.mp3")

	var distance int
	if c.Level > 13 {
		distance = int(Round(14 * c.vh))
	} else {
		distance = int(Round(float64(28-c.Level) * c.vh))
	}

	for i := 0; i <= c.Width; i = i + distance {
		for j := 0; j <= c.Height; j = j + distance {
			if rand.Float64() > 0.5 {
				continue
			}

			x := float64(i) - 10*c.vh + rand.Float64()*20*c.vh
			y := float64(j) - 10*c.vh + rand.Float64()*20*c.vh

			if x < 20*c.vh && y < 20*c.vh {
				continue
			}

			if x > float64(c.Width)-25*c.vh && y > float64(c.Height)-25*c.vh {
				continue
			}
			a := Asteroid{}
			a.img = img
			a.Canvas = c
			a.radius = (rand.Float64()+0.25)*2.5*c.vh + 1.5*c.vh
			a.y = y
			a.x = x
			var s Sprite
			s = &a
			c.Sprites = append(c.Sprites, s)
			c.asteroids = append(c.asteroids, &a)
			a.audio = &audio
		}
	}
}

func (c *Canvas) levelUp() {
	c.audio.Files["whoosh"].Play()
	c.Level++
	c.Reset()
}

//Initialize creates all elements within canvas
func (c *Canvas) Initialize() {
	c.Level = 0
	c.vh = float64(c.Height) / 100
	c.vw = float64(c.Width) / 100
	c.Ship.Initialize()
	c.createAsteroids()
	c.explosion = js.Global.Get("Image").New()
	c.explosion.Set("src", "./assets/images/explosion.png")
	c.wormhole = createWormhole(c, 100, 100, 10, "./assets/images/wormhole.png")
	c.Sprites = append(c.Sprites, c.wormhole)
	c.blackhole = createWormhole(c, -10, -10, 10, "./assets/images/blackhole.png")
	c.Sprites = append(c.Sprites, c.blackhole)
	audio := audio.CreateStore()
	audio.Add("whoosh", "./assets/audio/whoosh.mp3")
	c.audio = &audio
}

//Reset the canvas state
func (c *Canvas) Reset() {
	c.asteroids = nil
	c.Sprites = nil
	c.Ship.reset()
	c.blackhole.radius = 10 * c.vh
	c.Sprites = append(c.Sprites, c.Ship)
	c.Sprites = append(c.Sprites, c.wormhole)
	c.Sprites = append(c.Sprites, c.blackhole)
	c.createAsteroids()
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

	if c.wormhole.intersects(c.Ship) == true {
		c.levelUp()
	}

	if c.blackhole.intersects(c.Ship) == true && c.Ship.explodeFrame == 0 {
		c.Ship.startExplosion()
	}

	c.blackhole.radius = c.blackhole.radius + 0.025*c.vh*float64(c.Level+1)
	c.blackhole.x = c.blackhole.x + 0.025*c.vh*float64(c.Level+1)
	c.blackhole.x = c.blackhole.y + 0.025*c.vh*float64(c.Level+1)

	c.Ctx.ClearRect(0, 0, int(c.Width), int(c.Height))

	for _, a := range c.asteroids {
		a.intersects(c.Ship)
		if intersects(c.blackhole.x, c.blackhole.y, c.blackhole.radius, a.x, a.y, a.radius) == true && a.explodeFrame == 0 {
			a.startExplosion()
		}
	}

	for _, s := range c.Sprites {
		s.Draw()
	}
}
