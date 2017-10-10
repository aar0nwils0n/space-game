package main

import (
	"honnef.co/go/js/dom"
	"fmt"
)

type KeyboardState struct {
	up bool
	down bool
	left bool
	right bool
}

func (s *KeyboardState) handleKeyDown(event dom.Event) {
	keyCode := event.(*dom.KeyboardEvent).KeyCode
	fmt.Println(keyCode)
	if(keyCode == 37) {
		s.left = true
	}
	if(keyCode == 38) {
		s.up = true
	}
	if(keyCode == 39) {
		s.right = true
	}
	if(keyCode == 40) {
		s.down = true
	}
}

func (s *KeyboardState) handleKeyUp(event dom.Event) {
	keyCode := event.(*dom.KeyboardEvent).KeyCode
	if(keyCode == 37) {
		s.left = false;
	}
	if(keyCode == 38) {
		s.up = false
	}
	if(keyCode == 39) {
		s.right = false
	}
	if(keyCode == 40) {
		s.down = false
	}
}