package main

import (
	"testing"
)

func Test(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{
			name: "Test-1",
			args: args{
				nums: []int{1, 7, 3, 6, 5, 6},
			},
			expected: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex(tt.args.nums); got != tt.expected {
				t.Errorf("pivotIndex() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
