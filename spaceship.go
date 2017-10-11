package main

import (
	"honnef.co/go/js/dom"
	"strconv"
	"github.com/gopherjs/gopherjs/js"
	"fmt"
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
	s.x = 400;
	s.y = 400;
	s.img = js.Global.Get("Image").New()
	s.img.Set("src", "./ship.svg")
	s.img.Call("addEventListener", "load", func() {
		s.Draw()
	}, false)
}

func (s *Ship) Draw() {
	fmt.Println("In draw")
	s.ctx.ClearRect(0, 0, 800, 800);
	s.ctx.Save(); // save current state
	s.ctx.Translate(int(s.x), int(s.y));
    s.ctx.Rotate(s.rotation); // rotate
	s.ctx.Call("drawImage", s.img, -60, -60, 200, 200)
	s.ctx.Restore();
}

func (s *Ship) setRotation() {
	htmlEl := s.element.(dom.HTMLElement)
	htmlEl.Style().SetProperty("transform", "rotate(" + strconv.FormatFloat(s.rotation, 'f', 4, 64) + "rad)", "")
}
