package main

import (
	"reflect"
	"testing"
)

func Test_isSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected bool
	}{
		{
			name: "Test-1",
			args: []string{
				"abc",
				"ahbgdc",
			},
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args[0], tt.args[1]); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("isSubsequence() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
