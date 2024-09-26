package linkedlist

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	type args[T any] struct {
		head *ListNode[T]
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want args[T]
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
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
				}},
			want: args[int]{
				head: &ListNode[int]{
					Val: 5,
					Next: &ListNode[int]{
						Val: 4,
						Next: &ListNode[int]{
							Val: 3,
							Next: &ListNode[int]{
								Val: 2,
								Next: &ListNode[int]{
									Val:  1,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "test for nil",
			args: args[int]{
				head: &ListNode[int]{}},
			want: args[int]{
				head: &ListNode[int]{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want.head) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
