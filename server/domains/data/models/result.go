package models

type SingleResult[Model any] struct {
	Value Model
	Error error
}

type SliceResult[Model any] struct {
	Value []Model
	Error error
}
