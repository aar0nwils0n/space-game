package main

import (
	"honnef.co/go/js/dom"
	"github.com/gopherjs/gopherjs/js"
	"strconv"
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
}

func New() *js.Object {
	return js.MakeWrapper(&Ship{})
}

func (s *Ship) SetPosition() {
	htmlEl := s.element.(dom.HTMLElement)
	htmlEl.Style().SetProperty("left", strconv.FormatFloat(s.x, 'E', -1, 64) + "px", "")
	htmlEl.Style().SetProperty("top", strconv.FormatFloat(s.y, 'E', -1, 64) + "px", "")
}

type KeyboardState struct {
	up bool
	down bool
	left bool
	right bool
}

func (s *Ship) Initialize() {
	s.element = dom.GetWindow().Document().CreateElement("div")
	dom.GetWindow().Document().QuerySelector("body").AppendChild(s.element)
	s.element.Class().SetString("ship")
	htmlEl := s.element.(dom.HTMLElement);
	htmlEl.Style().SetProperty("background", "red", "")
	htmlEl.Style().SetProperty("width", "20px", "")
	htmlEl.Style().SetProperty("height", "20px", "")
	htmlEl.Style().SetProperty("position", "absolute", "")
	s.x = 500
	s.y = 500
	s.SetPosition()
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

func (s *Ship) setRotation() {
	htmlEl := s.element.(dom.HTMLElement)
	htmlEl.Style().SetProperty("transform", "rotate(" + strconv.FormatFloat(s.rotation, 'f', 4, 64) + "deg)", "")
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

func cycle(ks *KeyboardState, ship *Ship) func() {
	return func() {
		if(ks.up) {
			ship.velocity = ship.velocity + 0.25
		}

		if(ks.down) {
			ship.velocity = ship.velocity - 0.25
		}

		if(ks.left) {
			ship.rotationalSpeed = ship.rotationalSpeed - 0.25
		}

		if(ks.right) {
			ship.rotationalSpeed = ship.rotationalSpeed + 0.25
		}

		if(ship.rotationalSpeed != 0) {
			fmt.Println(ship.rotation, ship.rotationalSpeed)
			ship.rotation = ship.rotation + ship.rotationalSpeed;
			ship.setRotation()
		}
		
		if(ship.velocity != 0) {
			ship.y = ship.y - ship.velocity
			ship.SetPosition()
		}


		dom.GetWindow().SetTimeout(cycle(ks, ship), 50);
	}
}

func main() {
	js.Global.Set("ship", map[string]interface{}{
		"New": New,
	});

	dom.GetWindow().Document().AddEventListener("DOMContentLoaded", true, Initialize);
}

func Initialize(e dom.Event) {
	keyboardState := KeyboardState{};
	dom.GetWindow().Document().AddEventListener("keydown", true, keyboardState.handleKeyDown)
	dom.GetWindow().Document().AddEventListener("keyup", true, keyboardState.handleKeyUp)
	ship := Ship{}
	ship.Initialize();
	
	cycle(&keyboardState, &ship)();
}