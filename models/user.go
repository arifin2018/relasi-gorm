package models

type (
	User struct {
		ID   	int `json:"id" form:"id" gorm:"primaryKey"`
		Name 	string `json:"name" form:"name" gorm:"not null"`
		Locker 	LockerResponse `json:"locker"`
		Post 	[]UserPost `json:"post"`
	}

	LockerUser struct {
		ID   	int `json:"id" form:"id" gorm:"primaryKey"`
		Name 	string `json:"name" form:"name" gorm:"not null"`
	}

	PostUser struct {
		ID   	int `json:"id" form:"id" gorm:"primaryKey"`
		Name 	string `json:"name" form:"name" gorm:"not null"`
	}
)

func (LockerUser) TableName() string {
	return "users"
}
func (PostUser) TableName() string {
	return "users"
}