package linkedlist

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args[T Ordered] struct {
		head *ListNode[T]
		k    int
	}
	type testCase[T Ordered] struct {
		name string
		args args[T]
		want *ListNode[T]
	}
	tests := []testCase[int]{
		{
			name: "test for int",
			args: args[int]{
				head: &ListNode[int]{
					Val: 1,
					Next: &ListNode[int]{
						Val: 2,
						Next: &ListNode[int]{
							Val: 3,
							Next: &ListNode[int]{
								Val: 4,
								Next: &ListNode[int]{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
				k: 2,
			},
			want: &ListNode[int]{
				Val: 4,
				Next: &ListNode[int]{
					Val: 5,
					Next: &ListNode[int]{
						Val: 1,
						Next: &ListNode[int]{
							Val: 2,
							Next: &ListNode[int]{
								Val:  3,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotate(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
