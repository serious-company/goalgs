package main

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name: "Test-1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.expected) {
				t.Errorf("rotate(): %v, expected: %v", tt.args.nums, tt.expected)
			}
		})
	}
}
