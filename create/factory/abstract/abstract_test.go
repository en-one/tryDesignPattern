package abstract

import (
	"factory/simple"
	"reflect"
	"testing"
)

func TestSimpleFruitFactory_CreateApple(t *testing.T) {
	tests := []struct {
		name string
		want simple.Apple
	}{
		{
			name: "apple",
			want: simple.Apple{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := AbstractFruitFactory{}
			if got := j.CreateApple(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRuleParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
