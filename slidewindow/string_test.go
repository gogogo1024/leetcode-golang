package slidewindow

import (
	"reflect"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lengthOfLongestSubstring1",
			args: args{s: "abcabcbb"},
			want: 3,
		}, {
			name: "lengthOfLongestSubstring2",
			args: args{s: "bbbbb"},
			want: 1,
		},
		{
			name: "lengthOfLongestSubstring3",
			args: args{s: "pwwkew"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shortestSeq(t *testing.T) {
	type args[T Ordered] struct {
		big   []T
		small []T
	}
	type testCase[T Ordered] struct {
		name string
		args args[T]
		want []int
	}
	tests := []testCase[int]{
		{
			name: "shortestSeq1",
			args: args[int]{big: []int{7, 5, 9, 0, 2, 1, 3, 5, 7, 9, 1, 1, 5, 8, 8, 9, 7}, small: []int{1, 5, 9}},
			want: []int{7, 10},
		}, {
			name: "shortestSeq2",
			args: args[int]{big: []int{1, 2, 3}, small: []int{4}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestSeq(tt.args.big, tt.args.small); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "minWindow",
			args: args{s: "ADOBECODEBANC", t: "ABC"},
			want: "BANC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
