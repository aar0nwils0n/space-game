package main

import (
	"honnef.co/go/js/dom"
	"github.com/gopherjs/gopherjs/js"
)

type Ship struct {
	Exploder
	element dom.Element
	xSpeed float64
	ySpeed float64
	rotationalSpeed float64
	rotation float64
	link string
	ship *js.Object
	shipEngineOn *js.Object
	ks *KeyboardState
}

func (s *Ship) Initialize() {
	s.radius = 30
	s.reset()
	s.ship = js.Global.Get("Image").New()
	s.ship.Set("src", "./assets/images/ship.svg")
	s.shipEngineOn = js.Global.Get("Image").New()
	s.shipEngineOn.Set("src", "./assets/images/ship-engine-on.svg")
}

func(s *Ship) reset() {
	s.x = 400
	s.y = 400
	s.rotationalSpeed = 0
	s.rotation = 0
	s.xSpeed = 0
	s.ySpeed = 0
	s.explodeFrame = 0
}

func (s *Ship) Draw() {
	if(s.explodeFrame == 0) {
		s.canvas.ctx.Save(); // save current state
		s.canvas.ctx.Translate(int(s.x), int(s.y));
		s.canvas.ctx.Rotate(s.rotation); // rotate

		var img *js.Object

		if(s.ks.up) {
			img = s.shipEngineOn
		} else {
			img = s.ship
		}
		s.canvas.ctx.Call("drawImage", img, -s.radius * 1.25, -s.radius *1.25,  s.radius * 2 * 1.25, s.radius * 2 * 1.25)
		s.canvas.ctx.Restore();
	} else if (s.exploded() == false) {
		s.explode()
	}
}