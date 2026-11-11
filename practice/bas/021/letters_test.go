package practice

import "testing"

func TestLetters(t *testing.T) {
	got := Letters()

	want := map[int]string{1: "A", 2: "B", 3: "C"}
	for key, wantVal := range want {
		if got[key] != wantVal {
			t.Errorf("Letters()[%d] = %q, want %q", key, got[key], wantVal)
		}
	}

	if len(got) != len(want) {
		t.Errorf("Letters() length = %d, want %d", len(got), len(want))
	}

	if got[99] != "" {
		t.Errorf("Letters()[99] = %q, want empty string", got[99])
	}
	if _, ok := got[99]; ok {
		t.Error("Letters()[99] should not exist")
	}
}
