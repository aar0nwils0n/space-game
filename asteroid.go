package main

import (
	"github.com/gopherjs/gopherjs/js"
)

type Asteroid struct {
	radius int
	image string
	img *js.Object
}

func (a *Asteroid) Draw() {
	
}