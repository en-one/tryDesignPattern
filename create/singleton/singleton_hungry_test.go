package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInstance(t *testing.T) {
	assert.Equal(t, GetHungryInstance(), GetHungryInstance())
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetHungryInstance() != GetHungryInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
