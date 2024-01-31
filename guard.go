package guard

import (
	"encoding/json"
)

// Guard is a dead simple json marshal guard
type Guard[T any] struct {
	raw     []byte
	data    T
	success bool
}

// UnmarshalJSON implements the json unmarshal interface
func (g *Guard[T]) UnmarshalJSON(b []byte) error {
	g.raw = b
	err := json.Unmarshal(b, &g.data)
	if err != nil {
		return nil
	}

	g.success = true
	return nil
}

// IsSuccess indicates if the unmarshal is success
func (g *Guard[T]) IsSuccess() bool {
	return g.success
}

func (g *Guard[T]) Success(data T) {
	g.success = true
	g.data = data
}

func (g *Guard[T]) Fail(raw []byte) {
	g.raw = raw
}

// Get gets the data
func (g *Guard[T]) Get() T {
	return g.data
}

// GetPtr gets the pointer of data
func (g *Guard[T]) GetPtr() *T {
	return &g.data
}

// GetRaw gets the original marshal bytes
func (g *Guard[T]) GetRaw() []byte {
	return g.raw
}
