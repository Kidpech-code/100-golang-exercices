package practice

import "testing"

func TestParseHumanValid(t *testing.T) {
	b := []byte(`{"Name":"Alice","Description":"Engineer"}`)
	h, err := ParseHuman(b)
	if err != nil {
		t.Fatalf("ParseHuman returned error: %v", err)
	}
	if h.Name != "Alice" || h.Description != "Engineer" {
		t.Fatalf("unexpected Human: %+v", h)
	}
}

func TestParseHumanInvalid(t *testing.T) {
	b := []byte(`{"Name":"Alice",}`)
	if _, err := ParseHuman(b); err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
}
