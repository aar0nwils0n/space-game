package game

import (
	"fmt"

	"honnef.co/go/js/dom"
)

//KeyboardState holds the curent key state and handles key events
type KeyboardState struct {
	up    bool
	down  bool
	left  bool
	right bool
}

//HandleKeyDown responds to event listener by settings the appropriate key state to true
func (s *KeyboardState) HandleKeyDown(event dom.Event) {
	keyCode := event.(*dom.KeyboardEvent).KeyCode
	fmt.Println(keyCode)
	if keyCode == 37 {
		s.left = true
	}
	if keyCode == 38 {
		s.up = true
	}
	if keyCode == 39 {
		s.right = true
	}
	if keyCode == 40 {
		s.down = true
	}
}

//HandleKeyUp responds to event listener by settings the appropriate key state to false
func (s *KeyboardState) HandleKeyUp(event dom.Event) {
	keyCode := event.(*dom.KeyboardEvent).KeyCode
	if keyCode == 37 {
		s.left = false
	}
	if keyCode == 38 {
		s.up = false
	}
	if keyCode == 39 {
		s.right = false
	}
	if keyCode == 40 {
		s.down = false
	}
}
