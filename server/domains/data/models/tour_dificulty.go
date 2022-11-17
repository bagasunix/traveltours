package models

type TourDifficulty struct {
	BaseModel
	Name string
}

// Builder Object for TourDificulty
type TourDifficultyBuilder struct {
	BaseModelBuilder
	name string
}

// Constructor for TourDifficultyBuilder
func NewTourDifficultyBuilder() *TourDifficultyBuilder {
	o := new(TourDifficultyBuilder)
	return o
}

// Build Method which creates TourDificulty
func (b *TourDifficultyBuilder) Build() *TourDifficulty {
	o := new(TourDifficulty)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.Name = b.name
	return o
}

// Setter method for the field name of type string in the object TourDifficultyBuilder
func (t *TourDifficultyBuilder) SetName(name string) {
	t.name = name
}
