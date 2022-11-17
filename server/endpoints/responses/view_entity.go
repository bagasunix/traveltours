package responses

import (
	"encoding/json"

	"github.com/bagasunix/traveltours/pkg/errors"
)

type ViewEntity[T any] struct {
	Data T `json:"data"`
}

func (a *ViewEntity[T]) ToJSON() []byte {
	j, err := json.Marshal(a)
	errors.HandlerReturnedVoid(err)
	return j
}

type ViewEntityBuilder[T any] struct {
	data T
}

func NewViewEntityBuilder[T any]() *ViewEntityBuilder[T] {
	a := new(ViewEntityBuilder[T])
	return a
}

func (a *ViewEntityBuilder[T]) SetData(data T) *ViewEntityBuilder[T] {
	a.data = data
	return a
}

func (a *ViewEntityBuilder[T]) Build() *ViewEntity[T] {
	b := new(ViewEntity[T])
	b.Data = a.data
	return b
}
