package method

import (
	"reflect"
	"testing"
)

func TestNewFruitFactoryMethod(t *testing.T) {
	type args struct {
		t string
	}

	tests := []struct {
		name string
		args args
		want FruitFactoryMethod
	}{
		{
			name: "apple",
			args: args{t: "apple"},
			want: AppleFactory{},
		},
		{
			name: "banana",
			args: args{t: "banana"},
			want: BananaFactory{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFruitFactoryMethod(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFruitFactoryMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}
