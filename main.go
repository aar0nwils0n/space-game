package main

import (
	"github.com/haronius/space-ship/game"
	"honnef.co/go/js/dom"
)

func main() {
	dom.GetWindow().Document().AddEventListener("DOMContentLoaded", true, initialize)
}

func initialize(e dom.Event) {
	keyboardState := game.KeyboardState{}
	dom.GetWindow().Document().AddEventListener("keydown", true, keyboardState.HandleKeyDown)
	dom.GetWindow().Document().AddEventListener("keyup", true, keyboardState.HandleKeyUp)
	domCanvas := dom.GetWindow().Document().GetElementByID("game-canvas").(*dom.HTMLCanvasElement)
	ctx := domCanvas.GetContext2d()
	canvas := game.Canvas{Ctx: ctx, Width: 800, Height: 800}
	ship := game.Ship{Ks: &keyboardState}
	ship.Canvas = &canvas
	canvas.Ship = ship
	var sSprite game.Sprite
	sSprite = &ship
	canvas.Sprites = append(canvas.Sprites, sSprite)
	canvas.Initialize()

	dom.GetWindow().SetInterval(func() {
		canvas.Draw()
		canvas.Ship.Cycle()
	}, 50)
}
