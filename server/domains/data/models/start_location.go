package models

type StartLocation struct {
	BaseModel
	Type      string
	VillageId *int64
	Coordinate
	Village       *Village `gorm:"foreignKey:VillageId"`
	StreetAddress string   `gorm:"size:100"`
	Description   string
}
