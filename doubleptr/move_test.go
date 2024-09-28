package doubleptr

import (
	"reflect"
	"testing"
)

func Test_moveZero(t *testing.T) {
	type args struct {
		num []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test for moveZero",
			args: args{num: []int{0, 1, 0, 3, 12}},
			want: []int{1, 3, 12, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveZero(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
