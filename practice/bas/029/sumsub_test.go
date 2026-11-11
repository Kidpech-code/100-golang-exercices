package practice

import "testing"

func TestSumSub(t *testing.T) {
	tests := []struct {
		name             string
		a, b             int
		wantSum, wantSub int
	}{
		{"positive", 5, 3, 8, 2},
		{"negative", -5, -3, -8, -2},
		{"mixed", 5, -3, 2, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSum, gotSub := SumSub(tt.a, tt.b)
			if gotSum != tt.wantSum || gotSub != tt.wantSub {
				t.Fatalf("SumSub(%d, %d) = (%d, %d), want (%d, %d)", tt.a, tt.b, gotSum, gotSub, tt.wantSum, tt.wantSub)
			}
		})
	}
}
