package main

import (
	"honnef.co/go/js/dom"
)

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
	
	dom.GetWindow().SetInterval(func() {
		canvas.Draw()
		canvas.ship.cycle()
	}, 50);
}