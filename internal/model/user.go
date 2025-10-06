package model

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-"` // store hashed passwords
	Tasks    []Task `json:"tasks"`
}
