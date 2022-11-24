package models

type Image struct {
	BaseModel
	FileName string
}

// Builder Object for Image
type ImageBuilder struct {
	BaseModelBuilder
	fileName string
}

// Constructor for ImageBuilder
func NewImageBuilder() *ImageBuilder {
	o := new(ImageBuilder)
	return o
}

// Build Method which creates Image
func (b *ImageBuilder) Build() *Image {
	o := new(Image)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.FileName = b.fileName
	return o
}

// Setter method for the field fileName of type string in the object ImageBuilder
func (i *ImageBuilder) SetFileName(fileName string) {
	i.fileName = fileName
}
