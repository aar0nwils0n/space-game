package main

import (
	"strconv"

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
	rect := dom.GetWindow().Document().GetElementByID("game-body").GetBoundingClientRect()
	domCanvas := dom.GetWindow().Document().GetElementByID("game-canvas").(*dom.HTMLCanvasElement)
	domCanvas.SetAttribute("width", strconv.FormatFloat(rect.Height, 'f', 6, 64))
	domCanvas.SetAttribute("height", strconv.FormatFloat(rect.Height, 'f', 6, 64))
	ctx := domCanvas.GetContext2d()
	canvas := game.Canvas{Ctx: ctx, Height: rect.Height, Width: rect.Height}
	ship := game.Ship{Ks: &keyboardState}
	ship.Canvas = &canvas
	canvas.Ship = ship
	var sSprite game.Sprite
	sSprite = &canvas.Ship
	canvas.Sprites = append(canvas.Sprites, sSprite)
	canvas.Initialize()

	dom.GetWindow().SetInterval(func() {
		canvas.Draw()
		canvas.Ship.Cycle()
	}, 50)
}
