package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name: "Test-1",
			args: args{
				nums: []int{-4, -1, 0, 3, 10},
			},
			expected: []int{0, 1, 9, 16, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.nums); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("sortedSquares() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
