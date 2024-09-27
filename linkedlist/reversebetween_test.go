package linkedlist

import (
	"reflect"
	"testing"
)

func isEqual[T any](l1, l2 *ListNode[T]) bool {
	for l1 != nil && l2 != nil {
		if !reflect.DeepEqual(l1.Val, l2.Val) { // 使用 reflect.DeepEqual 比较值
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil // 确保两个链表长度相同
}

func Test_reversebetwwen(t *testing.T) {
	type args[T any] struct {
		head  *ListNode[T]
		left  int
		right int
	}
	type testCase[T any] struct {
		name string
		args args[T]
		//want *ListNode[T]
		want args[T]
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
				left:  2,
				right: 4,
			},
			want: args[int]{
				head: &ListNode[int]{
					Val: 1,
					Next: &ListNode[int]{
						Val: 4,
						Next: &ListNode[int]{
							Val: 3,
							Next: &ListNode[int]{
								Val: 2,
								Next: &ListNode[int]{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversebetwwen(tt.args.head, tt.args.left, tt.args.right); !isEqual(got, tt.want.head) {
				t.Errorf("reversebetwwen() = %v, want %v", got, tt.want)
			}
		})
	}
}
