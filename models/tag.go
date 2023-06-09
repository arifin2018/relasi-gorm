package models

type (
	Tag struct{
		ID		int		`json:"id" form:"id" gorm:"primaryKey"`
		Name	string	`json:"name" form:"name" gorm:"not null"`
	}
)