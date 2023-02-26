package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
