package abstract

import (
	"reflect"
	"testing"
)

func TestSimpleFruitFactory_CreateApple(t *testing.T) {
	tests := []struct {
		name string
		want Apple
	}{
		{
			name: "apple",
			want: Apple{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := SimpleFruitFactory{}
			if got := j.CreateApple(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRuleParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
