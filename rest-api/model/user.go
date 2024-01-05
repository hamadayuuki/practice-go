package model

import "time"

type User struct {
//  変数      型        JSONのキー と 制約
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID uint `json:"id" gorm:"primaryKey`
	Email string `json:"email" gorm:"unique"`
}
