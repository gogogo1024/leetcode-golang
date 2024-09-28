package doubleptr

import (
	"reflect"
	"testing"
)

func Test_parisum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name:    "parisum",
			args:    args{nums: []int{5, 6, 5}, target: 11},
			wantAns: [][]int{{5, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := parisum(tt.args.nums, tt.args.target); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("parisum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
