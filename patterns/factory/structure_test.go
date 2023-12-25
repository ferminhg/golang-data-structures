package factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreator(t *testing.T) {
	c := Creator{&ConcreteCreatorA{}}
	assert.Equal(t, "ConcreteProductA", c.someOperation())
	d := Creator{&ConcreteCreatorB{}}
	assert.Equal(t, "ConcreteProductB", d.someOperation())
}
