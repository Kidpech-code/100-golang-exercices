package practice

import "testing"

func TestWeather(t *testing.T) {
	tests := []struct {
		name string
		temp int
		want string
	}{
		{"cold boundary below", 19, "too cold"},
		{"perfect lower boundary", 20, "perfect"},
		{"perfect upper boundary", 30, "perfect"},
		{"hot boundary above", 31, "so hot"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Weather(tt.temp); got != tt.want {
				t.Fatalf("Weather(%d) = %q, want %q", tt.temp, got, tt.want)
			}
		})
	}
}
