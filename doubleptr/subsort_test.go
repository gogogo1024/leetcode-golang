package doubleptr

import (
	"reflect"
	"testing"
)

func Test_findUnsortedSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "findUnsortedSubarray",
			args: args{nums: []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}},
			want: []int{3, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findUnsortedSubarray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findUnsortedSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
