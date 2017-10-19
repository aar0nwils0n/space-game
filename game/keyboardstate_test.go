package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"honnef.co/go/js/dom"
)

func TestHandleKeyDownUp(t *testing.T) {
	leftCode := 37
	upCode := 38
	rightCode := 39
	downCode := 40

	assert := assert.New(t)

	keyEvent := dom.KeyboardEvent{}
	var event dom.Event
	event = &keyEvent
	keyEvent.KeyCode = upCode
	ks := KeyboardState{}

	ks.HandleKeyDown(event)

	assert.Equal(ks.up, true)

	ks.HandleKeyUp(event)

	assert.Equal(ks.up, false)

	keyEvent.KeyCode = downCode

	ks.HandleKeyDown(event)

	assert.Equal(ks.down, true)

	ks.HandleKeyUp(event)

	assert.Equal(ks.down, false)

	keyEvent.KeyCode = leftCode

	ks.HandleKeyDown(event)

	assert.Equal(ks.left, true)

	ks.HandleKeyUp(event)

	assert.Equal(ks.left, false)

	keyEvent.KeyCode = rightCode

	ks.HandleKeyDown(event)

	assert.Equal(ks.right, true)

	ks.HandleKeyUp(event)

	assert.Equal(ks.right, false)

}
