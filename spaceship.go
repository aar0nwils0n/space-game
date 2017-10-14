package main

import (
	"honnef.co/go/js/dom"
	"github.com/gopherjs/gopherjs/js"
)

type Ship struct {
	Exploder
	element dom.Element
	velocity float64
	direction float64
	rotationalSpeed float64
	rotation float64
	link string
	img *js.Object
}

func (s *Ship) Initialize() {
	s.radius = 50
	s.reset()
	s.img = js.Global.Get("Image").New()
	s.img.Set("src", "./ship.svg")
	s.img.Call("addEventListener", "load", func() {
		s.Draw()
	}, false)
}

func(s *Ship) reset() {
	s.x = 400
	s.y = 400
	s.rotationalSpeed = 0
	s.rotation = 0
	s.direction = 0
	s.velocity = 0
	s.direction = 0
	s.explodeFrame = 0
}

func (s *Ship) Draw() {
	if(s.explodeFrame == 0) {
		s.canvas.ctx.Save(); // save current state
		s.canvas.ctx.Translate(int(s.x), int(s.y));
		s.canvas.ctx.Rotate(s.rotation); // rotate
		s.canvas.ctx.Call("drawImage", s.img, -s.radius/2, -s.radius/2,  s.radius * 2, s.radius * 2)
		s.canvas.ctx.Restore();
	} else if (s.exploded() == false) {
		s.explode()
	}
}