package flyweight

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlyWeightFactory_GetFlyWeight(t *testing.T) {
	factory := NewFlyWeightFactory()

	zhang := factory.GetFlyWeight("zhang beauty")
	chen := factory.GetFlyWeight("chen beauty")

	assert.Len(t, factory.pool, 2)
	assert.Equal(t, zhang, factory.pool["zhang beauty"])
	assert.Equal(t, chen, factory.pool["chen beauty"])

}
