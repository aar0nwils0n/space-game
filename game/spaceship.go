package game

import (
	"math"

	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

//Ship is an object that moves upon the canvas according to the keyboard state
type Ship struct {
	Exploder
	element                dom.Element
	xSpeed                 float64
	ySpeed                 float64
	rotationalSpeed        float64
	rotationalAcceleration float64
	rotation               float64
	link                   string
	ship                   *js.Object
	shipEngineOn           *js.Object
	Ks                     *KeyboardState
	acceleration           float64
}

//Initialize creates ship with default properties
func (s *Ship) Initialize() {
	s.radius = 25
	s.rotation = math.Pi * 0.75
	s.reset()
	s.acceleration = 0.25
	s.rotationalAcceleration = 0.01
	s.ship = js.Global.Get("Image").New()
	s.ship.Set("src", "./assets/images/ship.svg")
	s.shipEngineOn = js.Global.Get("Image").New()
	s.shipEngineOn.Set("src", "./assets/images/ship-engine-on.svg")
}

func (s *Ship) reset() {
	s.x = 50
	s.y = 50
	s.rotationalSpeed = 0
	s.rotation = math.Pi * 0.75
	s.xSpeed = 0
	s.ySpeed = 0
	s.explodeFrame = 0
}

//Draw checks if out of bounds and exploded if needed
func (s *Ship) Draw() {

	if s.outOfBounds() == true && s.explodeFrame == 0 && s.exploded() == false {
		s.explodeFrame = 1
	}

	if s.explodeFrame == 0 {
		s.Canvas.Ctx.Save() // save current state
		s.Canvas.Ctx.Translate(int(s.x), int(s.y))
		s.Canvas.Ctx.Rotate(s.rotation) // rotate

		var img *js.Object

		if s.Ks.up {
			img = s.shipEngineOn
		} else {
			img = s.ship
		}
		s.Canvas.Ctx.Call("drawImage", img, -s.radius*1.25, -s.radius*1.25, s.radius*2*1.25, s.radius*2*1.25)
		s.Canvas.Ctx.Restore()
	} else if s.exploded() == false {
		s.explode()
	}
}

func (s *Ship) outOfBounds() bool {
	return s.x < 0 || s.y < 0 || s.y > 800 || s.x > 800
}

//Cycle checks keyboard state and moves to corresponding coordinates
func (s *Ship) Cycle() {
	if s.explodeFrame == 0 {
		oposite := math.Sin(s.rotation) * s.acceleration
		adjacent := math.Cos(s.rotation) * s.acceleration

		if s.Ks.up {
			s.ySpeed -= adjacent
			s.xSpeed += oposite
		}

		if s.Ks.left {
			s.rotationalSpeed = s.rotationalSpeed - s.rotationalAcceleration
		}

		if s.Ks.right {
			s.rotationalSpeed = s.rotationalSpeed + s.rotationalAcceleration
		}

		if s.rotationalSpeed != 0 {
			s.rotation = s.rotation + s.rotationalSpeed
		}

		if s.xSpeed != 0 {
			s.x += s.xSpeed
		}

		if s.ySpeed != 0 {
			s.y += s.ySpeed
		}
	}
}