package main

import (
	"reflect"
	"testing"
)

func Test_runningSum(t *testing.T) {
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
				nums: []int{1, 2, 3, 4},
			},
			expected: []int{1, 3, 6, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runningSum(tt.args.nums); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("runningSum() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
