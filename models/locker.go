package models

type (
	Locker struct {
		ID   	int 	`json:"id" form:"id" gorm:"primaryKey"`
		Code 	string 	`json:"code" form:"code" gorm:"unique;not null"`
		UserId	int 	`json:"user_id" form:"user_id" gorm:"not null"`
		User	User	
	}
)