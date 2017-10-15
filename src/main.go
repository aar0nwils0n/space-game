package main

import (
	"honnef.co/go/js/dom"
	"github.com/gopherjs/gopherjs/js"
	"math"
)


func New() *js.Object {
	return js.MakeWrapper(&Ship{})
}

func (s *Ship) cycle() func() {
	return func() {
		if(s.explodeFrame == 0) {
			oposite := math.Sin(s.rotation) * 0.25
			adjacent := math.Cos(s.rotation) * 0.25
			
			if(s.ks.up) {
				s.ySpeed -= adjacent
				s.xSpeed += oposite
			}

			if(s.ks.left) {
				s.rotationalSpeed = s.rotationalSpeed - 0.01
			}

			if(s.ks.right) {
				s.rotationalSpeed = s.rotationalSpeed + 0.01
			}
			
			if(s.rotationalSpeed != 0) {
				s.rotation = s.rotation + s.rotationalSpeed;
			}
			
			if(s.xSpeed != 0) {
				s.x += s.xSpeed
			}
			
			if(s.ySpeed != 0) {
				s.y += s.ySpeed
			}
		}
		
		s.canvas.Draw()
	}
}

func main() {
	dom.GetWindow().Document().AddEventListener("DOMContentLoaded", true, Initialize);
}

func Initialize(e dom.Event) {
	keyboardState := KeyboardState{};
	dom.GetWindow().Document().AddEventListener("keydown", true, keyboardState.handleKeyDown)
	dom.GetWindow().Document().AddEventListener("keyup", true, keyboardState.handleKeyUp)
	domCanvas := dom.GetWindow().Document().GetElementByID("game-canvas").(*dom.HTMLCanvasElement)
	ctx := domCanvas.GetContext2d()
	canvas := Canvas{ctx: ctx, width: 800, height: 800}
	ship := Ship{ks: &keyboardState}
	ship.canvas = &canvas
	canvas.ship = ship
	canvas.Initialize();
	
	dom.GetWindow().SetInterval(canvas.ship.cycle(), 50);
}