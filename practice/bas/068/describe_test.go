package practice

import "testing"

func TestDescribe(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  string
	}{
		{"int", 42, "int"},
		{"string", "hello", "string"},
		{"bool", true, "bool"},
		{"nil", nil, "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Describe(tt.input)
			if got != tt.want {
				t.Errorf("Describe(%v) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
