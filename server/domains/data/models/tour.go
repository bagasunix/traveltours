package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Tour struct {
	BaseModel
	Name            string
	Slug            string
	Duration        int
	MaxGroupSize    int
	DifficultyId    uuid.UUID
	Difficulty      *TourDifficulty `gorm:"foreignKey:DifficultyId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	RatingsAverage  float32
	RatingsQuantity int64
	Price           int64
	Summary         string `gorm:"size:100"`
	Description     string
	Cover           uuid.UUID
	CoverId         *Image  `gorm:"foreignKey:Cover;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	Images          []Image `gorm:"type:uuid[]"`
	// ImagesId        *Image  `gorm:"foreignKey:Images;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	StartDates    time.Time
	SecretTour    int8 `gorm:"size:1"`
	StartLocation uuid.UUID
	LocationId    StartLocation  `gorm:"foreignKey:StartLocation;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	Location      []LocationTour `gorm:"type:uuid[]"`
	Guides        []User         `gorm:"type:uuid[]"`
	// UserId        User           `gorm:"foreignKey:Guides;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	IsActive int8 `gorm:"size:1"`
}

// Builder Object for Tour
type TourBuilder struct {
	BaseModelBuilder
	name            string
	slug            string
	duration        int
	maxGroupSize    int
	difficultyId    uuid.UUID
	difficulty      *TourDifficulty
	ratingsAverage  float32
	ratingsQuantity int64
	price           int64
	summary         string
	description     string
	cover           uuid.UUID
	coverId         *Image
	images          []Image
	// imagesId        *Image
	startDates    time.Time
	secretTour    int8
	startLocation uuid.UUID
	locationId    StartLocation
	// location        []LocationTour
	guides []User
	// userId   User
	isActive int8
}

// Constructor for TourBuilder
func NewTourBuilder() *TourBuilder {
	o := new(TourBuilder)
	return o
}

// Build Method which creates Tour
func (b *TourBuilder) Build() *Tour {
	o := new(Tour)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.Name = b.name
	o.Slug = b.slug
	o.Duration = b.duration
	o.MaxGroupSize = b.maxGroupSize
	o.DifficultyId = b.difficultyId
	o.Difficulty = b.difficulty
	o.RatingsAverage = b.ratingsAverage
	o.RatingsQuantity = b.ratingsQuantity
	o.Price = b.price
	o.Summary = b.summary
	o.Description = b.description
	o.Cover = b.cover
	o.CoverId = b.coverId
	o.Images = b.images
	// o.ImagesId = b.imagesId
	o.StartDates = b.startDates
	o.SecretTour = b.secretTour
	o.StartLocation = b.startLocation
	o.LocationId = b.locationId
	// o.Location = b.location
	o.Guides = b.guides
	// o.UserId = b.userId
	o.IsActive = b.isActive
	return o
}

// Setter method for the field name of type string in the object TourBuilder
func (t *TourBuilder) SetName(name string) {
	t.name = name
}

// Setter method for the field slug of type string in the object TourBuilder
func (t *TourBuilder) SetSlug(slug string) {
	t.slug = slug
}

// Setter method for the field duration of type int in the object TourBuilder
func (t *TourBuilder) SetDuration(duration int) {
	t.duration = duration
}

// Setter method for the field maxGroupSize of type int in the object TourBuilder
func (t *TourBuilder) SetMaxGroupSize(maxGroupSize int) {
	t.maxGroupSize = maxGroupSize
}

// Setter method for the field difficultyId of type uuid.UUID in the object TourBuilder
func (t *TourBuilder) SetDifficultyId(difficultyId uuid.UUID) {
	t.difficultyId = difficultyId
}

// Setter method for the field difficulty of type *TourDifficulty in the object TourBuilder
func (t *TourBuilder) SetDifficulty(difficulty *TourDifficulty) {
	t.difficulty = difficulty
}

// Setter method for the field ratingsAverage of type float32 in the object TourBuilder
func (t *TourBuilder) SetRatingsAverage(ratingsAverage float32) {
	t.ratingsAverage = ratingsAverage
}

// Setter method for the field ratingsQuantity of type int64 in the object TourBuilder
func (t *TourBuilder) SetRatingsQuantity(ratingsQuantity int64) {
	t.ratingsQuantity = ratingsQuantity
}

// Setter method for the field price of type int64 in the object TourBuilder
func (t *TourBuilder) SetPrice(price int64) {
	t.price = price
}

// Setter method for the field summary of type string in the object TourBuilder
func (t *TourBuilder) SetSummary(summary string) {
	t.summary = summary
}

// Setter method for the field description of type string in the object TourBuilder
func (t *TourBuilder) SetDescription(description string) {
	t.description = description
}

// Setter method for the field startDates of type time.Time in the object TourBuilder
func (t *TourBuilder) SetStartDates(startDates time.Time) {
	t.startDates = startDates
}

// Setter method for the field secretTour of type int8 in the object TourBuilder
func (t *TourBuilder) SetSecretTour(secretTour int8) {
	t.secretTour = secretTour
}

// Setter method for the field guides of type []uuid.UUID in the object TourBuilder
func (t *TourBuilder) SetGuides(guides []User) {
	t.guides = guides
}

// Setter method for the field userId of type []User in the object TourBuilder
// func (t *TourBuilder) SetUserId(userId User) {
// 	t.userId = userId
// }

// Setter method for the field isActive of type int8 in the object TourBuilder
func (t *TourBuilder) SetIsActive(isActive int8) {
	t.isActive = isActive
}

// Setter method for the field cover of type uuid.UUID in the object TourBuilder
func (t *TourBuilder) SetCover(cover uuid.UUID) {
	t.cover = cover
}

// Setter method for the field coverId of type *Image in the object TourBuilder
func (t *TourBuilder) SetCoverId(coverId *Image) {
	t.coverId = coverId
}

// Setter method for the field images of type []uuid.UUID in the object TourBuilder
func (t *TourBuilder) SetImages(images []Image) {
	t.images = images
}

// Setter method for the field imagesId of type *Image in the object TourBuilder
// func (t *TourBuilder) SetImagesId(imagesId *Image) {
// 	t.imagesId = imagesId
// }

// Setter method for the field startLocation of type *StartLocation in the object TourBuilder
func (t *TourBuilder) SetStartLocation(startLocation uuid.UUID) {
	t.startLocation = startLocation
}

// Setter method for the field locationId of type StartLocation in the object TourBuilder
func (t *TourBuilder) SetLocationId(locationId StartLocation) {
	t.locationId = locationId
}
