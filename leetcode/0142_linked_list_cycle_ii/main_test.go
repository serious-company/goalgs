package main

import (
	"testing"
)

func Test_detectCycle(t *testing.T) {
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
			if got := detectCycle(tt.args.head); got != tt.expected {
				t.Errorf("detectCycle() = %v, expected %v", got, tt.expected)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle2(tt.args.head); got != tt.expected {
				t.Errorf("detectCycle2() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
