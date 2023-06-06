package models

type (
	User struct {
		ID   int `json:"id" form:"id" gorm:"primaryKey"`
		Name int `json:"name" form:"name" gorm:"primaryKey"`
	}
)