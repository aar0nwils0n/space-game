package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShipCycle(t *testing.T) {
	assert := assert.New(t)
	ks := KeyboardState{}
	ship := Ship{x: 0, y: 0, ks: &ks, acceleration: 0.25}

	//It should turn around
	ks.up = true
	ship.cycle()
	assert.Equal(ship.y, ship.acceleration)
	ship.rotation = math.Pi //180 deg
	ship.cycle()
	assert.Equal(ship.y, float64(0))

	//It should accelerate straight forward
	ship.rotation = 0
	ship.ySpeed = 0
	ship.cycle()
	ship.cycle()
	ship.cycle()
	assert.Equal(ship.ySpeed, ship.acceleration*3)
	assert.Equal(ship.y, ship.acceleration+ship.acceleration*2+ship.acceleration*3)

	//It should accelerate up and to the left
	ship.ySpeed = 0
	ship.y = 0
	ship.rotation = math.Pi / 4
	ship.cycle()
	ship.cycle()
	ship.cycle()
	totalDistance := ship.acceleration + ship.acceleration*2 + ship.acceleration*3
	oposite := totalDistance * -math.Sin(ship.rotation)
	adjacent := totalDistance * math.Cos(ship.rotation)

	assert.Equal(int(ship.y*10000), int(oposite*10000))
	assert.Equal(int(ship.x*10000), int(adjacent*10000))

	//It should increase in rotational speed
	ship.rotation = 0
	ks.right = true
	ship.cycle()
	ship.cycle()
	ship.cycle()
	assert.Equal(ship.rotationalSpeed, ship.rotationalAcceleration*3)
	assert.Equal(ship.rotation, ship.rotationalAcceleration+ship.rotationalAcceleration*2+ship.rotationalAcceleration*3)

}
