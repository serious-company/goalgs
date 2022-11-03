package main

import (
	"reflect"
	"testing"
)

func Test_firstBadVersion(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{
			name: "Test-1",
			args: args{
				n: 5,
			},
			expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstBadVersion(tt.args.n); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("firstBadVersion() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
