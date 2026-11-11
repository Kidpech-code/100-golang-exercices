package practice

import (
	"reflect"
	"testing"
)

func TestExercise053BirdsValid(t *testing.T) {
	input := []byte(`{"birds":{"robin":"coo","eagle":"screech"}}`)
	got, err := Birds(input)
	if err != nil {
		t.Fatalf("Birds() error = %v", err)
	}
	want := map[string]any{"robin": "coo", "eagle": "screech"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Birds() = %#v, want %#v", got, want)
	}
}

func TestExercise053BirdsMissingKey(t *testing.T) {
	if _, err := Birds([]byte(`{"cats":{"siamese":"meow"}}`)); err == nil {
		t.Fatal("Birds() error = nil, want error for missing key")
	}
}

func TestExercise053BirdsWrongType(t *testing.T) {
	inputs := []string{
		`{"birds":["robin","eagle"]}`,
		`{"birds":"not-a-map"}`,
		`{"birds":42}`,
		`{"birds":true}`,
		`{"birds":null}`,
	}
	for _, input := range inputs {
		if _, err := Birds([]byte(input)); err == nil {
			t.Fatalf("Birds(%s) error = nil, want error for wrong type", input)
		}
	}
}

func TestExercise053BirdsInvalidJSON(t *testing.T) {
	if _, err := Birds([]byte(`{invalid`)); err == nil {
		t.Fatal("Birds() error = nil, want error for invalid JSON")
	}
}
