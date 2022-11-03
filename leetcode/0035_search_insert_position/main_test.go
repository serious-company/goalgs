package main

import (
	"testing"
)

func Test(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{
			name: "Test-1",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.expected {
				t.Errorf("searchInsert() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
