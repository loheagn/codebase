package io_example

import "testing"

func Test_copyExample(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			copyExample()
		})
	}
}
