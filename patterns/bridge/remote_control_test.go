package bridge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoteControl(t *testing.T) {
	rcTV := NewRemoteControl(NewTV())
	assert.True(t, rcTV.TogglePower())

	rcRadio := NewAdvancedRemoteControl(NewRadio())
	assert.True(t, rcRadio.TogglePower())
}
