package json

import "encoding/json"

func Unmarshal[T any](payload []byte) *T {
	out := new(T)
	err := json.Unmarshal(payload, out)
	if err != nil {
		panic(err)
	}

	return out
}
