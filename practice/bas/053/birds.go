package practice

import (
	"encoding/json"
	"errors"
)

func Birds(b []byte) (map[string]any, error) {
	var data map[string]any
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}

	birds, ok := data["birds"].(map[string]any)
	if !ok {
		return nil, errors.New("birds key is missing or not an object")
	}

	return birds, nil
}
