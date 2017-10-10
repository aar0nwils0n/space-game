package main

import (
	"honnef.co/go/js/dom"
	"github.com/gopherjs/gopherjs/js"
	"math"
)


func New() *js.Object {
	return js.MakeWrapper(&Ship{})
}

func cycle(ks *KeyboardState, ship *Ship) func() {
	return func() {
		if(ks.up) {
			ship.velocity = ship.velocity + 0.25
		}

		if(ks.down) {
			ship.velocity = ship.velocity - 0.25
		}

		if(ks.left) {
			ship.rotationalSpeed = ship.rotationalSpeed - 0.005
		}

		if(ks.right) {
			ship.rotationalSpeed = ship.rotationalSpeed + 0.005
		}
		

		if(ship.rotationalSpeed != 0) {
			ship.rotation = ship.rotation + ship.rotationalSpeed;
			ship.setRotation()
		}
		
		if(ship.velocity != 0) {
			oposite := math.Sin(ship.rotation) * ship.velocity
			adjecent := math.Cos(ship.rotation) * ship.velocity
			ship.y = ship.y - oposite
			ship.x = ship.x - adjecent

			ship.SetPosition()
		}
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
	canvas := dom.GetWindow().Document().GetElementByID("game-canvas").(dom.HTMLCanvasElement)
	ctx := canvas.GetContext2d()
	ship := Ship{ctx: ctx}
	ship.Initialize();
	
	dom.GetWindow().SetInterval(cycle(&keyboardState, &ship), 50);
}