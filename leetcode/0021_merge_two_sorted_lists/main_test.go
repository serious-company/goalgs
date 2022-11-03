package main

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	tests := []struct {
		name     string
		args     []*ListNode
		expected *ListNode
	}{
		{
			name: "Test-1",
			args: []*ListNode{
				{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
				{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
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
			if got := mergeTwoLists(tt.args[0], tt.args[1]); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("mergeTwoLists() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
