package requests

import (
	"encoding/json"

	"github.com/bagasunix/traveltours/pkg/errors"
)

type BaseList struct {
	Limit     int64   `json:"limit"`
	Keyword   string  `json:"keyword"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (c *BaseList) ToJSON() []byte {
	j, err := json.Marshal(c)
	errors.HandlerReturnedVoid(err)
	return j
}

// BaseListBuilder Builder Object for BaseList
type BaseListBuilder struct {
	limit     int64
	keyword   string
	latitude  float64
	longitude float64
}

// NewBaseListBuilder Constructor for BaseListBuilder
func NewBaseListBuilder() *BaseListBuilder {
	o := new(BaseListBuilder)
	return o
}

// Build Method which creates BaseList
func (b *BaseListBuilder) Build() *BaseList {
	o := new(BaseList)
	o.Limit = b.limit
	o.Keyword = b.keyword
	o.Latitude = b.latitude
	o.Longitude = b.longitude
	return o
}

// SetLimit Limit Builder method to set the field limit in BaseListBuilder
func (b *BaseListBuilder) SetLimit(v int64) *BaseListBuilder {
	b.limit = v
	return b
}

// SetKeyword Keyword Builder method to set the field keyword in BaseListBuilder
func (b *BaseListBuilder) SetKeyword(v string) *BaseListBuilder {
	b.keyword = v
	return b
}

// SetLatitude Setter method for the field latitude of type float64 in the object BaseListBuilder
func (b *BaseListBuilder) SetLatitude(v float64) *BaseListBuilder {
	b.latitude = v
	return b
}

// SetLongitude Setter method for the field longitude of type float64 in the object BaseListBuilder
func (b *BaseListBuilder) SetLongitude(v float64) *BaseListBuilder {
	b.longitude = v
	return b
}
