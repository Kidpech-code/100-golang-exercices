package practice

import "testing"

func TestConcat(t *testing.T) {
	tests := []struct {
		name string
		a, b string
		want string
	}{
		{"both non-empty", "Hello World!", "It's a me, Mario", "Hello World! It's a me, Mario"},
		{"empty left", "", "world", " world"},
		{"empty right", "hello", "", "hello "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concat(tt.a, tt.b); got != tt.want {
				t.Fatalf("Concat(%q, %q) = %q, want %q", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
