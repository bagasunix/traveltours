package responses

import (
	"encoding/json"

	"github.com/bagasunix/traveltours/pkg/errors"
)

type ListEntity[T any] struct {
	Data []T `json:"data"`
}

func (a *ListEntity[T]) ToJSON() []byte {
	j, err := json.Marshal(a)
	errors.HandlerReturnedVoid(err)
	return j
}

type ListEntityBuilder[T any] struct {
	data []T
}

func NewListEntityBuilder[T any]() *ListEntityBuilder[T] {
	a := new(ListEntityBuilder[T])
	return a
}

func (a *ListEntityBuilder[T]) SetData(data []T) *ListEntityBuilder[T] {
	a.data = data
	return a
}

func (a *ListEntityBuilder[T]) Build() *ListEntity[T] {
	b := new(ListEntity[T])
	b.Data = a.data
	return b
}
