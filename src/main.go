package main

import (
	"honnef.co/go/js/dom"
	"github.com/gopherjs/gopherjs/js"
	"math"
)


func New() *js.Object {
	return js.MakeWrapper(&Ship{})
}

func cycle(ks *KeyboardState, canvas *Canvas) func() {
	return func() {
		if(canvas.ship.explodeFrame == 0) {
			oposite := math.Sin(canvas.ship.rotation) * 0.25
			adjacent := math.Cos(canvas.ship.rotation) * 0.25
			
			if(ks.up) {
				canvas.ship.ySpeed -= adjacent
				canvas.ship.xSpeed += oposite
			}

			if(ks.left) {
				canvas.ship.rotationalSpeed = canvas.ship.rotationalSpeed - 0.005
			}

			if(ks.right) {
				canvas.ship.rotationalSpeed = canvas.ship.rotationalSpeed + 0.005
			}
			
			if(canvas.ship.rotationalSpeed != 0) {
				canvas.ship.rotation = canvas.ship.rotation + canvas.ship.rotationalSpeed;
			}
			
			if(canvas.ship.xSpeed != 0) {
				canvas.ship.x += canvas.ship.xSpeed
			}
			
			if(canvas.ship.ySpeed != 0) {
				canvas.ship.y += canvas.ship.ySpeed
			}
		}
		
		canvas.Draw()
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
	
	dom.GetWindow().SetInterval(cycle(&keyboardState, &canvas), 50);
}