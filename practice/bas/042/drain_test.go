package practice

import (
	"reflect"
	"testing"
)

func TestDrainPreservesOrder(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	want := []int{1, 2, 3, 4, 5}

	got := Drain(input)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Drain(%v) = %v, want %v", input, got, want)
	}
}

func TestDrainEmpty(t *testing.T) {
	input := []int{}
	got := Drain(input)

	if len(got) != 0 {
		t.Fatalf("Drain(%v) = %v, want empty slice", input, got)
	}
}
