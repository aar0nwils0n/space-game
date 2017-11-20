package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsteroidIntersects(t *testing.T) {
	assert := assert.New(t)
	store := createTestAudioStore()
	a := Asteroid{}
	a.audio = &store
	a.radius = 100
	a.x = 0
	a.y = 0

	s := Ship{}
	s.audio = &store
	s.radius = 100
	s.x = 200
	s.y = 0

	assert.Equal(a.intersects(&s), false)

	s.x = 199
	assert.Equal(a.intersects(&s), true)

	s.x = 0
	s.y = 199

	assert.Equal(a.intersects(&s), true)

	s.x = 0
	s.y = 200
	assert.Equal(a.intersects(&s), false)
}
