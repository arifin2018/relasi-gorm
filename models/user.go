package models

type (
	User struct {
		ID   	int `json:"id" form:"id" gorm:"primaryKey"`
		Name 	string `json:"name" form:"name" gorm:"not null"`
		Locker 	[]LockerResponse `json:"locker"`
		Post 	[]Post `json:"post"`
	}

	LockerUser struct {
		ID   	int `json:"id" form:"id" gorm:"primaryKey"`
		Name 	string `json:"name" form:"name" gorm:"not null"`
	}
)

func (LockerUser) TableName() string {
	return "users"
  }