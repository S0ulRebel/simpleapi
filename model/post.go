package model

type Post struct {
	ID     string `json:"ID" gorm:"size:100;not null" bson:"_id"`
	UserID string `gorm:"not null" json:"UserID" validate:"required"`
	Title  string `gorm:"size:100;not null" json:"Title" validate:"required"`
	Body   string `gorm:"not null" json:"Body" validate:"required"`
}
