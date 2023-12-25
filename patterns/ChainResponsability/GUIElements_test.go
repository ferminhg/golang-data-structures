package ChainResponsability

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExample(t *testing.T) {
	result := buildPanel()
	result.ShowHelp()

	assert.Equal(t, "Panel Help", result.ShowHelp())
}
