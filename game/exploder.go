package game

import (
	"github.com/haronius/space-ship/audio"
)

//Exploder draws a progressing explison upon the canvas
type Exploder struct {
	Canvas       *Canvas
	x            float64
	y            float64
	radius       float64
	explodeFrame float64
	audio        *audio.Store
}

func (e *Exploder) exploded() bool {
	return e.explodeFrame >= 10
}

func (e *Exploder) startExplosion() {
	file := e.audio.Files["explosion"]
	file.Play()
	e.explodeFrame = 1
}

func (e *Exploder) explode() {
	explosionSize := e.radius*2 + e.explodeFrame*e.Canvas.vh*5
	e.Canvas.Ctx.Call("drawImage", e.Canvas.explosion, e.x-explosionSize/2, e.y-explosionSize/2, explosionSize, explosionSize)
	e.explodeFrame++
}
