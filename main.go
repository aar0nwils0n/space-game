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
			canvas.ship.velocity = ship.velocity + 0.25
		}

		if(ks.down) {
			canvas.ship.velocity = ship.velocity - 0.25
		}

		if(ks.left) {
			canvas.ship.rotationalSpeed = ship.rotationalSpeed - 0.005
		}

		if(ks.right) {
			canvas.ship.rotationalSpeed = ship.rotationalSpeed + 0.005
		}
		

		if(canvas.ship.rotationalSpeed != 0) {
			canvas.ship.rotation = ship.rotation + ship.rotationalSpeed;
		}
		
		if(canvas.ship.velocity != 0) {
			oposite := math.Sin(ship.rotation) * ship.velocity
			adjacent := math.Cos(ship.rotation) * ship.velocity
			canvas.ship.y = ship.y - adjacent
			canvas.ship.x = ship.x + oposite

		}

		canvas.Draw()
	}
}

func main() {
	js.Global.Set("ship", map[string]interface{}{
		"New": New,
	});

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
	ship.Initialize();
	
	dom.GetWindow().SetInterval(cycle(&keyboardState, &canvas), 50);
}