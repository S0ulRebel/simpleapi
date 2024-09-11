package model

type User struct {
	ID        string `json:"ID" gorm:"size:100;not null" bson:"_id"`
	FirstName string `json:"firstName" gorm:"size:100;not null" validate:"required"`
	LastName  string `json:"lastName" gorm:"size:100;not null" validate:"required"`
	Email     string `json:"email" gorm:"uniqueIndex;not null" validate:"required,email"`
	Password  string `json:"password" gorm:"not null" validate:"required"`
}
