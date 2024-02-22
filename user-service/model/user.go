package model

type User struct {
	BaseModel
	Name        string `json:"name" gorm:"column:name;type:varchar(255);not null"`
	UserName    string `json:"user_name" gorm:"column:user_name;unique;type:varchar(255);not null"`
	Password    string `json:"password" gorm:"column:password;type:varchar(255);not null"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number;type:varchar(11);not null"`
	Address     string `json:"address" gorm:"column:address;type:text;not null"`
	IsHelper    bool   `json:"is_helper" gorm:"column:is_helper;type:boolean;not null"`
}

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
