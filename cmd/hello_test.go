package cmd

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test with a valid name", "Alice", "Hello, Alice"},
		{"Test with an empty name", "", "Hello, "},
		{"Test with a space in the name", "Bob Smith", "Hello, Bob Smith"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Hello(tt.input)
			if result != tt.expected {
				t.Errorf("got %s, want %s", result, tt.expected)
			}
		})
	}
}
