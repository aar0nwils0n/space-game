package main

import (
	"strconv"

	"github.com/haronius/space-ship/audio"
	"github.com/haronius/space-ship/game"
	"honnef.co/go/js/dom"
)

func main() {
	dom.GetWindow().Document().AddEventListener("DOMContentLoaded", true, initialize)
}

func getCanvasSize() (float64, float64) {
	bodyRect := dom.GetWindow().Document().GetElementByID("game-body").GetBoundingClientRect()
	controlRect := dom.GetWindow().Document().GetElementByID("controls").GetBoundingClientRect()
	height := bodyRect.Height - controlRect.Height

	return bodyRect.Width, height
}

func createGameCanvas() *game.Canvas {
	width, height := getCanvasSize()
	keyboardState := game.KeyboardState{}
	domCanvas := dom.GetWindow().Document().GetElementByID("game-canvas").(*dom.HTMLCanvasElement)
	domCanvas.SetAttribute("width", strconv.FormatFloat(width, 'f', 6, 64))
	domCanvas.SetAttribute("height", strconv.FormatFloat(height, 'f', 6, 64))
	ctx := domCanvas.GetContext2d()
	canvas := game.Canvas{Ctx: ctx, Height: int(height), Width: int(width)}
	ship := game.Ship{Ks: &keyboardState}
	ship.Canvas = &canvas
	canvas.Ship = &ship
	var sSprite game.Sprite
	sSprite = canvas.Ship
	canvas.Sprites = append(canvas.Sprites, sSprite)
	return &canvas
}

func startSound() {
	store := audio.CreateStore()
	store.Add("bg", "assets/audio/Papergirl.mp3")
	store.Files["bg"].LoopFull()
}

func initialize(e dom.Event) {
	startSound()
	canvas := createGameCanvas()
	dom.GetWindow().Document().AddEventListener("keydown", true, canvas.Ship.Ks.HandleKeyDown)
	dom.GetWindow().Document().AddEventListener("keyup", true, canvas.Ship.Ks.HandleKeyUp)

	left := dom.GetWindow().Document().GetElementByID("left")
	left.AddEventListener("touchstart", true, canvas.Ship.Ks.SetLeftTrue)
	left.AddEventListener("touchend", true, canvas.Ship.Ks.SetLeftFalse)
	left.AddEventListener("mousedown", true, canvas.Ship.Ks.SetLeftTrue)
	left.AddEventListener("mouseup", true, canvas.Ship.Ks.SetLeftFalse)

	up := dom.GetWindow().Document().GetElementByID("up")
	up.AddEventListener("touchstart", true, canvas.Ship.Ks.SetUpTrue)
	up.AddEventListener("touchend", true, canvas.Ship.Ks.SetUpFalse)
	up.AddEventListener("mousedown", true, canvas.Ship.Ks.SetUpTrue)
	up.AddEventListener("mouseup", true, canvas.Ship.Ks.SetUpFalse)

	right := dom.GetWindow().Document().GetElementByID("right")
	right.AddEventListener("touchstart", true, canvas.Ship.Ks.SetRightTrue)
	right.AddEventListener("touchend", true, canvas.Ship.Ks.SetRightFalse)
	right.AddEventListener("mousedown", true, canvas.Ship.Ks.SetRightTrue)
	right.AddEventListener("mouseup", true, canvas.Ship.Ks.SetRightFalse)

	canvas.Initialize()

	interval := setInterval(canvas)
	handlePlayPause(interval, canvas)

	reset := dom.GetWindow().Document().GetElementByID("reset")
	reset.AddEventListener("click", true, func(e dom.Event) {
		canvas.Level = 0
		canvas.Reset()
		dom.GetWindow().Document().GetElementByID("reset-overlay").SetAttribute("class", "reset-overlay hidden")
	})
}

func handlePlayPause(interval int, canvas *game.Canvas) {
	pause := dom.GetWindow().Document().GetElementByID("pause")
	pause.AddEventListener("click", true, func(e dom.Event) {
		if pause.Class().Contains("pause") {
			pause.SetAttribute("class", "play")
			dom.GetWindow().ClearInterval(interval)
		} else {
			pause.SetAttribute("class", "pause")
			interval = setInterval(canvas)
		}
	})
}

func setInterval(canvas *game.Canvas) int {
	return dom.GetWindow().SetInterval(func() {
		canvas.Draw()
		canvas.Ship.Cycle()
	}, 50)
}
