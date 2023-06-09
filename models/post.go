package models

type (
	Post struct {
		ID   	int 		`json:"id" form:"id" gorm:"primaryKey"`
		Title 	string 		`json:"title" form:"name" gorm:"not null"`
		Body 	string 		`json:"body" form:"body" gorm:"not null"`
		User_id uint 		`json:"user_id" form:"user_id" gorm:"not null"`
		User 	PostUser 	`json:"user"`
	}

	UserPost struct{
		ID   	int 		`json:"id" form:"id"`
		Title 	string 		`json:"title" form:"name"`
		Body 	string 		`json:"body" form:"body"`
		User_id uint 		`json:"-" form:"user_id"`
	}
)
func (UserPost) TableName() string {
	return "posts"
}