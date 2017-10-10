package main

import (
	"honnef.co/go/js/dom"
	"math"
	"strconv"
	"github.com/gopherjs/gopherjs/js"
)

type Ship struct {
	element dom.Element
	velocity float64
	x float64
	y float64
	direction float64
	rotationalSpeed float64
	rotation float64
	link string
	img *js.Object
	ctx *dom.CanvasRenderingContext2D
}

func (s *Ship) Initialize() {
	s.img = js.Global.Get("Image").New()
	s.img.Set("src", "./ship.svg")
	s.img.Call("addEventListener", "load", func() {
		s.Draw()
	}, false)
}

func (s *Ship) Draw() {
	s.ctx.Save(); // save current state
    s.ctx.Rotate(s.rotation); // rotate
	s.ctx.Call("drawImage", s.img, s.x, s.y, 30, 30)
	s.ctx.Restore();
}

func (s *Ship) SetPosition() {
	htmlEl := s.element.(dom.HTMLElement)
	htmlEl.Style().SetProperty("left", strconv.FormatFloat(s.x, 'E', -1, 64) + "px", "")
	htmlEl.Style().SetProperty("top", strconv.FormatFloat(s.y, 'E', -1, 64) + "px", "")
}


func (s *Ship) setRotation() {
	htmlEl := s.element.(dom.HTMLElement)
	htmlEl.Style().SetProperty("transform", "rotate(" + strconv.FormatFloat(s.rotation, 'f', 4, 64) + "rad)", "")
}
