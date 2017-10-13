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
		if(ks.up) {
			canvas.ship.velocity = canvas.ship.velocity + 0.25
		}

		if(ks.down) {
			canvas.ship.velocity = canvas.ship.velocity - 0.25
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
		
		if(canvas.ship.velocity != 0) {
			oposite := math.Sin(canvas.ship.rotation) * canvas.ship.velocity
			adjacent := math.Cos(canvas.ship.rotation) * canvas.ship.velocity
			canvas.ship.y = canvas.ship.y - adjacent
			canvas.ship.x = canvas.ship.x + oposite

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
	ship := Ship{ctx: ctx}
	canvas := Canvas{ctx: ctx, ship: ship}
	canvas.Initialize();
	
	dom.GetWindow().SetInterval(cycle(&keyboardState, &canvas), 50);
}