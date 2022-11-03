package main

import (
	"testing"
)

func Test_binarySearch(t *testing.T) {
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
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			expected: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.expected {
				t.Errorf("search() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
