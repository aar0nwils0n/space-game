package main

type Exploder struct {
	canvas       *Canvas
	x            float64
	y            float64
	radius       float64
	explodeFrame float64
}

func (e *Exploder) exploded() bool {
	return e.explodeFrame >= 10
}

func (e *Exploder) explode() {
	explosionSize := e.radius * e.explodeFrame * 2
	e.canvas.ctx.Call("drawImage", e.canvas.explosion, e.x-explosionSize/2, e.y-explosionSize/2, explosionSize, explosionSize)
	e.explodeFrame++
}
