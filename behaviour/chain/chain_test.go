package chain

import (
	"fmt"
	"testing"
)

func Test_chan1(t *testing.T) {
	s := NewIService()

	res, err := s.HelloWorldChain("煎鱼13231")
	fmt.Println(res, err)

	// 会返回错误输出
	res, err = s.HelloWorldChain("12")
	fmt.Println(res, err)
}

// func Test_chan12t *testing.T) {
// 	chain := &Chain{}
// 	chain.AddFilter(&AdSensitiveWordFilter{})
// 	assert.Equal(t, false, chain.Filter("test"))

// 	chain.AddFilter(&PoliticalWordFilter{})
// 	assert.Equal(t, true, chain.Filter("test"))
// }
