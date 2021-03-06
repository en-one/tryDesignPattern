package simple

import (
	"reflect"
	"testing"
)

func TestNewFruitFactory(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want Fruit
	}{
		{
			name: "Apple",
			args: args{t: "Apple"},
			want: Apple{},
		},
		{
			name: "Banana",
			args: args{t: "Banana"},
			want: Banana{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFruit(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFruitFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
