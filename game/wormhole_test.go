package game 

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestWormholeIntersects(t *testing.T) {
	assert := assert.New(t)
	
	w := Wormhole{}
	w.radius = 100
	w.x = 0
	w.y = 0

	s := Ship{}
	s.radius = 100
	s.x = 200
	s.y = 0

	assert.Equal(w.intersects(&s), false)


	s.x = 199
	assert.Equal(w.intersects(&s), true)

	s.x = 0
	s.y = 199
	
	assert.Equal(w.intersects(&s), true)

	s.x = 0
	s.y = 200
	assert.Equal(w.intersects(&s), false)
}