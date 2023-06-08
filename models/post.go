package models

type (
	Post struct {
		ID   	int `json:"id" form:"id" gorm:"primaryKey"`
		Title 	string `json:"title" form:"name" gorm:"not null"`
		Body 	string `json:"body" form:"body" gorm:"not null"`
		User_id uint `json:"user_id" form:"user_id" gorm:"not null"`
	}
)