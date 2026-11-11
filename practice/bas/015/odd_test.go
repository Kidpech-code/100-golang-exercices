package practice

import "testing"

func TestIsOdd(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"zero", 0, false},
		{"one", 1, true},
		{"two", 2, false},
		{"negative three", -3, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOdd(tt.n); got != tt.want {
				t.Errorf("IsOdd(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
