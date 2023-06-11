package models

type (
	Post struct {
		ID   	int 		`json:"id" form:"id" gorm:"primaryKey"`
		Title 	string 		`json:"title" form:"name" gorm:"not null"`
		Body 	string 		`json:"body" form:"body" gorm:"not null"`
		User_id uint 		`json:"user_id" form:"user_id" gorm:"not null"`
		TagID 	[]int 		`json:"tag_id" form:"tag_id" gorm:"-"`
		User 	PostUser 	`json:"user"`
		Tag 	[]Tag 		`json:"tag" gorm:"many2many:Post_Tags"`
	}

	UserPost struct{
		ID   	int 		`json:"id" form:"id"`
		Title 	string 		`json:"title" form:"name"`
		Body 	string 		`json:"body" form:"body"`
		User_id uint 		`json:"-" form:"user_id"`
	}

	ResponsePostWithTag struct {
		ID   	int 		`json:"id" form:"id" gorm:"primaryKey"`
		Title 	string 		`json:"title" form:"name"`
		Body 	string 		`json:"body" form:"body"`
		User_id uint 		`json:"-" form:"user_id"`
		TagID 	[]int 		`json:"-" form:"tag_id" gorm:"-"`
		User 	PostUser 	`json:"user"`
		Tag 	[]Tag 		`json:"tag" gorm:"many2many:Post_Tags;foreignKey:ID;joinForeignKey:post_id;References:ID;joinReferences:tag_id"`
	}
)
func (UserPost) TableName() string {
	return "posts"
}

func (ResponsePostWithTag) TableName() string {
	return "posts"
}