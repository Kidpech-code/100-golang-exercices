package practice

import (
	"reflect"
	"testing"
)

func TestFib(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"n=0", 0, []int{}},
		{"n=1", 1, []int{0}},
		{"n=2", 2, []int{0, 1}},
		{"n=7", 7, []int{0, 1, 1, 2, 3, 5, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Fib(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Fib(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
