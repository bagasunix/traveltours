package models

import "github.com/gofrs/uuid"

type TourReview struct {
	BaseModel
	Rating float32
	TourId uuid.UUID
	Tour   *Tour `gorm:"foreignKey:TourId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	UserId uuid.UUID
	User   *User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
}

// Builder Object for TourReview
type TourReviewBuilder struct {
	BaseModelBuilder
	rating float32
	tourId uuid.UUID
	tour   *Tour
	userId uuid.UUID
	user   *User
}

// Constructor for TourReviewBuilder
func NewTourReviewBuilder() *TourReviewBuilder {
	o := new(TourReviewBuilder)
	return o
}

// Build Method which creates TourReview
func (b *TourReviewBuilder) Build() *TourReview {
	o := new(TourReview)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.Rating = b.rating
	o.TourId = b.tourId
	o.Tour = b.tour
	o.UserId = b.userId
	o.User = b.user
	return o
}

// Setter method for the field rating of type float32 in the object TourReviewBuilder
func (t *TourReviewBuilder) SetRating(rating float32) {
	t.rating = rating
}

// Setter method for the field tourId of type uuid.UUID in the object TourReviewBuilder
func (t *TourReviewBuilder) SetTourId(tourId uuid.UUID) {
	t.tourId = tourId
}

// Setter method for the field tour of type *Tour in the object TourReviewBuilder
func (t *TourReviewBuilder) SetTour(tour *Tour) {
	t.tour = tour
}

// Setter method for the field userId of type uuid.UUID in the object TourReviewBuilder
func (t *TourReviewBuilder) SetUserId(userId uuid.UUID) {
	t.userId = userId
}

// Setter method for the field user of type *User in the object TourReviewBuilder
func (t *TourReviewBuilder) SetUser(user *User) {
	t.user = user
}
