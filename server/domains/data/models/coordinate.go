package models

type Coordinate struct {
	Latitude  float64
	Longitude float64
	Accuracy  float64
}

// CoordinateBuilder Builder Object for Coordinate
type CoordinateBuilder struct {
	latitude  float64
	longitude float64
	accuracy  float64
}

// NewCoordinateBuilder Constructor for CoordinateBuilder
func NewCoordinateBuilder() *CoordinateBuilder {
	o := new(CoordinateBuilder)
	return o
}

// Build Method which creates Coordinate
func (b *CoordinateBuilder) Build() *Coordinate {
	o := new(Coordinate)
	o.Latitude = b.latitude
	o.Longitude = b.longitude
	o.Accuracy = b.accuracy
	return o
}

// SetLatitude Setter method for the field latitude of type float64 in the object CoordinateBuilder
func (b *CoordinateBuilder) SetLatitude(latitude float64) {
	b.latitude = latitude
}

// SetLongitude Setter method for the field longitude of type float64 in the object CoordinateBuilder
func (b *CoordinateBuilder) SetLongitude(longitude float64) {
	b.longitude = longitude
}

// SetAccuracy Setter method for the field latitude of type float64 in the object CoordinateBuilder
func (b *CoordinateBuilder) SetAccuracy(accuracy float64) {
	b.accuracy = accuracy
}
