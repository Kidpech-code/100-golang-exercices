package practice

import "testing"

func TestMiddle3(t *testing.T) {
	tests := []struct {
		in   []int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 3, 4}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{3, 4, 5}},
		{[]int{1, 2}, []int{1, 2}},
	}

	for _, tt := range tests {
		got := Middle3(tt.in)
		if !equalInts(got, tt.want) {
			t.Errorf("Middle3(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestMiddle3DoesNotMutateInput(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	original := []int{1, 2, 3, 4, 5}
	_ = Middle3(in)
	if !equalInts(in, original) {
		t.Errorf("Middle3 mutated input: got %v, want %v", in, original)
	}
}

func equalInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
