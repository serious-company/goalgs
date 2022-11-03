package main

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}

	node := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  0,
			Next: nil,
		},
	}
	node.Next.Next = node

	tests := []struct {
		name     string
		args     args
		expected *ListNode
	}{
		{
			name: "Test-1",
			args: args{
				head: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  0,
								Next: node,
							},
						},
					},
				},
			},
			expected: node,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); reflect.DeepEqual(got, tt.expected) {
				t.Errorf("reverseList() = %v, expected %v", got, tt.expected)
			}
		})

		if got := reverseList2(tt.args.head); reflect.DeepEqual(got, tt.expected) {
			t.Errorf("reverseList2() = %v, expected %v", got, tt.expected)
		}
	}
}
