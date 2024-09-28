package doubleptr

import (
	"reflect"
	"testing"
)

func Test_sortColor2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sortColor",
			args: args{nums: []int{2, 0, 2, 1, 1, 0}},
			want: []int{0, 0, 1, 1, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortColor(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
