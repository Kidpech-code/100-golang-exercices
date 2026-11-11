package practice

import "encoding/json"

type Human struct {
	Name        string
	Description string
}

func ParseHuman(b []byte) (Human, error) {
	var h Human
	if err := json.Unmarshal(b, &h); err != nil {
		return Human{}, err
	}
	return h, nil
}
