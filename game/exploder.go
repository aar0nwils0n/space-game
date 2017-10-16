package game

//Exploder draws a progressing explison upon the canvas
type Exploder struct {
	Canvas       *Canvas
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
	e.Canvas.Ctx.Call("drawImage", e.Canvas.explosion, e.x-explosionSize/2, e.y-explosionSize/2, explosionSize, explosionSize)
	e.explodeFrame++
}
