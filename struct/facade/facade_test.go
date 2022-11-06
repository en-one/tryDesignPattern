package facade

import (
	"testing"
)

func TestNewSampleFacade(t *testing.T) {
	facade := NewSampleFacade()
	facade.sendLetter(
		"hello Song, this is facade pattern testing...",
		"austsxk@vip.qq.com")
}
