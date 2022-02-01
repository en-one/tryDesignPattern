package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_demo(t *testing.T) {
	chain := &Chain{}
	chain.AddFilter(&AdSensitiveWordFilter{})
	assert.Equal(t, false, chain.Filter("test"))

	chain.AddFilter(&PoliticalWordFilter{})
	assert.Equal(t, true, chain.Filter("test"))
}
