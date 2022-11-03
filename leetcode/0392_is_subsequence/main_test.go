package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected bool
	}{
		{
			name: "Test-1",
			args: []string{
				"egg",
				"add",
			},
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args[0], tt.args[1]); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("isIsomorphic() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
